package pokedex

import "github.com/Dhinihan/pokedexcli/internal/pokeapi"

type Pokedex map[string]pokeapi.PokemonDetails

func (p Pokedex) Add(pm pokeapi.PokemonDetails) {
	p[pm.Name] = pm
}

func (p Pokedex) Get(name string) (pokeapi.PokemonDetails, bool) {
	pm, ok := p[name]
	return pm, ok
}

func NewPokedex() Pokedex {
	return Pokedex{}
}
