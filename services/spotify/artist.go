package services

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/Santiago-Balcero/discord-bot/models"
	client "github.com/Santiago-Balcero/discord-bot/services/http"
)

func GetArtist(artistName string) (models.Artist, error) {
	artistData := models.Artist{}
	url := os.Getenv("SPOTIFY_SERVICE")
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return artistData, err
	}
	query := request.URL.Query()
	query.Add("name", artistName)
	request.URL.RawQuery = query.Encode()

	log.Println("Calling Spotify Service API:", request)
	responseBytes, err := client.GetClient().Do(request)
	if err != nil {
		return artistData, err
	}
	if responseBytes.StatusCode != 200 {
		return artistData, fmt.Errorf(
			"error in response status: %v",
			responseBytes.StatusCode,
		)
	}
	body, err := io.ReadAll(responseBytes.Body)
	if err != nil {
		return artistData, err
	}
	response := models.ArtistResponse{}
	err = json.Unmarshal(body, &response)
	if err != nil {
		return artistData, err
	}
	log.Println("Response from Spotify Service API:", response.Data)
	artistData = response.Data
	return artistData, nil
}
