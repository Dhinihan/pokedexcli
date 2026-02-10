package main

import (
	"github.com/Dhinihan/pokedexcli/internal/pokeapi"
	"github.com/Dhinihan/pokedexcli/internal/pokedex"
)

type config struct {
	client  *pokeapi.Api
	pokeDex pokedex.Pokedex
}

func main() {
	config := NewConfig()
	replyer(config)
}

func NewConfig() config {
	return config{
		client:  pokeapi.NewPokeapi(),
		pokeDex: pokedex.NewPokedex(),
	}
}
