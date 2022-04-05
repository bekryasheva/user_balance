package database

import (
	"database/sql"
	"errors"
	"fmt"
	"user_balance/internal/app/pkg"
)

func Withdraw(tx *sql.Tx, id *int, amount *float64) error {
	var balance float64

	err := tx.QueryRow("SELECT balance FROM users WHERE id = $1", id).Scan(&balance)
	if errors.Is(err, sql.ErrNoRows) {
		return pkg.ErrNoUser
	}

	if err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to get user balance %v", err)
	}

	if balance - *amount < 0 {
		return pkg.ErrInsufficientFunds
	}

	_, err = tx.Exec("UPDATE users SET balance = balance - $1 WHERE id = $2", amount, id)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("error occurred while updating user balance %v", err)
	}

	return nil
}
