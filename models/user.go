package models

import (
	"errors"
	"time"

	"github.com/jorge-dev/ev-book/db"
	hash "github.com/jorge-dev/ev-book/utils"
)

type AuthUser struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password,omitempty" binding:"required"`
}

type User struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name" binding:"required"`
	CreatedAt time.Time `json:"created_at"`
	AuthUser
}

func (u *User) Save() error {
	creationTime := time.Now()
	u.CreatedAt = creationTime
	query := `INSERT INTO users (name, username, email, password, createdAt) VALUES (?, ?, ?, ?, ?)`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		errorMessage := "Error preparing the query to save the user: " + err.Error()
		return errors.New(errorMessage)
	}
	defer stmt.Close()
	hashedPassword, err := hash.HashPassword(u.Password)
	if err != nil {
		return err
	}
	result, err := stmt.Exec(u.Name, u.Username, u.Email, hashedPassword, creationTime)
	if err != nil {
		errorMessage := "Error executing the query to save the user: " + err.Error()
		return errors.New(errorMessage)
	}
	id, err := result.LastInsertId()
	u.ID = id
	return err
}

func (u *AuthUser) ValidateCredentials() error {
	query := `SELECT password FROM users WHERE username = ? OR email = ?`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		errorMessage := "Error preparing the query to get the user: " + err.Error()
		return errors.New(errorMessage)
	}
	defer stmt.Close()
	row := stmt.QueryRow(u.Username, u.Email)

	var retrievePwd string
	err = row.Scan(&retrievePwd)
	if err != nil {
		errorMessage := "invalid credentials. Please check your username or email"
		return errors.New(errorMessage)
	}

	if !hash.ComparePasswords(retrievePwd, u.Password) {
		errorMessage := "invalid credentials. Please check your username or email"
		return errors.New(errorMessage)
	}

	return nil

}
