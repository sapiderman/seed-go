package handlers

import (
	"fmt"
	"net/http"

	"github.com/sapiderman/seed-go/internal/connector"
)

// ListUsers lists all users
func ListUsers(w http.ResponseWriter, r *http.Request) {

	users, err := connector.ListAllUsers()
	if err != nil {

		w.WriteHeader(http.StatusNotImplemented)
	}

	fmt.Println(users)
	w.WriteHeader(http.StatusOK)

	// w.Write([]byte(users))
}

// ListDevices lists all users
func ListDevices(w http.ResponseWriter, r *http.Request) {

	users, err := connector.ListAllDevices()
	if err != nil {

		w.WriteHeader(http.StatusNotImplemented)
	}

	fmt.Println(users)
	w.WriteHeader(http.StatusOK)

}

// AddDevice adds a device to database
func AddDevice(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)

}

// AddUser adds a user to database
func AddUser(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)

}
