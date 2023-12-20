package main

import (
	"fmt"

	"matryer/internal/db"
	"matryer/internal/platform/server"
	"matryer/internal/platform/server/handlers"

	"matryer/pkg/logger"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

var (
	apiPort    = os.Getenv("API_PORT")
	dbType     = os.Getenv("DB_TYPE")
	dbUser     = os.Getenv("DB_USER")
	dbPassword = os.Getenv("DB_PASSWORD")
	dbHost     = os.Getenv("DB_HOST")
	dbPort     = os.Getenv("DB_PORT")
	dbName     = os.Getenv("DB_NAME")
)

func run() error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbHost, dbPort, dbName)
	client, err := sqlx.Connect(dbType, dsn)
	if err != nil {
		return err
	}
	defer client.Close()

	logger := logger.New()
	userStore := db.NewMySQLUserStore(client, logger)
	userHandler := handlers.NewUserHandler(logger, userStore)

	router := chi.NewRouter()
	srv := server.New(router, logger, userHandler)
	addr := fmt.Sprintf(":%s", apiPort)

	logger.Infof("listening and serving %s", addr)

	return http.ListenAndServe(addr, srv)
}
