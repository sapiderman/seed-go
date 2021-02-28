package connector

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4"

	"github.com/sapiderman/seed-go/internal/config"
	log "github.com/sirupsen/logrus"
)

// PgxConn is the Pgx connection pool
type PgxConn struct {
	conn *pgx.Conn
}

var (
	pgxLog = log.WithField("module", "pgx")
)

// NewConnection initialize connection pol
func NewConnection() (*PgxConn, error) {
	pgxConnectorStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s",
		config.Get("psql.host"), config.Get("psql.port"), config.Get("psql.user"), config.Get("psql.pass"), config.Get("psql.dbname"))

	ctx := context.Background()
	conn, err := pgx.Connect(ctx, pgxConnectorStr)
	if err != nil {
		log.Error("error: ", err)
		return nil, nil
	}

	pg := PgxConn{conn: conn}
	return &pg, nil
}

// CloseConnection close
func (p *PgxConn) CloseConnection(ctx context.Context) error {

	p.conn.Close(ctx)
	return nil
}
