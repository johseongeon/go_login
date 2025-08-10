package pkg

import (
	"bufio"
	"bytes"
	_ "crypto/tls" // for https local test
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func Register(reader *bufio.Reader) {

	// username
	fmt.Print("Enter username: ")
	username, _ := reader.ReadString('\n')
	username = strings.TrimSpace(username)

	// user_id
	fmt.Print("Enter user ID: ")
	userID, _ := reader.ReadString('\n')
	userID = strings.TrimSpace(userID)

	// password
	password := InsertPassword(reader)

	// Register information
	registerData := map[string]string{
		"username": username,
		"user_id":  userID,
		"password": password,
	}
	jsonData, err := json.Marshal(registerData)
	if err != nil {
		panic(err)
	}

	client := &http.Client{
		// for https local test
		/*
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			},
		*/
	}

	// Generate request
	req, err := http.NewRequest("POST", os.Getenv("SERVER_URI")+"/register", bytes.NewBuffer(jsonData))
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json")

	// Send request
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// Read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Status: %s\n", resp.Status)
	fmt.Printf("Response: %s\n", string(body))
}
