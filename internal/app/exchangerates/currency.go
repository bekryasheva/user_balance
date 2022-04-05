package exchangerates

import (
	"github.com/labstack/echo/v4"
	"user_balance/internal/app"
	"user_balance/internal/app/model"
)

func ValidateCurrency(c echo.Context, config app.Config, user *model.UserBalance) error {
	currency := c.QueryParam("currency")

	if currency == "" || currency == "RUB" {
		user.Currency = "RUB"
		return nil
	}

	err := ExchangeRates(currency, user, config)
	if err != nil {
		return err
	}

	return nil
}
