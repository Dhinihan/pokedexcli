package main

import (
	"bufio"
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func(config, []string) error
	args        []string
}

func replyer(c config) {
	comandos := getCommands()
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		text := scanner.Text()
		cleanedInput := cleanInput(text)
		if len(cleanedInput) == 0 {
			fmt.Println()
			continue
		}
		comando, ok := comandos[cleanedInput[0]]
		args := cleanedInput[1:]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}
		err := comando.callback(c, args)
		if err != nil {
			fmt.Println(err.Error())
		}
	}
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Explore the next locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Go Back on the previous locations",
			callback:    commandMapBack,
		},
		"explore": {
			name:        "explore <location_name>",
			description: "Show pokémon at location",
			callback:    commandExplore,
		},
	}
}
