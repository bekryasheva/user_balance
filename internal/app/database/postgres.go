package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
	"net"
	"time"
	"user_balance/internal/app"
)

func NewPostgresDB(cfg app.Config, log *zap.Logger) (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.Name,
		cfg.Database.Sslmode)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	for attempt := 0; attempt < 10; attempt++ {
		err = db.Ping()
		if err == nil {
			break
		}

		if err, ok := err.(net.Error); !ok {
			return nil, err
		}

		log.Warn("failed to connect to database",
			zap.Error(err),
			zap.Int("attempt:", attempt))

		time.Sleep(5 * time.Second)
	}

	return db, nil
}

