package handlers

import (
	"bytes"
	"fmt"
	"net/http"

	"encoding/json"

	"github.com/sapiderman/seed-go/internal/connector"
	"github.com/sapiderman/seed-go/internal/models"
)

// MyHandlers wraps all needed connectors
type MyHandlers struct {
	repo *connector.DbPool
}

// NewHandlers instantiates myHandler
func NewHandlers(p *connector.DbPool) *MyHandlers {
	return &MyHandlers{}
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

// AddUser adds a user to database
func (h *MyHandlers) AddUser(w http.ResponseWriter, r *http.Request) {

	newuser := models.NewUser{}
	err := json.NewDecoder(r.Body).Decode(&newuser)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.repo.InsertUser(&newuser)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

}
