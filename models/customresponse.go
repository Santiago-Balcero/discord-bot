package models

type CustomResponse struct {
	Code    int
	Message string
}

type ArtistResponse struct {
	CustomResponse
	Data Artist
}
