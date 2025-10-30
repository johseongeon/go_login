package pkg

import (
	"database/sql"
	"encoding/json"
	"log"
	"main/internal/user"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

// Request method: POST
func RegisterHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if r.Header.Get("Content-Type") != "application/json" {
			http.Error(w, "Content-Type must be application/json", http.StatusUnsupportedMediaType)
			return
		}

		var newUser user.User
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&newUser); err != nil {
			http.Error(w, "Invalid JSON body: "+err.Error(), http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		err := user.Register(db, &newUser)
		if err != nil {
			http.Error(w, "Failed to register user: "+err.Error(), http.StatusInternalServerError)
			return
		}
		log.Printf("User %s registered successfully", newUser.Username)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("User registered successfully"))
	}
}
