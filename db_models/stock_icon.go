package db_models

type Icons struct {
	ID   int    `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
	Slug string `db:"slug" json:"slug"`
	SVG  string `db:"svg" json:"svg"`
	Tag  string `db:"tag" json:"tag"`
}
