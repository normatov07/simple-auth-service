package postgres

import (
	"github.com/normatov07/auth-service/core/model"
)

type UserRepo struct{}

func (r UserRepo) GetUserByLogin(login string) (m model.UserModel, err error) {
	err = conn.QueryRow("SELECT id,login,password,first_name,last_name,address,created_at,updated_at FROM users WHERE login=$1", login).Scan(&m.ID, &m.Login, &m.Password, &m.FirstName, &m.LastName, &m.Address, &m.CreatedAt, &m.UpdatedAt)
	return
}

func (r UserRepo) CreateUser(model model.UserModel) (err error) {
	_, err = conn.Exec("INSERT INTO users (id,login,password,first_name,last_name,address,created_at,updated_at) VALUES ($1,$2,$3,$4,$5,$6,$7,$8)", model.ID, model.Login, model.Password, model.FirstName, model.LastName, model.Address.String, model.CreatedAt, model.UpdatedAt)
	return
}
