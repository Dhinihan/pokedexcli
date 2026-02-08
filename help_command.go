package main

import "fmt"

func commandHelp(c config) error {
	fmt.Print("Welcome to the Pokedex!\nUsage:\n\n")
	comandos := getCommands()
	for nome, comando := range comandos {
		fmt.Printf("%s:	%s\n", nome, comando.description)
	}
	return nil
}

