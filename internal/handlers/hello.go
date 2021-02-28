package handlers

import (
	"fmt"
	"net/http"
)

// Hello handles /hello calls
func Hello(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)

	resp := []byte("hello back")

	_, err := w.Write(resp)
	if err != nil {
		fmt.Println("resp error: ", err)
	}

}
