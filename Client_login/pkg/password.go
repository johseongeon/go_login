package pkg

import (
	"bufio"
	"fmt"
	"os"

	"golang.org/x/term"
)

func InsertPassword(reader *bufio.Reader) string {
	// password
	fmt.Print("Enter password: ")
	passwordBytes, err := term.ReadPassword(int(os.Stdin.Fd()))
	if err != nil {
		panic(err)
	}
	fmt.Println()

	fmt.Print("Re-enter password: ")
	rePasswordBytes, err := term.ReadPassword(int(os.Stdin.Fd()))
	if err != nil {
		panic(err)
	}
	fmt.Println()

	if string(passwordBytes) != string(rePasswordBytes) {
		fmt.Println("Passwords do not match. Please try again.")
		return InsertPassword(reader)
	}

	return string(passwordBytes)
}
