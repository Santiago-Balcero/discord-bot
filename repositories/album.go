package repositories

import (
	"discord-spotify-bot/config"
	"discord-spotify-bot/models"
	"fmt"
)

func SaveAlbum(artistId int, album models.Album) (int, error) {
	query :=
		`INSERT INTO album (
			artist_id,
			spotify_id,
			name,
			type,
			release_date,
			url,
			image,
			max_danceability,
			max_danceability_track,
			max_energy,
			max_energy_track,
			tracks_count,
			duration_ms
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)
		RETURNING album_id`
	result := config.DB.QueryRow(
		query,
		artistId,
		album.SpotifyId,
		album.Name,
		album.Type,
		album.ReleaseDate,
		album.Url,
		album.Image,
		album.MaxDanceability,
		album.MaxDanceabilityTrack,
		album.MaxEnergy,
		album.MaxEnergyTrack,
		album.TracksCount,
		album.DurationMs,
	)
	var newAlbumId int
	if err := result.Scan(&newAlbumId); err != nil {
		return 0, fmt.Errorf("error in save album: %v", err)
	}
	return newAlbumId, nil
}

func GetArtistAlbums(artistId int) ([]models.Album, error) {
	query :=
		`SELECT
			spotify_id,
			name,
			type,
			release_date,
			url,
			image,
			max_danceability,
			max_danceability_track,
			max_energy,
			max_energy_track,
			tracks_count,
			duration_ms,
			created_at,
			updated_at
		FROM album
		WHERE artist_id = $1`
	result, err := config.DB.Query(
		query,
		artistId,
	)
	var albums []models.Album
	if err != nil {
		return albums, fmt.Errorf("error in query: %v", err)
	}
	for result.Next() {
		var album models.Album
		err = result.Scan(
			&album.SpotifyId,
			&album.Name,
			&album.Type,
			&album.ReleaseDate,
			&album.Url,
			&album.Image,
			&album.MaxDanceability,
			&album.MaxDanceabilityTrack,
			&album.MaxEnergy,
			&album.MaxEnergyTrack,
			&album.TracksCount,
			&album.DurationMs,
			&album.CreatedAt,
			&album.UpdatedAt,
		)
		if err != nil {
			return albums, fmt.Errorf("error in rows scan: %v", err)
		}
		albums = append(albums, album)
	}
	return albums, nil
}
