package pokecalc

import (
	"fmt"
	"math"
	"math/rand/v2"

	"github.com/Dhinihan/pokedexcli/internal/pokeapi"
)

const catch_factor = 0.00377619213

type Rng interface {
	Float64() float64
}

type defaultRng struct{}

func (r defaultRng) Float64() float64 {
	return rand.Float64()
}

func NewRng() Rng {
	return defaultRng{}
}

func ThrowPokeball(p pokeapi.PokemonDetails, rng Rng) bool {
	odds := calcCatchOdds(p.BaseExperience)
	roll := rng.Float64()
	fmt.Printf("roll: %.3f, target %.3f\n", roll, odds)
	return roll <= odds
}

func calcCatchOdds(b int) float64 {
	return 0.99*math.Exp(-catch_factor*float64(b)) + 0.01
}
