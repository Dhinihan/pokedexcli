package main

import (
	"fmt"

	"github.com/Dhinihan/pokedexcli/internal/pokecalc"
)

func commandCatch(c config, args []string) error {
	if len(args) < 1 {
		fmt.Println("É preciso informar o pokémon, 'catch <pokemon>'")
	}

	pName := args[0]

	pokemon, err := c.client.GetPokemonDetails(pName)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pName)

	pegou := pokecalc.ThrowPokeball(pokemon, pokecalc.NewRng())

	if pegou {
		fmt.Printf("%s was caught\n", pName)
		c.pokeDex.Add(pokemon)
	} else {
		fmt.Printf("%s escaped!\n", pName)
	}
	return nil
}
