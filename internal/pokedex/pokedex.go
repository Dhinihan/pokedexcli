package pokedex

import (
	"fmt"

	"github.com/Dhinihan/pokedexcli/internal/pokeapi"
)

type Pokedex map[string]pokeapi.PokemonDetails

func (p Pokedex) Add(pm pokeapi.PokemonDetails) {
	p[pm.Name] = pm
}

func (p Pokedex) Get(name string) (pokeapi.PokemonDetails, bool) {
	pm, ok := p[name]
	return pm, ok
}

func (p Pokedex) List() string {
	var saida string
	for name := range p {
		saida = fmt.Sprintf("%s -%12s\n", saida, name)
	}
	return saida
}

func NewPokedex() Pokedex {
	return Pokedex{}
}
