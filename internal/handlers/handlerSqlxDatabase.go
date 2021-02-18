package handlers

import (
	"bytes"
	"fmt"
	"net/http"

	"encoding/json"

	"github.com/sapiderman/seed-go/internal/connector"
)

// MyHandlers wraps all needed connectors
type MyHandlers struct {
	repo *connector.DbPool
}

// NewHandlers instantiates myHandler
func NewHandlers(p *connector.DbPool) (*MyHandlers, error) {

	nh := MyHandlers{repo: p}

	return &nh, nil
}

// ListUsers lists all users
func (h *MyHandlers) ListUsers(w http.ResponseWriter, r *http.Request) {

	users, err := h.repo.ListAllUsers()
	if err != nil {

		w.WriteHeader(http.StatusNotImplemented)
	}

	fmt.Println(users)
	w.WriteHeader(http.StatusOK)

	// w.Write([]byte(users))
}

// ListDevices lists all users
func (h *MyHandlers) ListDevices(w http.ResponseWriter, r *http.Request) {

	devlist, err := h.repo.ListAllDevices()
	if err != nil {

		w.WriteHeader(http.StatusNotImplemented)
	}

	reqBodyBytes := new(bytes.Buffer)
	json.NewEncoder(reqBodyBytes).Encode(devlist)

	w.WriteHeader(http.StatusOK)

	w.Write(reqBodyBytes.Bytes())
}

// AddDevice adds a device to database
func (h *MyHandlers) AddDevice(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)

}

// NewUser to check input
type NewUser struct {
	Name     string
	Email    string
	Mobileno string
	Password string
}

// AddUser adds a user to database
func (h *MyHandlers) AddUser(w http.ResponseWriter, r *http.Request) {

	newuser := NewUser{}
	err := json.NewDecoder(r.Body).Decode(&newuser)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)

}
