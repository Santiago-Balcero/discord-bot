package services

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"discord-spotify-bot/models"
	"discord-spotify-bot/repositories"
	"discord-spotify-bot/utils"

	"github.com/zmb3/spotify/v2"
)

func GetArtist(client *spotify.Client, artistName string) (string, error) {
	var artistData models.Artist
	ctx := context.Background()
	artistData, err := searchArtist(ctx, client, artistName)
	if err != nil {
		return fmt.Sprintf("Artist not found: %s", artistName), err
	}

	dbArtist, err := repositories.GetArtist(artistData.Name)
	if err != nil {
		return "Internal error.", err
	}
	if dbArtist.Name != "" {
		log.Println("Artist found in database")
		err = getArtistAlbums(dbArtist.ArtistId, &dbArtist)
		if err != nil {
			return "Internal error.", err
		}
		return artistToString(&dbArtist), nil
	}

	err = analyseArtist(ctx, client, &artistData)
	if err != nil {
		return "Please try again in a few minutes.", err
	}
	if err = saveArtistData(&artistData); err != nil {
		return "Internal error.", err
	}
	return artistToString(&artistData), nil
}

func searchArtist(
	ctx context.Context,
	client *spotify.Client,
	artistName string,
) (models.Artist, error) {
	artist := models.Artist{}
	result, err := client.Search(ctx, artistName, spotify.SearchTypeArtist)
	if err != nil {
		return artist, fmt.Errorf("error searching artist: %v", err)
	}

	artistFound := false

	for _, a := range result.Artists.Artists {
		name := utils.ClearString(a.Name)
		if name == strings.ReplaceAll(artistName, " ", "") {
			artist.SpotifyId = a.ID.String()
			artist.Name = a.Name
			artist.Popularity = a.Popularity
			artist.Genres = a.Genres
			artist.Url = a.ExternalURLs["spotify"]
			artist.Followers = int(a.Followers.Count)
			artist.Image = string(a.Images[2].URL)
			log.Println("Artist found:", artist.Name, artist.Url)
			artistFound = true
			break
		}
	}

	if !artistFound {
		return artist, fmt.Errorf("artist not found")
	}

	return artist, nil
}

func analyseArtist(
	ctx context.Context,
	client *spotify.Client,
	artistData *models.Artist,
) error {
	artist, err := client.GetArtist(ctx, spotify.ID(artistData.SpotifyId))
	if err != nil {
		return fmt.Errorf("error in GetArtist: %v", err)
	}

	loc, _ := time.LoadLocation("America/Bogota")
	startTime := time.Now().In(loc)

	albumTypesSearch := []spotify.AlbumType{
		spotify.AlbumTypeAlbum,
		spotify.AlbumTypeAppearsOn,
		spotify.AlbumTypeCompilation,
		spotify.AlbumTypeSingle,
	}

	albums, err := client.GetArtistAlbums(
		ctx,
		spotify.ID(artistData.SpotifyId),
		albumTypesSearch,
	)
	if err != nil {
		return fmt.Errorf("error in GetArtistAlbums: %v", err)
	}

	for _, album := range albums.Albums {
		tracks, err := client.GetAlbumTracks(ctx, album.ID)
		if err != nil {
			return fmt.Errorf("error in GetAlbumTracks: %v", err)
		}
		albumData := models.Album{
			SpotifyId:   album.ID.String(),
			Name:        album.Name,
			Type:        album.AlbumType,
			ReleaseDate: album.ReleaseDate,
			Url:         album.ExternalURLs["spotify"],
			Image:       string(album.Images[2].URL),
			Tracks:      []models.Track{},
		}
		for _, track := range tracks.Tracks {
			if utils.ArtistInList(track.Artists, artist.Name) {
				artistData.TracksCount++
				albumData.DurationMs += track.Duration
				albumData.DurationMs += track.Duration
				trackData := models.Track{
					SpotifyId:  track.ID.String(),
					Name:       track.Name,
					Url:        track.ExternalURLs["spotify"],
					DurationMs: track.Duration,
				}
				err := AnalyseTrack(client, &trackData)
				if err != nil {
					return fmt.Errorf("error in AnalyseTrack: %v", err)
				}
				checkAlbumMaximums(&trackData, &albumData)
				albumData.Tracks = append(albumData.Tracks, trackData)
				albumData.TracksCount++
			}
		}
		addToArtistDiscography(&albumData, artistData)
		checkArtistMaximums(&albumData, artistData)
	}

	fetchTime := time.Since(startTime)
	log.Println("Artist data was fetched in:", fetchTime)
	return nil
}

