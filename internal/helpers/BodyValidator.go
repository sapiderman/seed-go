package helpers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

//ValidateInput validates inputs against a struct tag
func ValidateInput(ctx context.Context, r *http.Request, mystruct interface{}) error {

	mystruct = validator.New()
	err := json.NewDecoder(r.Body).Decode(&mystruct)
	if err != nil {
		return err
	}

	err = validate.Struct(mystruct)
	if err != nil {
		return err
	}

	return nil
}
