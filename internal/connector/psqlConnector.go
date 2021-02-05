package connector

import (
	"database/sql"
	"fmt"

	// use postgress init
	_ "github.com/lib/pq"
	"github.com/sapiderman/seed-go/internal/config"
	log "github.com/sirupsen/logrus"
)

const (
	// DropAllTblSQL drops all table
	DropAllTblSQL = `DROP TABLE IF EXISTS users, device;`

	// CreateTblUser creates user table
	CreateTblUser = `CREATE TABLE IF NOT EXISTS users (
		id INT PRIMARY KEY,
		created_at TIMESTAMPTZ,
		updated_at TIMESTAMPTZ,
		deleted_at TIMESTAMPTZ,
		username VARCHAR(255) UNIQUE NOT NULL,
		phone VARCHAR UNIQUE NOT NULL, 
		email VARCHAR(100) UNIQUE NOT NULL ,
		password VARCHAR(255) NOT NULL,
		pin INT,
		device INT REFERENCES device(id)
		);`

	// CreateTblDevice creates device table
	CreateTblDevice = `CREATE TABLE IF NOT EXISTS device (
		id INT PRIMARY KEY,
		created_at TIMESTAMPTZ,
		updated_at TIMESTAMPTZ,
		deleted_at TIMESTAMPTZ,
		phone_brand VARCHAR(255) NOT NULL,
		phone_model VARCHAR(100) NOT NULL, 
		year VARCHAR(100) NOT NULL ,
		push_notif_id VARCHAR,
		device_id VARCHAR
);`
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
