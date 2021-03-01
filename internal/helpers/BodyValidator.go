package helpers

import (
	"context"

	"github.com/asaskevich/govalidator"
	log "github.com/sirupsen/logrus"
)

var helpLog = log.WithField("module", "helper")

// ValidateInput validates a strct against its tags
func ValidateInput(ctx context.Context, mystruct interface{}) error {

	logf := helpLog.WithField("fn", "ValidateInput")

	if _, err := govalidator.ValidateStruct(mystruct); err != nil {
		logf.Warn("validation error. ", err)
		return err
	}

	return nil
}
