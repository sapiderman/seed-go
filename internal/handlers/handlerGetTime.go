package handlers

import (
	"io/ioutil"
	"net/http"

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

	w.Write([]byte(body))

}
