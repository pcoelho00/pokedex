package main

import (
	"fmt"
)

// type cliCommand struct {
// 	name        string
// 	description string
// 	callback    func() error
// }

// func commandHelp() error {

// 	fmt.Println("test")

// }

// func createCommands() map[string]cliCommand {
// 	return map[string]cliCommand{
// 		"help": {
// 			name:        "help",
// 			description: "Displays a help message",
// 			callback:    commandHelp,
// 		},
// 		"exit": {
// 			name:        "exit",
// 			description: "Exit the Pokedex",
// 			callback:    commandExit,
// 		},
// 	}
// }

func main() {
	fmt.Println("Starting loop:")
	startRepl()
}
