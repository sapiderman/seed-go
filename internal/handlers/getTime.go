package handlers

import (
	"io"
	"net/http"
	"time"

	"github.com/sapiderman/seed-go/internal/helpers"
)

// GetTime
func GetTime(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	// w.WriteHeader(http.StatusOK)

	resp, err := http.Get("http://worldtimeapi.org/api/ip")
	if err != nil {
		helpers.HTTPResponseBuilder(r.Context(), w, r, http.StatusOK, "time server unavailable", err.Error())
		return
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		helpers.HTTPResponseBuilder(r.Context(), w, r, http.StatusOK, "time server response wierd", err.Error())
		return
	}

	helpers.HTTPResponseBuilder(r.Context(), w, r, http.StatusOK, "", string(body))

}

// GetTime
func GetTimeOut(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	time.Sleep(15 * time.Second)

	helpers.HTTPResponseBuilder(r.Context(), w, r, http.StatusRequestTimeout, "", "timeout")

}
