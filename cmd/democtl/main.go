package main

import (
	"os"

	"demo/internal/democtl/cmd"
)

func main() {
	command := cmd.NewDefaultDemoCtlCommand()
	if err := command.Execute(); err != nil {
		os.Exit(1)
	}
}
