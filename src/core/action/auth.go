package action

import (
	"database/sql"
)

type Login struct {
	Login    string `form:"login" binding:"required"`
	Password string `form:"password" binding:"required"`
}

type Register struct {
	Login     string         `form:"login" binding:"required"`
	Password  string         `form:"password" binding:"required"`
	FirstName string         `form:"first_name" binding:"required"`
	LastName  string         `form:"last_name"  binding:"required"`
	Address   sql.NullString `form:"address"`
}
