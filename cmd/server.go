package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/KingKeleos/photopro-user-service/config"
	"github.com/KingKeleos/photopro-user-service/database"
)

func main() {
	ctx := context.Background()
	slog.Info("Starting User Service")
	slog.Info("Migrating Database")

	config, err := config.ReadConfig()
	if err != nil {
		slog.Error("reading config", "error", err)
		os.Exit(0)
	}

	pg_con, err := database.Connect(ctx, config)
	if err != nil {
		slog.Error("connecting to database", "error", err)
		os.Exit(0)
	}
	err = database.Migrate(pg_con, config.DatabaseName)
	if err != nil {
		slog.Error("migrating database", "error", err)
		os.Exit(0)
	}
}
