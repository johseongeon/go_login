package pkg

import (
	"database/sql"
	"encoding/json"
	"main/internal/auth"
	"main/internal/user"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

// Request method: POST
func LoginHandler(db *sql.DB, skey []byte) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		var loginUser user.User
		err := json.NewDecoder(r.Body).Decode(&loginUser)
		if err != nil {
			http.Error(w, "Invalid request body: "+err.Error(), http.StatusBadRequest)
			return
		}

		if loginUser.User_id == "" || loginUser.Password == "" {
			http.Error(w, "Username and password are required", http.StatusBadRequest)
			return
		}

		loggedIn, err := user.Login(db, &loginUser)
		if err != nil {
			http.Error(w, "Failed to log in user: "+err.Error(), http.StatusInternalServerError)
			return
		}

		if loggedIn {
			jwt, err := auth.GenerateJWT(&loginUser, skey)
			if err != nil {
				http.Error(w, "Failed to generate JWT: "+err.Error(), http.StatusInternalServerError)
				return
			}
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(jwt))
		} else {
			http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		}
	}
}
