package app

import (
	"github.com/spf13/viper"
	"log"
)

type SectionAPI struct {
	Address string `mapstructure:"address"`
}

type SectionDatabase struct {
	Host string `mapstructure:"host"`
	Port string `mapstructure:"port"`
	User string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Name     string `mapstructure:"name"`
	Sslmode  string `mapstructure:"sslmode"`
}

type SectionCurrencyAPI struct {
	URL       string `mapstructure:"url"`
	AccessKey string `mapstructure:"access_key"`
}

type Config struct {
	API             SectionAPI         `mapstructure:"api"`
	Database        SectionDatabase    `mapstructure:"database"`
	FreeCurrencyApi SectionCurrencyAPI `mapstructure:"free_currency_api"`
}

func ReadConfigFromFile(path string) (Config, error) {
	config := &Config{
		API: SectionAPI{
			Address: ":8080",
		},
		Database: SectionDatabase{
			Host:     "localhost",
			Port:     "5432",
			User:     "userbalance",
			Password: "password",
			Name:     "userbalance",
			Sslmode:  "disable",
		},
		FreeCurrencyApi: SectionCurrencyAPI{
			URL: "https://freecurrencyapi.net/api/v2/latest",
		},
	}

	viper.SetConfigFile(path)
	viper.BindEnv("database.host", "DB_HOST")
	viper.BindEnv("database.port", "DB_PORT")
	viper.BindEnv("database.user", "DB_USER")
	viper.BindEnv("database.password", "DB_PASSWORD")
	viper.BindEnv("database.name", "DB_NAME")
	viper.BindEnv("database.sslmode", "DB_SSLMODE")

	viper.BindEnv("api.address", "API_ADDRESS")
	viper.BindEnv("free_currency_api.access_key", "CURRENCY_API_ACCESS_KEY")

	err := viper.ReadInConfig()
	if err != nil {
		log.Printf("%v\n", err)
		return Config{}, err
	}

	viper.GetString("api.address")

	err = viper.Unmarshal(&config)
	if err != nil {
		return Config{}, err
	}

	return *config, nil
}
