package model

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type UserModel struct {
	ID        uuid.UUID
	Login     string
	Password  string
	FirstName string
	LastName  string
	Address   sql.NullString
	CreatedAt time.Time
	UpdatedAt time.Time
}
