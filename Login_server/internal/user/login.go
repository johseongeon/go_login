package user

import (
	"database/sql"

	"log"

	"golang.org/x/crypto/bcrypt"
)

func Login(db *sql.DB, user *User) (bool, error) {
	var hashPw string

	err := db.QueryRow(
		`SELECT id, username, password_hash
		 FROM users
		 WHERE user_id = ?`,
		user.User_id,
	).Scan(&user.ID, &user.Username, &hashPw)
	if err != nil {
		log.Println("Query error:", err)
		return false, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashPw), []byte(user.Password))
	if err != nil {
		log.Println("Password comparison error:", err)
		return false, err
	}

	return true, nil
}
