package db_models

type Room struct {
	ID           int64   `json:"id" db:"id"`
	Rooms        int     `json:"rooms" db:"rooms"`
	AreaFull     float64 `json:"area_full" db:"area_full"`
	PriceFull    float64 `json:"price_full" db:"price_full"`
	PriceUnit    float64 `json:"price_unit" db:"price_unit"`
	NumberObject string  `json:"number_object" db:"number_object"`
	Status       string  `json:"status" db:"status"`
}

func (a Room) GetID() int64 {
	return a.ID
}
