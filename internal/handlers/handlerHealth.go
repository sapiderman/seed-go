package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// HealthResponse is just the resp to return
type healthResponse struct {
	ServerStatus     string `json:"serverStatus"`
	ServerTime       string `json:"serverTime"`
	ServerUpDuration uint64 `json:"serverUpDuration"`

	ServerVersion string `json:"serverVersion"`
}

// HandlerHealth handles /health calls
func HandlerHealth(w http.ResponseWriter, s *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	resp := &healthResponse{
		ServerStatus:     "I'm Aliiiive",
		ServerTime:       time.Now().String(),
		ServerUpDuration: 0,
		ServerVersion:    "0.0.0.0",
	}

	jsonData, err := json.Marshal(resp)
	if err != nil {
		fmt.Println("json error: ", err)
	}

	_, err = w.Write(jsonData)
	if err != nil {
		fmt.Println("json error: ", err)
	}

}
