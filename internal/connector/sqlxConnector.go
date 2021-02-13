package connector

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/sapiderman/seed-go/internal/config"
	log "github.com/sirupsen/logrus"
)

// User model for storing user table rows
type User struct {
	id        string
	createdAt string
	updatedat string
	deletedAt string
	username  string
	phone     string
	email     string
	password  string
	pin       int
	device    string
}

// Device model for storing device table rows
type Device struct {
	id          string
	createdAt   string
	updatedAt   string
	deletedAt   string
	phoneBrand  string
	phoneModel  string
	year        string
	pushNotifID string
	deviceID    string
}

const (
	// DropAllTblSQL drops all table
	DropAllTblSQL = `DROP TABLE IF EXISTS users, devices;`

	// CreateTblUsersSQL creates user table
	CreateTblUsersSQL = `CREATE TABLE IF NOT EXISTS users (
		id INT PRIMARY KEY,
		created_at TIMESTAMPTZ,
		updated_at TIMESTAMPTZ,
		deleted_at TIMESTAMPTZ,
		username VARCHAR(255) UNIQUE NOT NULL,
		phone VARCHAR UNIQUE NOT NULL, 
		email VARCHAR(100) UNIQUE NOT NULL ,
		password VARCHAR(255) NOT NULL,
		pin INT,
		devices INT REFERENCES device(id)
		);`

	// CreateTblDevicesSQL creates device table
	CreateTblDevicesSQL = `CREATE TABLE IF NOT EXISTS devices (
		id INT PRIMARY KEY,
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
	SelectAllDeviceSQL = `SELECT * from device;`

	// InsertUserSQL adds a user
	InsertUserSQL = `INSERT INTO users (username phone email password pin device) VALUES ($1 $2 $3 $4 $5 $6);`

	// InsertDeviceSQL adds a device
	InsertDeviceSQL = `INSERT INTO devices (phone_brannd phone_model push_id device_id) VALUES ($1 $2 $3 $4);`
)

var (
	//DB instance
	db *sqlx.DB
)

// InitializeDBInstance create db instance
func InitializeDBInstance() error {

	psgqlConnectStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s",
		config.Get("psql.host"), config.Get("psql.port"), config.Get("psql.user"), config.Get("psql.pass"), config.Get("psql.dbname"))
	db, err := sqlx.Connect("postgress", psgqlConnectStr)
	if err != nil {

		log.Error("Connection to database error: ", err)
		return err
	}
	defer db.Close()

	// ensure connection works: sqlx already pings db
	// err = db.Ping()
	// if err != nil {
	// 	log.Error("Ping to database error: ", err)
	// 	return err
	// }

	log.Info("db connection and ping successful...")

	return nil
}

// DropAllTables initializes the
func DropAllTables(db *sqlx.DB) error {

	_, err := db.Exec(DropAllTblSQL)
	if err != nil {
		log.Warn(err)
		return err
	}

	return nil
}

// CreateAllTables initializes the
func CreateAllTables(db *sqlx.DB) error {

	_, err := db.Exec(CreateTblUsersSQL)
	if err != nil {
		log.Warn(err)
		return err
	}

	_, err = db.Exec(CreateTblDevicesSQL)
	if err != nil {
		log.Warn(err)
		return err
	}

	return nil
}

// ListAllUsers list all users
func ListAllUsers() ([]User, error) {

	users := []User{}

	//_ , err := db.Exec(SelectAllUserSQL)
	err := db.Select(&users, SelectAllUserSQL)
	if err != nil {
		log.Warn(err)
		return nil, err
	}

	log.Info("got it:", users)
	return users, nil
}

// InsertUser inserts a single user to the database
func InsertUser(users *User) error {

	_, err := db.NamedExec(`INSERT INTO user () VALUES ()`, users)
	if err != nil {
		log.Warn(err)
		return err
	}

	return nil
}

// ListAllDevices list all devices
func ListAllDevices() ([]Device, error) {

	devices := []Device{}

	//_ , err := db.Exec(SelectAllUserSQL)
	err := db.Select(&devices, SelectAllDeviceSQL)
	if err != nil {
		log.Warn(err)
		return nil, err
	}

	return devices, nil
}

// InsertDevice a record into device table
func InsertDevice(dev Device) error {

	//_ , err := db.Exec(SelectAllUserSQL)
	// err := db.Select(&devices, SelectAllDeviceSQL)
	// if err != nil {
	// 	log.Warn(err)
	// 	return  err
	// }

	return nil
}
