package handlers

import (
	"io/ioutil"
	"net/http"

	log "github.com/sirupsen/logrus"
)

// HandlerGetTime que
func HandlerGetTime(w http.ResponseWriter, s *http.Request) {

	w.Header().Set("Content-Type", "text/plain")
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

	log.Info("resp: ", body)
	w.Write([]byte(body))

}
