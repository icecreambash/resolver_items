package db_models

import "github.com/google/uuid"

type SectionSyncEntrance struct {
	ID         int64     `json:"id" db:"id"`
	SectionID  uuid.UUID `json:"section_id" db:"section_id" gorm:"type:uuid;default:uuid_generate_v4();"`
	EntranceID uuid.UUID `json:"entrance_id" db:"entrance_id" gorm:"type:uuid;default:uuid_generate_v4();"`
}
