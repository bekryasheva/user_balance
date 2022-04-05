package model

type UserBalance struct {
	ID       int64  `json:"id"`
	Balance  float64 `json:"balance"`
	Currency string  `json:"currency"`
}


