package infraestructure

import (
	"database/sql"

	"github.com/fbFideles/fin-tracker/models/userModel"
)

type UserInfra struct {
	Transaction *sql.Tx
}

// Registers a user in the database
func (u *UserInfra) RegisterUser(req *userModel.ReqUserSingUp) (err error) {
	if _, err = u.Transaction.Exec(`
		INSERT INTO
		t_user (first_name, last_name, email, password)
		VALUES ($1, $2, $3, $4)
	`, req.FirstName, req.LastName, req.Email, req.Password); err != nil {
		return
	}

	return
}

// CheckEmail checks if the email passed is on the database
func (u *UserInfra) CheckEmail(email *string) (err error) {
	if err = u.Transaction.QueryRow(`
		SELECT 1
		FROM t_user
		WHERE email LIKE $1::TEXT
	`, email).Scan(); err != nil {
		if err == sql.ErrNoRows {
			return nil
		}

		return
	}

	return
}
