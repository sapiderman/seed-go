package models

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
