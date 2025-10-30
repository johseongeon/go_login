package user

import (
	"database/sql"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func Register(db *sql.DB, user *User) error {
	// hash + salt
	pwHash, err := bcrypt.GenerateFromPassword([]byte(user.Password+user.Username), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("error hashing password: %w", err)
	}

	// Start a transaction
	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("error starting transaction: %w", err)
	}
	defer tx.Rollback()

	// Prepare statement
	stmt, err := tx.Prepare("INSERT INTO users (username, user_id, password_hash) VALUES (?, ?, ?)")
	if err != nil {
		return fmt.Errorf("error preparing statement: %w", err)
	}
	defer stmt.Close()

	result, err := stmt.Exec(user.Username, user.User_id, string(pwHash))
	if err != nil {
		return fmt.Errorf("error executing insert: %w", err)
	}

	// Get auto-incremented ID
	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("error getting last insert ID: %w", err)
	}
	user.ID = int(lastInsertID)

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("error committing transaction: %w", err)
	}

	return nil
}
