package main

import (
	"fmt"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/spf13/cobra"
	"go.uber.org/zap"

	"user_balance/internal/app"
	"user_balance/internal/app/database"
	"user_balance/internal/app/handlers"
)

var configFile string

const (
	defaultConfigPath = "configs/config.yaml"
)

var rootCmd = &cobra.Command{
	Use:   "user-balance",
	Short: "Microservice for working with user balance",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {

		log, err := zap.NewProduction()
		if err != nil {
			fmt.Printf("failed to initialize logger: %v\n", err)
			return
		}

		defer log.Sync()

		config, err := app.ReadConfigFromFile(configFile)
		if err != nil {
			log.Fatal("failed to read config", zap.Error(err))
		}

		fmt.Printf("Using config file: %s\n", configFile)

		db, err := database.NewPostgresDB(config, log)
		if err != nil {
			log.Fatal("failed to open a DB connection", zap.Error(err))
		}
		defer db.Close()

		e := echo.New()

		e.GET("/balance/:id", handlers.BalanceHandler(db, log, config))
		e.POST("/deposit", handlers.DepositHandler(db, log))
		e.POST("/withdrawal", handlers.WithdrawalHandler(db, log))
		e.POST("/transfer", handlers.TransferHandler(db, log))
		e.GET("/history/:id", handlers.HistoryHandler(db, log))

		e.Logger.Fatal(e.Start(config.API.Address))
	},
}

func main() {
	err := rootCmd.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&configFile, "config","c", defaultConfigPath, "Setting the configuration file")
}