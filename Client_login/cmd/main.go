package main

import (
	"bufio"
	"fmt"
	"main/pkg"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("../.env")
	if err != nil {
		fmt.Println("Error loading .env file")
		return
	}
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Register or Login? (r/l): ")
		rol, _ := reader.ReadString('\n')
		rol = strings.TrimSpace(rol)

		if rol == "r" || rol == "R" {
			pkg.Register(reader)
			break
		} else if rol == "l" || rol == "L" {
			pkg.Login(reader)
			break
		} else {
			fmt.Println("Invalid option. Please enter 'r' for register or 'l' for login.")
		}
	}
}
