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

func (r *repository) RegisterUser(req *userModel.ReqUserSingUp) (err error) {
	return r.infra.RegisterUser(req)
}
