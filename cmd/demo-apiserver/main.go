package main

import (
	"os"

	"demo/internal/apiserver"
)

func main() {
	command := apiserver.NewAppCommand()
	if err := command.Execute(); err != nil {
		os.Exit(1)
	}
}
