package db_models

import "github.com/google/uuid"

type Floor struct {
	ID            uuid.UUID `json:"id" db:"id"`
	EntranceID    uuid.UUID `json:"entrance_id" db:"entrance_id"`
	IDx           int64     `json:"idx" db:"idx"`
	IsActiveLimit bool      `json:"is_active_limit" db:"is_active_limit"`
	ActiveLimit   int       `json:"active_limit" db:"active_limit"`
}
