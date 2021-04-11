package infraestructure

import (
	"database/sql"

	"github.com/fbFideles/fin-tracker/models/userModel"
)

type UserInfra struct {
	Transaction *sql.Tx
}

func (u *UserInfra) RegisterUser(req *userModel.ReqUserSingUp) (err error) {
	result, err := u.Transaction.Exec(`
		INSERT INTO
		t_user (first_name, last_name, email, password)
		VALUES ($1, $2, $3, $4)
	`, req.FirstName, req.LastName, req.Email, req.Password)

	return
}

func (u *UserInfra) CheckEmail(email *string) (err error) {
	rows, err := u.Transaction.Query(`
		SELECT 1
		FROM t_user
		WHERE email LIKE $1::TEXT
	`, email)

	return
}
