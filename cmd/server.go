package main

import (
	"context"
	"fmt"
	"log/slog"
	"net"
	"os"

	"github.com/KingKeleos/photopro-user-service"
	"github.com/KingKeleos/photopro-user-service/config"
	"github.com/KingKeleos/photopro-user-service/database"
	pb "github.com/KingKeleos/photopro-user-service/grpc"
	"google.golang.org/grpc"
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
	//Connecting to the database of the microservice
	pg_con, err := database.Connect(ctx, config)
	if err != nil {
		slog.Error("connecting to database", "error", err)
		os.Exit(0)
	}
	//Executing migration to microservice database
	err = database.Migrate(pg_con, config.DatabaseName)
	if err != nil {
		slog.Error("migrating database", "error", err)
		os.Exit(0)
	}
	slog.Info("finished migrating database")
	//Setting the created client for accessing the database
	database.PGClient = pg_con
	port := os.Getenv("PORT")
	if port == "" {
		slog.Info("reading port from env", "info", fmt.Errorf("PORT is empty, using default"))
		port = "8080"
	}
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		slog.Error("failed to listen", "error", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterUserServiceServer(grpcServer, userservice.NewUserServer())
	grpcServer.Serve(lis)
}
