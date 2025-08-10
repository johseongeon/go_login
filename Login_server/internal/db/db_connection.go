package db

import (
	"database/sql"
	"fmt"

	"log"

	_ "github.com/go-sql-driver/mysql"
)

// driver: "mysql", dataSource: os.Getenv("MYSQL_URI")
func ConnectToDB(driver, dataSource string) (*sql.DB, error) {
	db, err := sql.Open(driver, dataSource)
	if err != nil {
		return nil, fmt.Errorf("error opening database: %w", err)
	}

	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("error connecting to database: %w", err)
	}

	log.Println("Connected to the database successfully")
	return db, nil
}
