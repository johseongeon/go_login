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

	"golang.org/x/term"
)

func Login(reader *bufio.Reader) {

	// user_id
	fmt.Print("Enter user ID: ")
	userID, _ := reader.ReadString('\n')
	userID = strings.TrimSpace(userID)

	// password
	fmt.Print("Enter password: ")
	passwordBytes, err := term.ReadPassword(int(os.Stdin.Fd()))
	if err != nil {
		panic(err)
	}
	fmt.Println()
	password := string(passwordBytes)

	// login ingformation
	loginData := map[string]string{
		"user_id":  userID,
		"password": password,
	}

	// Marshal to JSON
	jsonData, err := json.Marshal(loginData)
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

	// generate request
	req, err := http.NewRequest("POST", os.Getenv("SERVER_URI")+"/login", bytes.NewBuffer(jsonData))
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json")

	// send request
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Status: %s\n", resp.Status)
	fmt.Printf("Response: %s\n", string(body))
}
