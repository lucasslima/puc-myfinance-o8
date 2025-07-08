package db

import (
	"database/sql"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
)
type DBConnector struct {
	Connection *sql.DB
}

func CreateConnection() *DBConnector {
	    // Capture connection properties.
    cfg := mysql.NewConfig()
    cfg.User = os.Getenv("DBUSER")
    cfg.Passwd = os.Getenv("DBPASS")
    cfg.Net = "tcp"
    cfg.Addr = "127.0.0.1:3306"
    cfg.DBName = "myfinance"

    // Get a database handle.
    var err error
    connection, err := sql.Open("mysql", cfg.FormatDSN())
    if err != nil {
        log.Fatalf("Error connecting to the database: %s", err)
    }
	log.Printf("Successfully connected to %s in database %s", cfg.Addr, cfg.DBName)
	return &DBConnector{Connection: connection}

}