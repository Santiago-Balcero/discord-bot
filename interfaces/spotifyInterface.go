package interfaces

import (
	"context"

	"github.com/zmb3/spotify/v2"
)

// SpotifyClient defines the methods needed by AnalyseTrack
type SpotifyClient interface {
	GetAudioFeatures(ctx context.Context, ids ...spotify.ID) ([]*spotify.AudioFeatures, error)
	// Add other methods you use here
}
