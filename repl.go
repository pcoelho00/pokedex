package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	// callback    func() error
}

func createNewCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			// callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			// callback:    commandExit,
		},
	}
}

func cleanInput(text string) []string {

	output := strings.TrimSpace(text)
	output = strings.ToLower(output)
	words := strings.Fields(output)
	return words
}

func startRepl() {

	reader := bufio.NewScanner(os.Stdin)
	for reader.Scan() {
		text := cleanInput(reader.Text())
		if len(text) == 0 {
			continue
		}

		command := text[0]

		commands := createNewCommands()

		switch command {
		case "help":
			fmt.Println("Available Commands:")
			for name := range commands {
				fmt.Println(" - ", name)
			}

		case "exit":
			os.Exit(0)
		default:
			fmt.Println("echo: ", text)
		}
	}
}
