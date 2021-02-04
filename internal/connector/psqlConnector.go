package connector

import (
	"database/sql"
	"fmt"

	// use postgress init
	_ "github.com/lib/pq"
	"github.com/sapiderman/seed-go/internal/config"
	log "github.com/sirupsen/logrus"
)

// GetDBInstance create db instance
func GetDBInstance() *sql.DB {

	psgqlConnectStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s",
		config.Get("psql.host"), config.Get("psql.port"), config.Get("psql.user"), config.Get("psql.pass"), config.Get("psql.dbname"))
	db, err := sql.Open("postgres", psgqlConnectStr)
	if err != nil {
		log.Error(err)
	}
	defer db.Close()

	// ensure connection
	err = db.Ping()
	if err != nil {
		log.Error(err)
	}

	return db
}
