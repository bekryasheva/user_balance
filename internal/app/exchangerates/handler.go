package exchangerates

import (
	"encoding/json"
	"fmt"
	"net/http"
	"user_balance/internal/app"
	"user_balance/internal/app/model"
	"user_balance/internal/app/pkg"
)

func ExchangeRates(currency string, user *model.UserBalance, config app.Config) error {

	url := fmt.Sprintf("%s?apikey=%s&base_currency=RUB", config.FreeCurrencyApi.URL, config.FreeCurrencyApi.AccessKey)

	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	decoder := json.NewDecoder(resp.Body)

	var rates model.RatesResponse

	err = decoder.Decode(&rates)
	if err != nil {
		return err
	}

	var rateExist bool

	for cy := range rates.Data {
		if cy == currency {
			rateExist = true
			break
		}
	}

	if !rateExist {
		return pkg.ErrUnknownCy
	}

	user.Balance = user.Balance * rates.Data[currency]
	user.Currency = currency

	return nil
}