func checkAlbumMaximums(trackData *models.Track, albumData *models.Album) {
	if trackData.Danceability > albumData.MaxDanceability {
		albumData.MaxDanceability = trackData.Danceability
		albumData.MaxDanceabilityTrack = trackData.Name
	}
	if trackData.Energy > albumData.MaxEnergy {
		albumData.MaxEnergy = trackData.Energy
		albumData.MaxEnergyTrack = trackData.Name
	}
}

func checkArtistMaximums(albumData *models.Album, artistData *models.Artist) {
	if albumData.MaxDanceability > artistData.MaxDanceability {
		artistData.MaxDanceability = albumData.MaxDanceability
		artistData.MaxDanceabilityTrack = albumData.MaxDanceabilityTrack
	}
	if albumData.MaxEnergy > artistData.MaxEnergy {
		artistData.MaxEnergy = albumData.MaxEnergy
		artistData.MaxEnergyTrack = albumData.MaxEnergyTrack
	}
}

func addToArtistDiscography(albumData *models.Album, artist *models.Artist) {
	if albumData.Type == "album" {
		artist.Albums = append(artist.Albums, *albumData)
		artist.AlbumsCount++
	} else if albumData.Type == "single" {
		artist.Singles = append(artist.Singles, *albumData)
		artist.SinglesCount++
	} else if albumData.Type == "compialtion" {
		artist.Compilations = append(artist.Compilations, *albumData)
		artist.CompilationsCount++
	}
}

func artistToString(artist *models.Artist) string {
	genres := "**Genres:** not found"
	if len(artist.Genres) > 0 {
		genres = fmt.Sprintf(
			"**Genres:** %s",
			strings.Join(artist.Genres, ", "),
		)
	}
	popularity := fmt.Sprintf("**Popularity:** %v", artist.Popularity)
	followers := fmt.Sprintf(
		"**Followers:** %v",
		utils.FormatInteger(artist.Followers),
	)
	albums := fmt.Sprintf("Total albums: %v", artist.AlbumsCount)
	singles := fmt.Sprintf("Total singles: %v", artist.SinglesCount)
	compilations := fmt.Sprintf(
		"Total compilations: %v",
		artist.CompilationsCount,
	)
	tracks := fmt.Sprintf(
		"Total tracks: %v",
		utils.FormatInteger(artist.TracksCount),
	)

	albumsInfo := "**Albums:**"
	for i := range artist.Albums {
		tracksText := "tracks"
		if artist.Albums[i].TracksCount == 1 {
			tracksText = "track"
		}
		albumsInfo += fmt.Sprintf(
			"\n\t• %s (%s), %v %s.",
			artist.Albums[i].Name,
			strings.Split(artist.Albums[i].ReleaseDate, "-")[0],
			artist.Albums[i].TracksCount,
			tracksText,
		)
	}

	singlesInfo := ""
	for i := range artist.Singles {
		singlesInfo = "\n\n**Singles:**"
		tracksText := ""
		if artist.Singles[i].TracksCount > 1 {
			tracksText = fmt.Sprintf(
				", %v tracks",
				artist.Singles[i].TracksCount,
			)
		}
		singlesInfo += fmt.Sprintf(
			"\n\t• %s (%s)%s.",
			artist.Singles[i].Name,
			strings.Split(artist.Singles[i].ReleaseDate, "-")[0],
			tracksText,
		)
	}

	compilationsInfo := ""
	for i := range artist.Compilations {
		compilationsInfo = "\n\n**Compilations:**"
		compilationsInfo += fmt.Sprintf(
			"\n\t• %s (%s) - track: %s.",
			artist.Compilations[i].Name,
			strings.Split(artist.Compilations[i].ReleaseDate, "-")[0],
			artist.Compilations[i].Tracks[0].Name,
		)
	}

	artistStr := fmt.Sprintf(
		"%s\n%s\n\n%s\n%s\n%s\n\n%s\n%s\n%s\n%s\n\n%s%s%s",
		fmt.Sprintf("**%s**", strings.ToUpper(artist.Name)),
		artist.Url,
		genres,
		popularity,
		followers,
		albums,
		singles,
		compilations,
		tracks,
		albumsInfo,
		singlesInfo,
		compilationsInfo,
	)
	return artistStr
}

