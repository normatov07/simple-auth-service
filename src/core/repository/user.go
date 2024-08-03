package repository

import (
	"github.com/normatov07/auth-service/core/model"
)

type UserRepo interface {
	GetUserByLogin(login string) (model.UserModel, error)
	CreateUser(model model.UserModel) error
}
