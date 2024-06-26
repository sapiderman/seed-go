package connector

// User model for storing user table rows
type User struct {
	ID        string `json:"id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"update_at"`
	DeletedAt string `json:"deleted_at"`
	Username  string `json:"username"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Pin       int    `json:"pin"`
	Device    string `json:"device"`
	Role      string `json:"role"`
}

// Device model for storing device table rows
type Device struct {
	ID         string `json:"id"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
	DeletedAt  string `json:"deleted_at"`
	PhoneBrand string `json:"phone_brand"`
	PhoneModel string `json:"phone_model"`
	PushID     string `json:"push_id"`
	DeviceID   string `json:"device_id"`
}

// NewUser to check input
type NewUser struct {
	Name     string `json:"Name" valid:"required"`
	Email    string `json:"Email" valid:"required,email"`
	Mobile   string `json:"Mobile" valid:"required,numeric"`
	Password string `json:"Password" valid:"required,minstringlength(8)"`
	DeviceID string `json:"DeviceID" valid:"optional"`
}

const (
	// DropAllTblSQL drops all table
	dropAllTblSQL = `DROP TABLE IF EXISTS users, devices;`

	// CreateTblUsersSQL creates user table
	createTblUsersSQL = `CREATE TABLE IF NOT EXISTS users (
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
	createTblDevicesSQL = `CREATE TABLE IF NOT EXISTS devices (
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
	// selectAllUserSQL = `SELECT * from users;`

	// SelectAllDeviceSQL queries device table
	// selectAllDeviceSQL = `SELECT * from devices;`

	// InsertUserSQL adds a user
	insertUserSQL = `INSERT INTO users 
(username, phone, email, password, pin, device) 
VALUES (:username, :phone, :email, :password, :pin, :device);`

	// InsertDeviceSQL adds a devics)
	InsertDeviceSQL = `INSERT INTO devices (created_at, updated_at, deleted_at, phone_brand, phone_model, push_id, device_id)
  VALUES (:created_at, :updated_at, :deleted_at, :phone_brand, :phone_model, :push_id, :device_id);`
)
