package database

import (
	"database/sql"
	"fmt"
)

func Deposit(tx *sql.Tx, id *int, amount *float64) error {
	res, err := tx.Exec("UPDATE users SET balance = balance + $1 WHERE id = $2", amount, id)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("error occurred while updating user balance: %v", err)
	}

	updates, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows updated: %v", err)
	}

	if updates > 0 {
		return nil
	}

	_, err = tx.Exec("INSERT INTO users (id, balance) VALUES ($1, $2)", id, amount)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to add user and deposit: %v", err)
	}

	return nil
}

