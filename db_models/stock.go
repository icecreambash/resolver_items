package db_models

import "time"

type Stock struct {
	ID          int64     `json:"id" db:"id"`
	Name        string    `json:"name" db:"name"`
	NameInChess string    `json:"chess_name" db:"name_in_chess"`
	Cef         float64   `json:"cef" db:"cef"`
	Desc        string    `json:"description" db:"description"`
	IconID      int64     `json:"icon_id" db:"icon_id"`
	IconColor   string    `json:"icon_color" db:"icon_color"`
	IsPublished bool      `json:"is_published" db:"is_published"`
	Format      int       `json:"format_id" db:"format_id"`
	Type        int       `json:"type_id" db:"type_id"`
	Public      bool      `json:"public" db:"public"`
	Affect      bool      `json:"affect_price" db:"affect_price"`
	Summed      bool      `json:"summed" db:"summed"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
	AllJK       bool      `json:"all_jk" db:"all_jk"`
}
