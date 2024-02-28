package services

import (
	"net/http"
)

func GetClient() *http.Client {
	return &http.Client{}
}
