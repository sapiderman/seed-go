package handlers

import (
	"fmt"
	"net/http"

	"encoding/json"

	"github.com/sapiderman/seed-go/internal/connector"
	"github.com/sapiderman/seed-go/internal/helpers"
)

// ListUsers lists all users
func (h *Handlers) ListUsers(w http.ResponseWriter, r *http.Request) {

	users, err := h.repo.ListAllUsers(r.Context())
	if err != nil {

		w.WriteHeader(http.StatusNotImplemented)
	}

	fmt.Println(users)
	w.WriteHeader(http.StatusOK)

	// w.Write([]byte(users))
}

// ListDevices lists all users
func (h *Handlers) ListDevices(w http.ResponseWriter, r *http.Request) {

	devlist, err := h.repo.ListAllDevices(r.Context())
	if err != nil {

		w.WriteHeader(http.StatusNotImplemented)
	}
	helpers.HTTPResponseBuilder(r.Context(), w, r, http.StatusOK, "List Devices", devlist)
}

// AddDevice adds a device to database
func (h *Handlers) AddDevice(w http.ResponseWriter, r *http.Request) {
	logf := hLog.WithField("fn", "AddDevice()")

	newDevice := connector.Device{}
	err := json.NewDecoder(r.Body).Decode(&newDevice)
	if err != nil {
		logf.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = helpers.ValidateInput(r.Context(), newDevice)
	if err != nil {
		logf.Error(err)
		return
	}

	// err = h.repo.InsertDevice(&newDevice)
	// if err != nil {
	// 	logf.Error(err)
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	return
	// }
}

// AddUser registers new users
func (h *Handlers) AddUser(w http.ResponseWriter, r *http.Request) {
	logf := hLog.WithField("fn", "AddUser()")

	newuser := connector.NewUser{}
	err := json.NewDecoder(r.Body).Decode(&newuser)
	if err != nil {
		logf.Error(err)
		helpers.HTTPResponseBuilder(r.Context(), w, r, http.StatusBadRequest, "Bad payload", nil)
		return
	}

	err = helpers.ValidateInput(r.Context(), newuser)
	if err != nil {
		logf.Error(err)
		helpers.HTTPResponseBuilder(r.Context(), w, r, http.StatusBadRequest, "Missing or wrong parameters", nil)
		return
	}

	// err = h.repo.InsertUser(&newuser)
	// if err != nil {
	// 	logf.Error(err)
	// 	helpers.HTTPResponseBuilder(r.Context(), w, r, http.StatusInternalServerError, "", nil)
	// 	return
	// }

}
