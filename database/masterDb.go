package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/microsoft/go-mssqldb"
)

// Database configurations
var masterDataConfig = DBConfig{
	Server:   "igsql.database.windows.net",
	Port:     1433,
	User:     "igsql",
	Password: "SqlDb$99",
	Database: "MasterData",
}

var salesDataConfig = DBConfig{
	Server:   "igsqlsales.database.windows.net",
	Port:     1433,
	User:     "igsqlsales",
	Password: "SqlDb$99",
	Database: "Sales",
}

// DBConfig represents the configuration for a database
type DBConfig struct {
	Server   string
	Port     int
	User     string
	Password string
	Database string
}

// ConnectDB connects to the specified database
func ConnectDB(config DBConfig) (*sql.DB, error) {
	// Build connection string
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;",
		config.Server, config.User, config.Password, config.Port, config.Database)

	// Create connection pool
	db, err := sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: ", err.Error())
		return nil, err
	}

	// Ping the database to verify the connection
	ctx := context.Background()
	err = db.PingContext(ctx)
	if err != nil {
		log.Fatal(err.Error())
		return nil, err
	}

	// Return the database connection and any error.
	return db, nil
}

func main() {
	ConnectDB(masterDataConfig)
	ConnectDB(salesDataConfig)
}
