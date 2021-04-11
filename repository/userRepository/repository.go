package userRepository

import (
	"database/sql"

	"github.com/fbFideles/fin-tracker/models/userModel"
	"github.com/fbFideles/fin-tracker/repository/userRepository/infraestructure"
)

type repository struct {
	infra *infraestructure.UserInfra
}

func NewRepository(tx *sql.Tx) IUser {
	return &repository{
		infra: &infraestructure.UserInfra{Transaction: tx},
	}
}

// RegisterUser is a data flux manager to build the repository method of registration
func (r *repository) RegisterUser(req *userModel.ReqUserSingUp) (err error) {
	return r.infra.RegisterUser(req)
}

// CheckEmail is a data flux manager to build the repository method of registration
func (r *repository) CheckEmail(email *string) (err error) {
	return r.infra.CheckEmail(email)
}
