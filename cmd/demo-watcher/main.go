package main

import (
	"os"

	"demo/internal/watcher"
)

func main() {
	command := watcher.NewWatcherCommand()
	if err := command.Execute(); err != nil {
		os.Exit(1)
	}
}
