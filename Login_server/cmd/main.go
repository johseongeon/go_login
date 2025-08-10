package main

import (
	"log"
	"main/internal/db"
	pkg "main/pkg/http"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	//driver: "mysql"
	database, err := db.ConnectToDB("mysql", os.Getenv("MYSQL_URI"))
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer database.Close()

	jwtSecret := []byte(os.Getenv("JWT_SECRET"))

	http.HandleFunc("/register", pkg.RegisterHandler(database))

	http.HandleFunc("/login", pkg.LoginHandler(database, jwtSecret))

	// http.ListenAndServeTLS(addr, certFile, keyFile string, handler Handler)

	if err := http.ListenAndServe(":443", nil); err != nil {
		log.Fatal(err)
	}
}
