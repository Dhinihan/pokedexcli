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

// Fórmula de captura
//	const k = 0.00377619213
//	return 0.99*math.Exp(-k*x) + 0.01
