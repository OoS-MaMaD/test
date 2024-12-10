package main

import (
	"fmt"
	"os"

	"github.com/OoS-MaMaD/testmaster/cmd"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	fmt.Println(os.Getenv("Password"))
	cmd.Execute()
}
