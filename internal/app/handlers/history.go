package handlers

import (
	"database/sql"
	"errors"
	"net/http"
	"user_balance/internal/app/database"
	"user_balance/internal/app/pkg"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

func HistoryHandler(db *sql.DB, log *zap.Logger) echo.HandlerFunc {
	return func(c echo.Context) error {
		var id int64
		errPath := echo.PathParamsBinder(c).Int64("id", &id).BindError()

		var (
			orderBy string
			sort string
			limit int64
			offset int64
		)

		errQuery := echo.QueryParamsBinder(c).
			String("order_by", &orderBy).
			String("sort", &sort).
			Int64("limit", &limit).
			Int64("offset", &offset).BindError()

		if errPath != nil || errQuery != nil {
			return echo.NewHTTPError(http.StatusBadRequest, pkg.ErrInvalidInput.Error())
		}

		err := pkg.ValidateHistoryParam(id, &orderBy, &sort, &limit, &offset)
		if err != nil {
			if errors.Is(err, pkg.ErrInvalidInput) {
				log.Error("input validation failed", zap.Error(err))
				return echo.NewHTTPError(http.StatusBadRequest, pkg.ErrInvalidInput.Error())
			}
			log.Error("error occurred while checking history param")
			return echo.ErrInternalServerError
		}

		transactions, err := database.GetHistory(db, id, orderBy, sort, limit, offset)
		if err != nil {
			if errors.Is(err, pkg.ErrNoUser) {
				return echo.NewHTTPError(http.StatusPreconditionFailed, pkg.ErrNoUser.Error())
			}
			log.Error("failed to get history of transaction", zap.Error(err))
			return echo.ErrInternalServerError
		}

		return c.JSON(http.StatusOK, transactions)
	}
}