package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// HealthResponse is just the resp to return
type healthResponse struct {
	ServerUptime     string `json:"serverUptime"`
	ServerUpDuration uint64 `json:"serverUpDuration"`
}

// HandleHealth handles /health calls
func HandleHealth(w http.ResponseWriter, s *http.Request) {
	// fmt.Println("helooooo!!! i'm alive...")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	resp := &healthResponse{
		ServerUptime: time.Now().String(),
	}

	jsonData, err := json.Marshal(resp)
	_, err = w.Write(jsonData)
	if err != nil {
		fmt.Println("Errrrorororroro")
		// log.WithFields(log.Fields{
		// 	"response object": err,
		// }).Error("Error while writing response object")
	}

}
