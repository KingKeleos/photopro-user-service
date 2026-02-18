package config

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

type DatabaseConfig struct {
	Host                  string
	Port                  string
	Username              string
	Password              string
	SSLMode               string
	DatabaseName          string
	MaxConnections        int
	IdleConnections       int
	MaxConnectionLifeTime time.Duration
}

func ReadConfig() (*DatabaseConfig, error) {
	host := os.Getenv("DB_HOST")
	if host == "" {
		return nil, fmt.Errorf("DB_HOST not set")
	}
	port := os.Getenv("DB_PORT")
	if port == "" {
		return nil, fmt.Errorf("DB_PORT not set")
	}
	name := os.Getenv("DB_USER")
	if name == "" {
		return nil, fmt.Errorf("DB_USER not set")
	}
	password := os.Getenv("DB_PASS")
	if password == "" {
		return nil, fmt.Errorf("DB_PASS not set")
	}
	databaseName := os.Getenv("DB_NAME")
	if databaseName == "" {
		return nil, fmt.Errorf("DB_NAME not set")
	}
	sslmode := os.Getenv("DB_SSL_MODE")
	if sslmode == "" {
		return nil, fmt.Errorf("DB_SSL_MODE not set")
	}
	maxCon, err := strconv.Atoi(os.Getenv("MAX_DATABASE_CONNECTIONS"))
	if err != nil {
		return nil, err
	}
	if maxCon == 0 {
		return nil, fmt.Errorf("MAX_DATABASE_CONNETIONS not set")
	}

	dbConfig := DatabaseConfig{
		Host:           host,
		Port:           port,
		Username:       name,
		Password:       password,
		DatabaseName:   databaseName,
		SSLMode:        sslmode,
		MaxConnections: maxCon,
	}

	return &dbConfig, nil
}
