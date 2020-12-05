package models

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/fbFideles/fin-tracker/utils"
	"golang.org/x/crypto/bcrypt"
)

// User defines a user structure for usage in the model layer
type User struct {
	ID        *int64     `json:"id,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	Name      *string    `json:"name,omitempty"`
	Email     *string    `json:"email,omitempty"`
	Password  *string    `json:"password,omitempty"`
}

// GetAllUsers is a method for getting
// all users in the database
func (u *User) GetAllUsers() (err error) {
	return
}

// RegisterUser defines a method for the user
// model to register a given user
func (u *User) RegisterUser(tx *sql.Tx) (err error) {
	hashBytes, err := bcrypt.GenerateFromPassword([]byte(*u.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("Error in hashing password")
	}
	u.Password = utils.StringPointer(string(hashBytes))

	if _, err := tx.Exec(`
		INSERT INTO t_user
			(name, email, password)
		VALUES
			(?, ?, ?)
		`); err != nil {
		return err
	}

	return
}
