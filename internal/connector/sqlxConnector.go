package connector

import (
	"fmt"

	//postgress import
	_ "github.com/lib/pq"

	"github.com/jmoiron/sqlx"
	"github.com/sapiderman/seed-go/internal/config"
	log "github.com/sirupsen/logrus"
)

const (
	// DropAllTblSQL drops all table
	DropAllTblSQL = `DROP TABLE IF EXISTS users, devices;`

	// CreateTblUsersSQL creates user table
	CreateTblUsersSQL = `CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		created_at TIMESTAMPTZ,
		updated_at TIMESTAMPTZ,
		deleted_at TIMESTAMPTZ,
		username VARCHAR(255) UNIQUE NOT NULL,
		phone VARCHAR UNIQUE NOT NULL, 
		email VARCHAR(100) UNIQUE NOT NULL ,
		password VARCHAR(255) NOT NULL,
		pin INT,
		devices INT REFERENCES devices(id)
		);`

	// CreateTblDevicesSQL creates device table
	CreateTblDevicesSQL = `CREATE TABLE IF NOT EXISTS devices (
		id SERIAL PRIMARY KEY,
		created_at TIMESTAMPTZ,
		updated_at TIMESTAMPTZ,
		deleted_at TIMESTAMPTZ,
		phone_brand VARCHAR(255) NOT NULL,
		phone_model VARCHAR(100) NOT NULL, 
		push_id VARCHAR,
		device_id VARCHAR
		);`
	// SelectAllUserSQL queries user table
	SelectAllUserSQL = `SELECT * from users;`

	// SelectAllDeviceSQL queries device table
	SelectAllDeviceSQL = `SELECT * from devices;`

	// InsertUserSQL adds a user
	InsertUserSQL = `INSERT INTO users 
	(username, phone, email, password, pin, device) 
	VALUES (:username, :phone, :email, :password, :pin, :device);`

	// InsertDeviceSQL adds a devics)
	InsertDeviceSQL = `INSERT INTO devices (created_at, updated_at, deleted_at, phone_brand, phone_model, push_id, device_id)
	  VALUES (:created_at, :updated_at, :deleted_at, :phone_brand, :phone_model, :push_id, :device_id);`
)

//DbPool struct wraps the db instance
type DbPool struct {
	Db *sqlx.DB
}

var (
	sqlxLog = log.WithField("module", "sqlx")
)

// SqlxNewInstance create DbPool instance
func SqlxNewInstance() (*DbPool, error) {
	logf := sqlxLog.WithField("func", "InitializeDBInstance")

	psgqlConnectStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.Get("psql.host"), config.Get("psql.port"), config.Get("psql.user"), config.Get("psql.pass"), config.Get("psql.dbname"))
	db, err := sqlx.Connect("postgres", psgqlConnectStr)
	if err != nil {
		logf.Error("Connection to database error: ", err)
		return nil, err
	}
	// defer db.Close()
	logf.Info("db connection and ping successful...")

	p := DbPool{Db: db}
	return &p, nil
}

// CloseConnection closes connetion
func (p *DbPool) CloseConnection() error {
	p.Db.Close()
	sqlxLog.WithField("func", "CloseConnection").Info("Closing database connection")
	return nil
}

// DropAllTables initializes the
func (p *DbPool) DropAllTables() error {

	logf := sqlxLog.WithField("func", "DropAllTables")
	_, err := p.Db.Exec(DropAllTblSQL)
	if err != nil {
		logf.Warn(err)
		return err
	}

	return nil
}

// CreateAllTables initializes the
func (p *DbPool) CreateAllTables() error {

	_, err := p.Db.Exec(CreateTblDevicesSQL)
	if err != nil {
		log.Warn(err)
		return err
	}

	_, err = p.Db.Exec(CreateTblUsersSQL)
	if err != nil {
		log.Warn(err)
		return err
	}

	return nil
}

// ListAllUsers list all users
func (p *DbPool) ListAllUsers() ([]User, error) {

	logf := sqlxLog.WithField("func", "ListAllUsers")
	users := []User{}

	//_ , err := db.Exec(SelectAllUserSQL)
	// err := db.Select(&users, SelectAllUserSQL)
	err := p.Db.Select(&users, `SELECT * from devices;`)
	if err != nil {
		logf.Warn(err)
		return nil, err
	}

	logf.Info("got it:", users)
	return users, nil
}

// InsertUser inserts a single user to the database
func (p *DbPool) InsertUser(users *NewUser) error {
	logf := sqlxLog.WithField("func", "InsertUser")

	_, err := p.Db.NamedExec(InsertUserSQL, users)
	if err != nil {
		logf.Error(err)
		return err
	}

	return nil
}

// ListAllDevices list all devices
func (p *DbPool) ListAllDevices() ([]Device, error) {
	logf := sqlxLog.WithField("func", "ListAllDevices")

	devices := []Device{}

	//_ , err := db.Exec(SelectAllUserSQL)
	err := p.Db.Select(&devices, "SELECT * from devices;")
	if err != nil {
		logf.Warn(err)
		return nil, err
	}

	return devices, nil
}

// InsertDevice a record into device table
func (p *DbPool) InsertDevice(d *Device) error {
	logf := sqlxLog.WithField("func", "ListAllInsertDeviceDevices")

	res, err := p.Db.NamedExec(InsertDeviceSQL, d)
	if err != nil {
		logf.Warn(err)
		return err
	}

	logf.Info("success: ", res)
	return nil
}
