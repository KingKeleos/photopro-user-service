package database

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/KingKeleos/photopro-user-service/config"
)

func Connect(ctx context.Context, conf *config.DatabaseConfig) (*sql.DB, error) {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		conf.Host, conf.Port, conf.Username, conf.Password, conf.DatabaseName, conf.SSLMode)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(conf.MaxConnections)
	db.SetMaxIdleConns(conf.IdleConnections)
	db.SetConnMaxLifetime(conf.MaxConnectionLifeTime)

	return db, nil
}
