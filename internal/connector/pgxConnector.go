package connector

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/sapiderman/seed-go/internal/config"
	log "github.com/sirupsen/logrus"
)

// PgxPool is the Pgx connection pool
type PgxPool struct {
	pool *pgxpool.Pool
}

var (
	pgxLog = log.WithField("module", "pgx")
)

// PgxNewConnection initialize connection pol
func PgxNewConnection(ctx context.Context) (*PgxPool, error) {
	logf := pgxLog.WithField("fn", "PgxNewConnection")

	pgxConnectorStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s",
		config.Get("psql.host"), config.Get("psql.port"), config.Get("psql.user"), config.Get("psql.pass"), config.Get("psql.dbname"))

	conn, err := pgxpool.Connect(ctx, pgxConnectorStr)
	if err != nil {
		logf.Error(err)
		return nil, nil
	}

	pg := PgxPool{pool: conn}
	return &pg, nil
}

// PgxCloseConnection close
func (p *PgxPool) PgxCloseConnection(ctx context.Context) error {

	p.pool.Close()
	return nil
}

// InsertUser inserts a single user to the database
func (p *PgxPool) InsertUser(ctx context.Context, users *NewUser) error {
	logf := sqlxLog.WithField("fn", "InsertUser")

	if _, err := p.pool.Exec(ctx, `insert into shortened_urls(id, url) values ($1, $2)	on conflict (id) do update set url=excluded.url`, users); err != nil {
		logf.Error(err)
		return err
	}

	return nil
}

// CreateAllTables queries for a user name
func (p *PgxPool) CreateAllTables(ctx context.Context) ([]Device, error) {
	logf := pgxLog.WithField("fn", "CreateAllTables")

	_, err := p.pool.Exec(ctx, createTblUsersSQL)
	if err != nil {
		logf.Error(err)
		return nil, err
	}
	_, err = p.pool.Exec(ctx, createTblDevicesSQL)
	if err != nil {
		logf.Error(err)
		return nil, err
	}

	devices := []Device{}
	return devices, nil
}

// InsertDevice inserts a single user to the database
func (p *PgxPool) InsertDevice(ctx context.Context, users *NewUser) error {
	// logf := sqlxLog.WithField("fn", "InsertDevice")

	// _, err := p.conn.
	// if err != nil {
	// 	logf.Error(err)
	// 	return err
	// }

	return nil
}

// InsertArticle inserts a single user to the database
func (p *PgxPool) InsertArticle(ctx context.Context, users *NewUser) error {
	// logf := sqlxLog.WithField("fn", "InsertArticle")

	// _, err := p.conn.
	// if err != nil {
	// 	logf.Error(err)
	// 	return err
	// }

	return nil
}

// FindUserName queries for a user name
func (p *PgxPool) FindUserName(ctx context.Context, users *NewUser) error {
	// logf := pgxLog.WithField("fn", "FindUserName")

	return nil
}

// ListAllUsers queries for a user name
func (p *PgxPool) ListAllUsers(ctx context.Context) ([]User, error) {
	// logf := pgxLog.WithField("fn", "ListAllUsers")

	users := []User{}

	return users, nil
}

// ListAllDevices queries for a user name
func (p *PgxPool) ListAllDevices(ctx context.Context) ([]Device, error) {
	logf := pgxLog.WithField("fn", "ListAllDevices")

	_, err := p.pool.Exec(ctx, `SELECT * FROM devices`)
	if err != nil {
		logf.Error(err)
		return nil, err
	}

	devices := []Device{}
	return devices, nil
}

// ListAllArticles queries for a user name
func (p *PgxPool) ListAllArticles(ctx context.Context) error {
	// logf := pgxLog.WithField("fn", "ListAllArticles")

	return nil
}
