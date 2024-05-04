package main

import (
	"fmt"
	"os"
)

func commandHelp() error {
	fmt.Println("Available Commands:")
	commands := createNewCommands()
	for _, cmd := range commands {
		fmt.Printf(" - %s: %s\n", cmd.name, cmd.description)
	}
	return nil
}

func commandExit() error {
	fmt.Println("Exiting program")
	os.Exit(0)
	return nil
}
