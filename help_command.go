package main

import "fmt"

func commandHelp(c config, args []string) error {
	fmt.Print("Welcome to the Pokedex!\nUsage:\n\n")
	comandos := getCommands()
	for _, comando := range comandos {
		fmt.Printf("%30s: %s\n", comando.name, comando.description)
	}
	return nil
}
