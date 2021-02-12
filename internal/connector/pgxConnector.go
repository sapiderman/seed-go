package connector

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4"

	"github.com/sapiderman/seed-go/internal/config"
	log "github.com/sirupsen/logrus"
)

var (
	conn *pgx.Conn
)

// InitPgxDBInstance initialize connection pol
func InitPgxDBInstance() error {
	pgxConnectorStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s",
		config.Get("psql.host"), config.Get("psql.port"), config.Get("psql.user"), config.Get("psql.pass"), config.Get("psql.dbname"))

	ctx := context.Background()
	_, err := pgx.Connect(ctx, pgxConnectorStr)
	if err != nil {
		log.Error("error: ", err)
		return nil
	}

	return nil
}
