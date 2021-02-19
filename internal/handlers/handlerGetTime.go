package handlers

import (
	"context"
	"io/ioutil"
	"net/http"

	"github.com/sapiderman/seed-go/internal/helpers"
	log "github.com/sirupsen/logrus"
)

// GetTime que
func GetTime(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	resp, err := http.Get("http://worldtimeapi.org/api/ip")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	helpers.HTTPResponseBuilder(ctx, w, r, http.StatusOK, "", body)

}
