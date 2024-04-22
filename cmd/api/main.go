package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"strconv"
	"time"

	_ "github.com/lib/pq"
	"greenlight.mounirbennacer.com/config"
)

// version is set at compile time
const version = "1.0.0"

type application struct {
	logger *slog.Logger
	config config.Config
}

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	flag.StringVar(&config.Envs.Port, "port", config.Envs.Port, "API server port")
	flag.StringVar(&config.Envs.Env, "env", config.Envs.Env, "Environment (development|staging|production)")
	flag.StringVar(&config.Envs.Host, "db-dsn", config.Envs.Host, "Postgresql DSN")

	flag.Parse()

	port, err := strconv.Atoi(config.Envs.Port)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	db, err := openDB(config.Envs)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	defer db.Close()

	logger.Info("database connection pool established")

	app := &application{
		config: config.Envs,
		logger: logger,
	}

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		ErrorLog:     slog.NewLogLogger(logger.Handler(), slog.LevelError),
	}

	logger.Info("starting server", "addr", srv.Addr, "env", config.Envs.Env)

	err = srv.ListenAndServe()
	logger.Error(err.Error())
	os.Exit(1)
}

func openDB(cfg config.Config) (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg.Host)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
