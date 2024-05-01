package repositories

import (
	"discord-spotify-bot/config"
	"discord-spotify-bot/models"
	"fmt"
)

func SaveTrack(albumId int, track models.Track) (int, error) {
	query :=
		`INSERT INTO track (
			album_id,
			spotify_id,
			name,
			url,
			danceability,
			energy,
			acousticness,
			loudness,
			liveness,
			instrumentalness,
			tempo,
			duration_ms
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
		RETURNING track_id`
	result := config.DB.QueryRow(
		query,
		albumId,
		track.SpotifyId,
		track.Name,
		track.Url,
		track.Danceability,
		track.Energy,
		track.Acousticness,
		track.Loudness,
		track.Liveness,
		track.Instrumentalness,
		track.Tempo,
		track.DurationMs,
	)
	var newTrackId int
	if err := result.Scan(&newTrackId); err != nil {
		return 0, fmt.Errorf("error in save track: %v", err)
	}
	return newTrackId, nil
}

func GetAlbumTracks(albumId int) ([]models.Track, error) {
	query :=
		`SELECT
			spotify_id,
			name,
			url,
			danceability,
			energy,
			acousticness,
			loudness,
			liveness,
			instrumentalness,
			tempo,
			duration_ms,
			created_at,
			updated_at
		FROM track
		WHERE album_id = $1`
	result, err := config.DB.Query(
		query,
		albumId,
	)
	var tracks []models.Track
	if err != nil {
		return tracks, fmt.Errorf("error in query: %v", err)
	}
	for result.Next() {
		var track models.Track
		err = result.Scan(
			&track.SpotifyId,
			&track.Name,
			&track.Url,
			&track.Danceability,
			&track.Energy,
			&track.Acousticness,
			&track.Loudness,
			&track.Liveness,
			&track.Instrumentalness,
			&track.Tempo,
			&track.DurationMs,
			&track.CreatedAt,
			&track.UpdatedAt,
		)
		if err != nil {
			return tracks, fmt.Errorf("error in rows scan: %v", err)
		}
		tracks = append(tracks, track)
	}
	return tracks, nil
}
