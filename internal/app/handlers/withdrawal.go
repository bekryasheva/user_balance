package handlers

import (
	"database/sql"
	"errors"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"net/http"
	"user_balance/internal/app/database"
	"user_balance/internal/app/pkg"
)

func WithdrawalHandler(db *sql.DB, log *zap.Logger) echo.HandlerFunc {
	return func(c echo.Context) error {

		w, err := GetBodyDataUser(c)
		if err != nil {
			log.Error("failed to get request body data")
			return echo.NewHTTPError(http.StatusBadRequest, pkg.ErrInvalidInput.Error())
		}

		tx, err := db.Begin()
		if err != nil {
			log.Error("failed to start a transaction")
			return echo.ErrInternalServerError
		}

		err = database.Withdraw(tx, &w.Id, &w.Amount)
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

		err = w.Save("withdrawal", tx)
		if err != nil {
			tx.Rollback()
			log.Error("failed to save WithdrawalTransaction")
			return echo.ErrInternalServerError
		}

		tx.Commit()

		return c.NoContent(http.StatusOK)
	}
}
