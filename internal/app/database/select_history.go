package database

import (
	"database/sql"
	"fmt"
	"user_balance/internal/app/pkg"
)

type Transaction struct {
	Id int `json:"id"`
	UserID          int    `json:"user_id"`
	TransactionType string `json:"type_of_transaction"`
	Date            string `json:"date"`
	Amount float64 `json:"amount"`
	FromID sql.NullInt64 `json:"from_id"`
	ToID sql.NullInt64 `json:"to_id"`
}

func GetHistory(db *sql.DB, id int64, orderBy, sort string, limit, offset int64) ([]Transaction, error) {

	query := fmt.Sprintf("SELECT * FROM transactions WHERE user_id = %d ORDER BY %s %s LIMIT %d OFFSET %d", id, orderBy, sort, limit, offset)
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if !rows.Next() {
		return nil, pkg.ErrNoUser
	}
	var transactions []Transaction

	for rows.Next() {
		var tr Transaction
		err = rows.Scan(&tr.Id, &tr.UserID, &tr.TransactionType, &tr.Date, &tr.Amount, &tr.FromID, &tr.ToID)
		if err != nil {
			return nil, err
		}

		transactions = append(transactions, tr)
	}

	return transactions, nil
}
