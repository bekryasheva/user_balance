package handlers

import (
	"database/sql"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"net/http"
	"user_balance/internal/app/database"
	"user_balance/internal/app/pkg"
)

func DepositHandler(db *sql.DB, log *zap.Logger) echo.HandlerFunc {
	return func(c echo.Context) error {

		d, err := GetBodyDataUser(c)
		if err != nil {
			log.Error("failed to get request body data", zap.Error(err))
			return echo.NewHTTPError(http.StatusBadRequest, pkg.ErrInvalidInput.Error())
		}

		tx, err := db.Begin()
		if err != nil {
			log.Error("failed to start a transaction", zap.Error(err))
			return echo.ErrInternalServerError
		}

		err = database.Deposit(tx, &d.Id, &d.Amount)
		if err != nil {
			log.Error("failed to deposit funds DB", zap.Error(err))
			return echo.ErrInternalServerError
		}

		err = d.Save("deposit", tx)
		if err != nil {
			tx.Rollback()
			log.Error("failed to save DepositTransaction", zap.Error(err))
			return echo.ErrInternalServerError
		}

		tx.Commit()

		return c.NoContent(http.StatusOK)
	}
}

func GetBodyDataUser(c echo.Context) (*database.DepositWithdrawalTransaction, error) {
	var d database.DepositWithdrawalTransaction

	decoder := json.NewDecoder(c.Request().Body)

	decoder.DisallowUnknownFields()

	err := decoder.Decode(&d)
	if err != nil {
		return nil, err
	}

	if d.Amount <= 0 || d.Id == 0{
		return nil, pkg.ErrInvalidInput
	}

	return &d, nil
}
