package main

import (
	"log"
	"main/internal/db"
	server "main/internal/handler"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load(".env")
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

	http.HandleFunc("/register", server.RegisterHandler(database))

	http.HandleFunc("/login", server.LoginHandler(database, jwtSecret))

	// http.ListenAndServeTLS(addr, certFile, keyFile string, handler Handler)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
