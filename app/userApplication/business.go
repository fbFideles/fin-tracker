package userApplication

import (
	"context"

	"github.com/fbFideles/fin-tracker/database"
	"github.com/fbFideles/fin-tracker/models/userModel"
	"github.com/fbFideles/fin-tracker/repository/userRepository"
	"github.com/fbFideles/fin-tracker/utils"
	"golang.org/x/crypto/bcrypt"
)

func SingUp(ctx context.Context, req *userModel.ReqUserSingUp) (token *string, err error) {
	tx, err := database.NewTransaction(ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	repo := userRepository.NewRepository(tx)

	repo.RegisterUser(req)

	hashedPasswordBytes, err := bcrypt.GenerateFromPassword([]byte(*req.Password), 15)
	if err != nil {
		return nil, err
	}
	req.Password = utils.StringPointer(string(hashedPasswordBytes))

	if err = tx.Commit(); err != nil {
		return nil, err
	}

	return
}
