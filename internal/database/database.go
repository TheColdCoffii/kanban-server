package database

import (
	"context"
	"errors"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

var dbPool *pgxpool.Pool

func InitializeDatabasePool() error {
	env := os.Getenv("ENV")
	if env == "development" {
		pool, err := pgxpool.New(context.Background(), os.Getenv("DEV_DB_CONN_STRING"))
		if err != nil {
			return err
		}
		log.Println("Connected to dev db")
		dbPool = pool
	} else if env == "production" {
		pool, err := pgxpool.New(context.Background(), os.Getenv("PROD_DB_CONN_STRING"))
		if err != nil {
			return err
		}
		log.Println("Connected to prod db")
		dbPool = pool
	} else {
		return errors.New("Could not initialize db connection")
	}
	return nil
}

func GetClient() *pgxpool.Conn {
	if dbPool == nil {
		panic(errors.New("Pool is not initialized"))
	}
	client, err := dbPool.Acquire(context.Background())
	if err != nil {
		panic(err)
	}
	return client
}
