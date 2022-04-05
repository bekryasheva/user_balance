package handlers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"net/http"
	"user_balance/internal/app/database"
	"user_balance/internal/app/pkg"
)

func TransferHandler(db *sql.DB, log *zap.Logger) echo.HandlerFunc {
	return func(c echo.Context) error {
		tr, err := GetBodyDataTransfer(c)
		if err != nil {
			log.Error("failed to get request body data")
			return echo.NewHTTPError(http.StatusBadRequest, pkg.ErrInvalidInput.Error())
		}

		tx, err := db.Begin()
		if err != nil {
			log.Error("failed to start a transaction")
			return echo.ErrInternalServerError
		}

		err = database.Withdraw(tx, &tr.FID, &tr.Amount)
		if err != nil {
			log.Error("failed to withdraw funds", zap.Error(err))
			if errors.Is(err, pkg.ErrNoUser) {
				return echo.NewHTTPError(http.StatusPreconditionFailed, pkg.ErrNoUser.Error())
			}

			if errors.Is(err, pkg.ErrInsufficientFunds) {
				return echo.NewHTTPError(http.StatusPreconditionFailed, pkg.ErrInsufficientFunds.Error())
			}

			return echo.ErrInternalServerError
		}

		err = database.Deposit(tx, &tr.TID, &tr.Amount)
		if err != nil {
			log.Error("failed to deposit funds DB", zap.Error(err))
			return echo.ErrInternalServerError
		}

		err = tr.Save(tx)
		if err != nil {
			tx.Rollback()
			log.Error("failed to save TransferTransaction")
			return echo.ErrInternalServerError
		}

		tx.Commit()

		return c.NoContent(http.StatusOK)
	}
}

func GetBodyDataTransfer(c echo.Context) (*database.TransferTransaction, error) {
	var transaction *database.TransferTransaction

	decoder := json.NewDecoder(c.Request().Body)

	decoder.DisallowUnknownFields()

	err := decoder.Decode(&transaction)
	if err != nil {
		return nil, err
	}

	if transaction.Amount <= 0 || transaction.TID == 0 || transaction.FID == 0 {
		return nil, pkg.ErrInvalidInput
	}

	return transaction, nil
}
