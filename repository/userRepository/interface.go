package userRepository

import "github.com/fbFideles/fin-tracker/models/userModel"

type IUser interface {
	RegisterUser(req *userModel.ReqUserSingUp) (err error)
}
