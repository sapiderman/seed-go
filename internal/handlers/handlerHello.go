package handlers

import (
	"fmt"
	"net/http"
)

// HandlerHello handles /hello calls
func HandlerHello(w http.ResponseWriter, s *http.Request) {

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)

	resp := []byte("hello back")

	_, err := w.Write(resp)
	if err != nil {
		fmt.Println("resp error: ", err)

	}

}
