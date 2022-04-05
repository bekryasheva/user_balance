package database

import (
	"database/sql"
	"time"
)

type DepositWithdrawalTransaction struct {
	Id     int `json:"id"`
	Amount float64 `json:"amount"`
}

func (d DepositWithdrawalTransaction) Save(transactionType string, tx *sql.Tx) error {
	_, err := tx.Exec("INSERT INTO transactions (user_id, type, date, amount) VALUES ($1, $2, $3, $4)", d.Id, transactionType, time.Now(), d.Amount)
	if err != nil {
		return err
	}
	return nil
}

type TransferTransaction struct {
	FID      int  `json:"fid"`
	TID      int  `json:"tid"`
	Amount   float64 `json:"amount"`
}

func (t TransferTransaction) Save(tx *sql.Tx) error {
	_, err := tx.Exec("INSERT INTO transactions (user_id, type, date, amount, from_id, to_id) VALUES ($1, $2, $3, $4, $5, $6)", t.TID, "transfer", time.Now(), t.Amount, t.FID, t.TID)
	if err != nil {
		return err
	}

	_, err = tx.Exec("INSERT INTO transactions (user_id, type, date, amount, from_id, to_id) VALUES ($1, $2, $3, $4, $5, $6)", t.FID, "transfer", time.Now(), t.Amount, t.FID, t.TID)
	if err != nil {
		return err
	}
	return nil
}


