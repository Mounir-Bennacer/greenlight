package main

import (
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"strconv"
	"time"

	_ "github.com/lib/pq"
	"greenlight.mounirbennacer.com/config"
	"greenlight.mounirbennacer.com/database"
	"greenlight.mounirbennacer.com/internal/data"
)

// version is set at compile time
const version = "1.0.0"

type application struct {
	logger *slog.Logger
	config config.Config
	models data.Models
}

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	flag.StringVar(&config.Envs.Port, "port", config.Envs.Port, "API server port")
	flag.StringVar(&config.Envs.Env, "env", config.Envs.Env, "Environment (development|staging|production)")
	flag.StringVar(&config.Envs.Host, "db-dsn", config.Envs.Host, "Postgresql DSN")

	flag.IntVar(&config.Envs.MaxOpenConns, "db-max-open-conns", config.Envs.MaxOpenConns, "PostgreSQL max open connections")
	flag.IntVar(&config.Envs.MaxIdleConns, "db-max-idle-conns", config.Envs.MaxIdleConns, "PostgreSQL max idle connections")
	flag.DurationVar(&config.Envs.MaxIdleTime, "db-max-idle-time", config.Envs.MaxIdleTime*time.Minute, "PostgreSQL max connection idle time")

	flag.Parse()

	port, err := strconv.Atoi(config.Envs.Port)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	db, err := database.OpenDB(config.Envs)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	defer db.Close()

	logger.Info("database connection pool established")

	app := &application{
		config: config.Envs,
		logger: logger,
		models: data.NewModels(db),
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
