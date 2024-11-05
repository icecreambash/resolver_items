package db_models

import "github.com/google/uuid"

type ChessEntrance struct {
	ID     uuid.UUID `json:"id" gorm:"type:uuid;default:uuid_generate_v4();" db:"id"`
	Title  string    `json:"title" db:"title"`
	IDx    int       `json:"idx" db:"idx"`
	IsLazy bool      `json:"is_lazy" db:"is_lazy"`
}
