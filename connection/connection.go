package connection

import (
	"database/sql"
	"fmt"
	"os"
)

// _ "github.com/microsoft/go-mssqldb"
func Connect() (*sql.DB, error) {
	var DBCONN = os.Getenv("DBCONN") //database connection string, uses Windows Authentication

	if DBCONN == "" {
		return nil, fmt.Errorf("DBCONN environment variable is not set")
	}

	db, err := sql.Open("sqlserver", DBCONN)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	if err := db.Ping(); err != nil {
		db.Close()
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}
	return db, nil
}
