package models

// User model for storing user table rows
type User struct {
	ID        string `db:"id"`
	CreatedAt string `db:"created_at"`
	UpdatedAt string `db:"update_at"`
	DeletedAt string `db:"deleted_at"`
	Username  string `db:"username"`
	Phone     string `db:"phone"`
	Email     string `db:"email"`
	Password  string `db:"password"`
	Pin       int    `db:"pin"`
	Device    string `db:"device"`
}

// Device model for storing device table rows
type Device struct {
	ID         string `db:"id"`
	CreatedAt  string `db:"created_at"`
	UpdatedAt  string `db:"updated_at"`
	DeletedAt  string `db:"deleted_at"`
	PhoneBrand string `db:"phone_brand"`
	PhoneModel string `db:"phone_model"`
	PushID     string `db:"push_id"`
	DeviceID   string `db:"device_id"`
}

// NewUser to check input
type NewUser struct {
	Name     string `db:"name"`
	Email    string `db:"email"`
	Mobile   string `db:"mobile_no"`
	Password string `db:"password"`
	DeviceID string `db:"DeviceID"`
}
