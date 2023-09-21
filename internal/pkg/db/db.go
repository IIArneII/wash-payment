package db

import (
	"context"
	"database/sql"
	"time"

	"github.com/gocraft/dbr/v2"
	_ "github.com/lib/pq"
	"github.com/pressly/goose"
)

func InitDB(dsn string) (dbPool *dbr.Connection, err error) {
	dbPool, err = dbr.Open("postgres", dsn, nil)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	ping := false
	for !ping {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
			if err = dbPool.PingContext(ctx); err == nil {
				ping = true
			} else {
				time.Sleep(100 * time.Millisecond)
			}
		}
	}

	dbPool.SetMaxOpenConns(10)
	dbPool.SetConnMaxLifetime(time.Second * 20)
	dbPool.SetConnMaxIdleTime(time.Second * 60)

	return
}

func UpMigrations(db *sql.DB, dir string) error {

	if err := goose.SetDialect("postgres"); err != nil {
		return err
	}

	if err := goose.Up(db, dir); err != nil {
		return err
	}

	return nil
}
