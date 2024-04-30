package repositories

import (
	"database/sql"
	"discord-spotify-bot/config"
	"discord-spotify-bot/models"
	"fmt"
)

func GetArtist(artistName string) (models.Artist, error) {
	var artist models.Artist
	query := `SELECT
				id,
				name,
				popularity,
				genres,
				ur,
				followers,
				image,
				max_danceability,
				max_danceability_track,
				max_energy,
				max_energy_track,
				albums_count,
				singles_count,
				compilations_count,
				tracks_count,
				duration_ms
			FROM artist
			WHERE artist_name = $1`
	result := config.DB.QueryRow(query, artistName)
	err := result.Scan(
		&artist.Id,
		&artist.Name,
		&artist.Popularity,
		&artist.Genres,
		&artist.Url,
		&artist.Followers,
		&artist.Image,
		&artist.MaxDanceability,
		&artist.MaxDanceabilityTrack,
		&artist.MaxEnergy,
		&artist.MaxEnergyTrack,
		&artist.AlbumsCount,
		&artist.SinglesCount,
		&artist.CompilationsCount,
		&artist.TracksCount,
		&artist.DurationMs,
	)
	if err == sql.ErrNoRows {
		return artist, nil
	}
	if err != nil {
		return artist, fmt.Errorf("error scaning rows: %v", err)
	}
	return artist, nil
}
