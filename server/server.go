package server

import (
	"discord-spotify-bot/config"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type HealthStatus struct {
	Code     int       `json:"code"`
	Datetime time.Time `json:"datetime"`
	Status   string    `json:"status"`
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	log.Println("Health check request from", r.RemoteAddr)
	loc, _ := time.LoadLocation("America/Bogota")
	status := HealthStatus{
		Code:     200,
		Datetime: time.Now().In(loc),
		Status:   "OK",
	}
	json.NewEncoder(w).Encode(status)
}

func InitServer() {
	log.Println("Assigned port:", config.Port)
	http.HandleFunc("/", healthCheck)
	go http.ListenAndServe(":"+config.Port, nil)
}
