package db_models

import "github.com/google/uuid"

type FloorItems struct {
	ID       uuid.UUID `json:"id" db:"id"`
	FloorID  uuid.UUID `json:"floor_id" db:"floor_id"`
	IDx      int64     `json:"idx" db:"idx"`
	ModelID  uuid.UUID `json:"model_id" db:"model_id"`
	IsGhost  bool      `json:"is_ghost" db:"is_ghost"`
	IsVanish bool      `json:"is_vanish" db:"is_vanish"`
	Slave    uuid.UUID `json:"slave" db:"slave" `
}
