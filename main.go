package main

import "github.com/Dhinihan/pokedexcli/internal/pokeapi"

type config struct {
	client *pokeapi.Api
}

func main() {
	config := NewConfig()
	replyer(config)
}

func NewConfig() config {
	return config{
		client: pokeapi.NewPokeapi(),
	}
}
