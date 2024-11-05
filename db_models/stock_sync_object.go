package db_models

type StockSyncObject struct {
	ID        int64  `json:"id" db:"id"`
	StockID   int64  `json:"stock_id" db:"stock_id"`
	ModelType string `json:"model_type" db:"model_type"`
	ModelID   int64  `json:"model_id" db:"model_id"`
}
