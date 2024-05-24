package services

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"discord-spotify-bot/constants"
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
		return fmt.Sprint(constants.ArtistNotFound, artistName), err
	}

	dbArtist, err := repositories.GetArtist(artistData.Name)
	if err != nil {
		return constants.InternalError, err
	}
	if dbArtist.Name != "" {
		log.Println("Artist found in database")
		err = getArtistAlbums(dbArtist.ArtistId, &dbArtist)
		if err != nil {
			return constants.InternalError, err
		}
		return artistDataToResponse(dbArtist), nil
	}

	err = analyseArtist(ctx, client, &artistData)
	if err != nil {
		return constants.TryAgain, err
	}
	if err = saveArtistData(&artistData); err != nil {
		return constants.InternalError, err
	}
	return artistDataToResponse(artistData), nil
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
				artistData.DurationMs += track.Duration
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
	} else if albumData.Type == "compilation" {
		artist.Compilations = append(artist.Compilations, *albumData)
		artist.CompilationsCount++
	}
}

func saveArtistData(artist *models.Artist) error {
	newArtistId, err := repositories.SaveArtist(*artist)
	if err != nil {
		return fmt.Errorf("artist %s not saved: %v", artist.Name, err)
	}
	log.Println(
		"Artist",
		utils.AddQuotes(artist.Name),
		constants.ArtistSaved,
		newArtistId,
	)

	discography := []models.Album{}
	discography = append(discography, artist.Albums...)
	discography = append(discography, artist.Singles...)
	discography = append(discography, artist.Compilations...)

	for _, album := range discography {
		newAlbumId, err := repositories.SaveAlbum(newArtistId, album)
		if err != nil {
			return fmt.Errorf(
				"%s %s not saved: %v",
				album.Type,
				album.Name,
				err,
			)
		}
		logType := "Album"
		if album.Type == "single" {
			logType = "Single"
		} else if album.Type == "compilation" {
			logType = "Compilation"
		}
		log.Println(
			logType,
			utils.AddQuotes(album.Name),
			constants.AlbumSaved,
			newAlbumId,
		)

		for _, track := range album.Tracks {
			newTrackId, err := repositories.SaveTrack(newAlbumId, track)
			if err != nil {
				return fmt.Errorf("track %s not saved: %v", track.Name, err)
			}
			log.Println(
				"Track",
				utils.AddQuotes(track.Name),
				constants.TrackSaved,
				newTrackId,
			)
		}
	}
	return nil
}

func getArtistAlbums(artistId int, artist *models.Artist) error {
	albums, err := repositories.GetArtistAlbums(artistId)
	if err != nil {
		return fmt.Errorf("error in get artist albums: %v", err)
	}

	for i := range albums {
		albums[i].Tracks, err = repositories.GetAlbumTracks(albums[i].AlbumId)
		if err != nil {
			return fmt.Errorf("error in get album tracks: %v", err)
		}
		if albums[i].Type == "album" {
			artist.Albums = append(artist.Albums, albums[i])
		} else if albums[i].Type == "single" {
			artist.Singles = append(artist.Singles, albums[i])
		} else if albums[i].Type == "compilation" {
			artist.Compilations = append(artist.Compilations, albums[i])
		}
	}
	return nil
}

func getTrackUrl(trackName string, artist models.Artist) string {
	discography := []models.Album{}
	discography = append(discography, artist.Albums...)
	discography = append(discography, artist.Singles...)
	discography = append(discography, artist.Compilations...)

	for i := range discography {
		for j := range discography[i].Tracks {
			if trackName == discography[i].Tracks[j].Name {
				return discography[i].Tracks[j].Url
			}
		}
	}
	return ""
}

func artistDataToResponse(artist models.Artist) string {
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

	albums := fmt.Sprintf("**Total albums:** %v", artist.AlbumsCount)

	singles := fmt.Sprintf("**Total singles:** %v", artist.SinglesCount)

	compilations := fmt.Sprintf(
		"**Total compilations:** %v",
		artist.CompilationsCount,
	)

	tracks := fmt.Sprintf(
		"**Total tracks:** %v",
		utils.FormatInteger(artist.TracksCount),
	)

	totalMusicTime := fmt.Sprintf(
		"**Total music time:** %s",
		utils.MillisecondsToTime(artist.DurationMs),
	)

	maxDanceabilityInfo := fmt.Sprintf(
		"**To dance:** %s [🎧 here](%s)",
		artist.MaxDanceabilityTrack,
		getTrackUrl(artist.MaxDanceabilityTrack, artist),
	)

	maxEnergyInfo := fmt.Sprintf(
		"**Full of energy:** %s [🎧 here](%s)",
		artist.MaxEnergyTrack,
		getTrackUrl(artist.MaxEnergyTrack, artist),
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
	if artist.SinglesCount > 0 {
		singlesInfo = "\n\n**Singles:**"
		for i := range artist.Singles {
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
	}

	compilationsInfo := ""
	if artist.CompilationsCount > 0 {
		compilationsInfo = "\n\n**Compilations:**"
		for i := range artist.Compilations {
			compilationsInfo += fmt.Sprintf(
				"\n\t• %s (%s) - track: %s.",
				artist.Compilations[i].Name,
				strings.Split(artist.Compilations[i].ReleaseDate, "-")[0],
				artist.Compilations[i].Tracks[0].Name,
			)
		}
	}

	artistStr := fmt.Sprintf(
		"%s\n%s\n\n%s\n%s\n%s\n\n%s\n%s\n%s\n%s\n%s\n\n%s\n%s\n\n%s%s%s",
		fmt.Sprintf("**%s**", strings.ToUpper(artist.Name)),
		artist.Url,
		genres,
		popularity,
		followers,
		albums,
		singles,
		compilations,
		tracks,
		totalMusicTime,
		maxDanceabilityInfo,
		maxEnergyInfo,
		albumsInfo,
		singlesInfo,
		compilationsInfo,
	)
	return artistStr
}
