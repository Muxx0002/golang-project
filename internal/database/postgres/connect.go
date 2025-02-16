package actions

import (
	"context"
	"log"
	"time"

	"github.com/google/logger"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/spf13/viper"
)

var DB *pgxpool.Pool

func InitDB() {
	var err error
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	DB, err = pgxpool.New(ctx, viper.GetString("database_url"))
	if err != nil {
		log.Fatalf("database connect error: %v", err)
	}
	DB.Config().MaxConns = int32(viper.GetInt32("max_conns"))
	DB.Config().MaxConnIdleTime = 3 * time.Second
	DB.Config().MaxConnLifetime = 10 * time.Minute
	if err := DB.Ping(ctx); err != nil {
		log.Fatalf("database connect error: %v", err)
	}
	logger.Info("connection to database successful!")
}
