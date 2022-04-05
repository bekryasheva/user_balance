package database

import (
	"database/sql"
	"user_balance/internal/app/model"
)

func GetBalance(db *sql.DB, user *model.UserBalance) error {
	err := db.QueryRow("SELECT balance FROM users where id = $1", user.ID).Scan(&user.Balance)
	if err != nil {
		return err
	}

	return nil
}
