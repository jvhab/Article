package postgres

import (
	"fmt"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	"grpc-article/config"
	"time"
)

const (
	setMaxOpenConns    = 60
	setConnMaxLifetime = 120
	setMaxIdleConns    = 30
	setConnMaxIdleTime = 20
)

func NewPostgresConn(cfg *config.Config) (*sqlx.DB, error) {
	dataSourceName := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.Login,
		cfg.Database.Password,
		cfg.Database.DBName,
		cfg.Database.SSLMode,
	)

	db, err := sqlx.Connect(cfg.Database.PGDriver, dataSourceName)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(setMaxOpenConns)
	db.SetConnMaxLifetime(setConnMaxLifetime * time.Second)
	db.SetMaxIdleConns(setMaxIdleConns)
	db.SetConnMaxIdleTime(setConnMaxIdleTime * time.Second)

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
