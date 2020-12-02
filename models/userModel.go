package models

import "time"

// User defines a user structure for usage in the model layer
type User struct {
	ID        *int64     `json:"id,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	Name      *string    `json:"name,omitempty"`
	Email     *string    `json:"email,omitempty"`
	Password  *string    `json:"password,omitempty"`
}

func (u *User) GetAllUsers() {

}

func (u *User) RegisterUser() {

}
