package db

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type DB struct {
	Pool *pgxpool.Pool
}

func (db *DB) Connect(connectionUrl string) error {
	connectionConfiguration, err := pgxpool.ParseConfig(connectionUrl)

	if err != nil {
		return fmt.Errorf("DB.New: parse connection configuration end with error %v", err)
	}

	connectionConfiguration.AfterConnect = func(ctx context.Context, conn *pgx.Conn) error {
		fmt.Println("DB.New - AfterConnectClosure: Connected")
		return nil
	}

	db.Pool, err = pgxpool.NewWithConfig(context.Background(), connectionConfiguration)

	if err != nil {
		return fmt.Errorf("DB.New - NewWithConfig: create pool end with error %v", err)
	}

	if err = db.Pool.Ping(context.Background()); err != nil {
		return fmt.Errorf("DB.New - pool.Ping: ping end with error %v", err)
	}

	return err
}

func (db *DB) Close() error {
	if db.Pool != nil {
		db.Pool.Close()
	}
	return nil
}
