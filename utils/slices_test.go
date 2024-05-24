package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zmb3/spotify/v2"
)

var (
	calypsa = spotify.SimpleArtist{
		Name: "Calypsa",
	}
	kika = spotify.SimpleArtist{
		Name: "Kika",
	}
	artists = []spotify.SimpleArtist{calypsa, kika}
)

func TestArtistInListArtistFound(t *testing.T) {
	result := ArtistInList(artists, "Kika")
	assert.Equal(t, true, result, "Should return true if artist in list")
}

func TestArtistInListArtistNotFound(t *testing.T) {
	result := ArtistInList(artists, "Santiago")
	assert.Equal(t, false, result, "Should return false if artist in list")
}
