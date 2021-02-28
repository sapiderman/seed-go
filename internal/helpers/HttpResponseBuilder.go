package helpers

import (
	"context"
	"encoding/json"
	"net/http"
)

// HTTPResponseBuilder builds the reponse headers and payloads
func HTTPResponseBuilder(ctx context.Context, w http.ResponseWriter, r *http.Request, httpStatus int, message string, data interface{}) {

	resp := make(map[string]interface{})
	resp["data"] = data
	if len(message) > 0 {
		resp["message"] = message
	}

	switch httpStatus {
	case http.StatusOK:
		resp["status"] = "OK"
	case http.StatusBadRequest:
		resp["status"] = "Bad Request"
	case http.StatusInternalServerError:
		resp["status"] = "Internal Server Error"
	case http.StatusForbidden:
		resp["status"] = "Forbidden"
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatus)
	json.NewEncoder(w).Encode(resp)
}
