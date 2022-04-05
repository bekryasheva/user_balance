package model

type RatesResponse struct {
	Query struct {
		Timestamp    int    `json:"timestamp"`
		BaseCurrency string `json:"base_currency"`
	} `json:"query"`
	Data map[string]float64 `json:"data"`
}
