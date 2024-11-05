package db_models

import (
	"github.com/google/uuid"
	"time"
)

type Chess struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;default:uuid_generate_v4();"`
	ModelID   string    `json:"model_id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}
