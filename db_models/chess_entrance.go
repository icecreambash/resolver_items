package db_models

import "github.com/google/uuid"

type ChessSyncEntrance struct {
	ID         int
	ChessID    string
	EntranceID uuid.UUID `json:"entrance_id" gorm:"type:uuid;default:uuid_generate_v4()"`
}
