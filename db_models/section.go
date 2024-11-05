package db_models

import "github.com/google/uuid"

type Section struct {
	ID      uuid.UUID `json:"id" gorm:"type:uuid;default:uuid_generate_v4();" db:"id"`
	ChessID uuid.UUID `json:"chess_id" gorm:"type:uuid;default:uuid_generate_v4();" db:"chess_id"`
	IDx     int64     `json:"idx" gorm:"type:" db:"idx"`
	Title   string    `json:"title" gorm:"type:" db:"title"`
}
