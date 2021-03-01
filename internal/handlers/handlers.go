package handlers

import (
	"github.com/sapiderman/seed-go/internal/connector"
	log "github.com/sirupsen/logrus"
)

// handler loggin
var hLog = log.WithField("module", "handlers")

// Handlers wraps all needed connectors
type Handlers struct {
	// repo *connector.DbPool
	repo *connector.PgxPool
}

// NewHandlers instantiates myHandler
// func NewHandlers(p *connector.DbPool) *Handlers {
func NewHandlers(p *connector.PgxPool) *Handlers {
	return &Handlers{repo: p}
}
