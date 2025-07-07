package db

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)
type DBConnector struct {
	connection *pgxpool.Pool
}

func (db DBConnector) queryRow() pgx.Rows{
	rows, err := db.connection.Query(context.Background(), "select 'Hello, world!'")
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}
	return rows
}