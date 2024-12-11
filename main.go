package main

import (
	"github.com/OoS-MaMaD/testmaster/cmd"
)

func main() {
	loadCredentials()
	cmd.Execute()
}
