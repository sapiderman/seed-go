package model

import (
	"database/sql"
	"errors"
)

type user struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Email float64 `json:"email"`
}

func (u *user) createUser(db *sql.DB) error {

	return errors.New("Not implemented")
}

func (u *user) listUser(db *sql.DB) error {

	return errors.New("Not implemented")
}

func (u *user) findUser(db *sql.DB) error {

	return errors.New("Not implemented")
}

func (u *user) deleteUser(db *sql.DB) error {

	return errors.New("Not implemented")
}