func saveArtistData(artist *models.Artist) error {
	newArtistId, err := repositories.SaveArtist(*artist)
	if err != nil {
		return fmt.Errorf("artist %s not saved: %v", artist.Name, err)
	}
	log.Println(
		"Artist",
		fmt.Sprintf("\"%s\"", artist.Name),
		"successfully saved with artist id:",
		newArtistId,
	)

	for _, album := range artist.Albums {
		newAlbumId, err := repositories.SaveAlbum(newArtistId, album)
		if err != nil {
			return fmt.Errorf("album %s not saved: %v", album.Name, err)
		}
		log.Println(
			"Album",
			fmt.Sprintf("\"%s\"", album.Name),
			"sucessfully saved with album id:",
			newAlbumId,
		)

		for _, track := range album.Tracks {
			newTrackId, err := repositories.SaveTrack(newAlbumId, track)
			if err != nil {
				return fmt.Errorf("track %s not saved: %v", track.Name, err)
			}
			log.Println(
				"Track",
				fmt.Sprintf("\"%s\"", track.Name),
				"sucessfully saved with track id:",
				newTrackId,
			)
		}
	}

	for _, single := range artist.Singles {
		newSingleId, err := repositories.SaveAlbum(newArtistId, single)
		if err != nil {
			return fmt.Errorf("single %s not saved: %v", single.Name, err)
		}
		log.Println(
			"Single",
			fmt.Sprintf("\"%s\"", single.Name),
			"sucessfully saved with album id:",
			newSingleId,
		)

		for _, track := range single.Tracks {
			newTrackId, err := repositories.SaveTrack(newSingleId, track)
			if err != nil {
				return fmt.Errorf("track %s not saved: %v", track.Name, err)
			}
			log.Println(
				"Track",
				fmt.Sprintf("\"%s\"", track.Name),
				"sucessfully saved with track id:",
				newTrackId,
			)
		}
	}

	for _, compilation := range artist.Compilations {
		newCompilationId, err := repositories.SaveAlbum(newArtistId, compilation)
		if err != nil {
			return fmt.Errorf("compilation %s not saved: %v", compilation.Name, err)
		}
		log.Println(
			"Compilation",
			fmt.Sprintf("\"%s\"", compilation.Name),
			"sucessfully saved with album id:",
			newCompilationId,
		)

		for _, track := range compilation.Tracks {
			newTrackId, err := repositories.SaveTrack(newCompilationId, track)
			if err != nil {
				return fmt.Errorf("track %s not saved: %v", track.Name, err)
			}
			log.Println(
				"Track",
				fmt.Sprintf("\"%s\"", track.Name),
				"sucessfully saved with track id:",
				newTrackId,
			)
		}
	}
	return nil
}

func getArtistAlbums(artistId int, artistData *models.Artist) error {
	var albums []models.Album
	var singles []models.Album
	var compilations []models.Album
	albums, err := repositories.GetArtistAlbums(artistId)
	if err != nil {
		return fmt.Errorf("error in get artist albums: %v", err)
	}

	for i, album := range albums {
		tracks, err := repositories.GetAlbumTracks(album.AlbumId)
		if err != nil {
			return fmt.Errorf("error in get album tracks: %v", err)
		}
		album.Tracks = tracks
		if album.Type == "single" {
			singles = append(singles, album)
			albums = append(albums[:i], albums[i+1:]...)
		} else if album.Type == "compilation" {
			compilations = append(compilations, album)
			albums = append(albums[:i], albums[i+1:]...)
		}
	}
	artistData.Albums = albums
	artistData.Singles = singles
	artistData.Compilations = compilations
	return nil
}
