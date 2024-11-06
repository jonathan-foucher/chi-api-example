package database

import (
	"context"
	"fmt"
	"os"
	"github.com/jackc/pgx/v5"
)

var connection *pgx.Conn

func InitDbConn() (*pgx.Conn) {
	ctx := context.Background()
	var err error
	connection, err = pgx.Connect(ctx, os.Getenv("DB_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	return connection
}

func GetDbConnection() (*pgx.Conn) {
	return connection
}
