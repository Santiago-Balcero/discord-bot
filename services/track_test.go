package services

import (
	"context"
	"discord-spotify-bot/models"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/zmb3/spotify/v2"
)

// Mock for Spotify client
type MockSpotifyClient struct {
	mock.Mock
}

func (m *MockSpotifyClient) GetAudioFeatures(
	ctx context.Context,
	ids ...spotify.ID,
) ([]*spotify.AudioFeatures, error) {
	args := m.Called(ctx, ids)
	return args.Get(0).([]*spotify.AudioFeatures), args.Error(1)
}

func TestAnalyseTrack(t *testing.T) {
	mockClient := new(MockSpotifyClient)

	track := &models.Track{
		Id:   "your_track_id",
		Name: "Your Track Name",
	}

	mockClient.On("GetAudioFeatures", mock.Anything, []spotify.ID{spotify.ID(track.Id)}).
		Return([]*spotify.AudioFeatures{{Danceability: 0.8, Energy: 0.7}}, nil)

	err := AnalyseTrack(mockClient, track)
	assert.Nil(t, err, "Error should be nil")
	assert.NotEmpty(
		t,
		track.Danceability,
		"Shoud set danceability",
	)
	assert.NotEmpty(
		t,
		track.Energy,
		"Shoud set energy",
	)

	mockClient.AssertExpectations(t)
}
