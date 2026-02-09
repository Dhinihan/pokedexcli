package main

import (
	"fmt"
)

func commandExplore(c config, args []string) error {

	if len(args) < 1 {
		fmt.Println("É preciso informar o local, 'explore <local>'")
	}

	lName := args[0]

	res, err := c.client.GetLocationDetails(lName)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf("Exploring %s...\n", lName)
	for _, e := range res.PokemonEncounters {
		fmt.Printf("- %s\n", e.Pokemon.Name)
	}
	return nil
}
