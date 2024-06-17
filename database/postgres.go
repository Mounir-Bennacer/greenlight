package database

import (
	"context"
	"database/sql"
	"time"

	_ "github.com/lib/pq"
	"greenlight.mounirbennacer.com/config"
)

func OpenDB(cfg config.Config) (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg.Host)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(cfg.MaxOpenConns)
	db.SetMaxIdleConns(cfg.MaxIdleConns)
	db.SetConnMaxIdleTime(cfg.MaxIdleTime)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
