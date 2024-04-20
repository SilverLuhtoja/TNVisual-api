// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package database

import (
	"database/sql"
	"encoding/json"

	"github.com/google/uuid"
)

type Project struct {
	ID          uuid.UUID
	Name        string
	Content     json.RawMessage
	Description sql.NullString
}

type User struct {
	ID       uuid.UUID
	Username string
	Password string
	ApiKey   string
}
