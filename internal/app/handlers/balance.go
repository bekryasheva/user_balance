package handlers

import (
	"database/sql"
	"errors"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"net/http"
	"user_balance/internal/app"
	"user_balance/internal/app/database"
	"user_balance/internal/app/exchangerates"
	"user_balance/internal/app/model"
	"user_balance/internal/app/pkg"
)

func BalanceHandler(db *sql.DB, log *zap.Logger, config app.Config) echo.HandlerFunc {
	return func(c echo.Context) error {

		var user model.UserBalance

		err := echo.PathParamsBinder(c).Int64("id", &user.ID).BindError()
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, pkg.ErrInvalidInput.Error())
		}

		err = pkg.ValidateId(user.ID)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, pkg.ErrInvalidInput.Error())
		}

		err = database.GetBalance(db, &user)

		if errors.Is(err, sql.ErrNoRows) {
			return echo.NewHTTPError(http.StatusPreconditionFailed, pkg.ErrNoUser.Error())
		}

		if err != nil {
			log.Error("error occurred while getting user balance", zap.Error(err))
			return echo.ErrInternalServerError
		}

		err = exchangerates.ValidateCurrency(c, config, &user)
		if err != nil {
			if errors.Is(err, pkg.ErrUnknownCy) {
				return echo.NewHTTPError(http.StatusBadRequest, pkg.ErrInvalidInput.Error())
			}
			return echo.ErrInternalServerError
		}

		return c.JSON(http.StatusOK, user)
	}
}
