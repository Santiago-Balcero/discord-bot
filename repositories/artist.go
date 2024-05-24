package repositories

import (
	"database/sql"
	"discord-spotify-bot/config"
	"discord-spotify-bot/models"
	"fmt"
	"strings"

	"github.com/lib/pq"
)

func GetArtist(artistName string) (models.Artist, error) {
	var artist models.Artist
	query :=
		`SELECT
			artist_id,
			spotify_id,
			name,
			popularity,
			genres,
			url,
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
			duration_ms,
			created_at,
			updated_at
		FROM artist
		WHERE name = $1`
	result := config.DB.QueryRow(query, artistName)
	err := result.Scan(
		&artist.ArtistId,
		&artist.SpotifyId,
		&artist.Name,
		&artist.Popularity,
		pq.Array(&artist.Genres),
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
		&artist.CreatedAt,
		&artist.UpdatedAt,
	)
	if err == sql.ErrNoRows {
		return artist, nil
	}
	if err != nil {
		return artist, fmt.Errorf("error in row scan: %v", err)
	}
	return artist, nil
}

func SaveArtist(artist models.Artist) (int, error) {
	query :=
		`INSERT INTO artist (
			spotify_id,
			name,
			popularity,
			genres,
			url,
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
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16)
		RETURNING artist_id`
	result := config.DB.QueryRow(
		query,
		artist.SpotifyId,
		artist.Name,
		artist.Popularity,
		fmt.Sprintf("{%s}", strings.Join(artist.Genres, ",")),
		artist.Url,
		artist.Followers,
		artist.Image,
		artist.MaxDanceability,
		artist.MaxDanceabilityTrack,
		artist.MaxEnergy,
		artist.MaxEnergyTrack,
		artist.AlbumsCount,
		artist.SinglesCount,
		artist.CompilationsCount,
		artist.TracksCount,
		artist.DurationMs,
	)
	var newArtistId int
	if err := result.Scan(&newArtistId); err != nil {
		return 0, fmt.Errorf("error in row scan: %v", err)
	}
	return newArtistId, nil
}
