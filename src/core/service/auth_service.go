package service

import (
	"database/sql"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/normatov07/auth-service/common/token"
	"github.com/normatov07/auth-service/common/utils"
	"github.com/normatov07/auth-service/core/action"
	"github.com/normatov07/auth-service/core/app_errors"
	"github.com/normatov07/auth-service/core/model"
	"github.com/normatov07/auth-service/core/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo repository.UserRepo
}

func GetUserService(repo repository.UserRepo) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) Login(atn action.Login) (string, error) {
	user, err := s.repo.GetUserByLogin(atn.Login)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", app_errors.NewAppErr(app_errors.USER_NOT_FOUND, atn.Login)
		}
		return "", app_errors.NewAppErr(app_errors.USER_NOT_FOUND, atn.Login)
	}
	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(atn.Password)); err != nil {
		return "", app_errors.NewAppErr(app_errors.LOGIN_ERROR)
	}

	tMaker, err := token.NewPasetoMaker()
	if err != nil {
		return "", app_errors.NewAppErr(app_errors.USER_NOT_FOUND, atn.Login)
	}

	return tMaker.CreateToken(user)
}

func (s *UserService) Register(atn action.Register) (string, error) {
	_, err := s.repo.GetUserByLogin(atn.Login)
	if err == nil {
		return "", app_errors.NewAppErr(app_errors.LOGIN_UNIQUE)
	} else if err != sql.ErrNoRows {
		log.Printf("register:check %v", err)
		return "", app_errors.NewAppErr(app_errors.SERVER_ERROR)
	}
	hash, _ := utils.HashPassword(atn.Password)

	model := model.UserModel{
		ID:        uuid.New(),
		Login:     atn.Login,
		Password:  hash,
		FirstName: atn.FirstName,
		LastName:  atn.LastName,
		Address:   atn.Address,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err = s.repo.CreateUser(model)
	if err != nil {
		log.Printf("register:create %v", err)
		return "", app_errors.NewAppErr(app_errors.SERVER_ERROR)
	}

	tMaker, err := token.NewPasetoMaker()
	if err != nil {
		return "", app_errors.NewAppErr(app_errors.USER_NOT_FOUND, atn.Login)
	}

	return tMaker.CreateToken(model)
}
