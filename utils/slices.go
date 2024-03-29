package utils

import "github.com/zmb3/spotify/v2"

func ArtistInList(list []spotify.SimpleArtist, artistName string) bool {
	for i := range list {
		if list[i].Name == artistName {
			return true
		}
	}
	return false
}
