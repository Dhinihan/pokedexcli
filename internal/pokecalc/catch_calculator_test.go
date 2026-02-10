package pokecalc

import (
	"testing"

	"github.com/Dhinihan/pokedexcli/internal/pokeapi"
)

type mockRng struct{ valor float64 }

func (r mockRng) Float64() float64 { return r.valor }

func TestThrowPokeball(t *testing.T) {
	caterpie := pokeapi.PokemonDetails{Name: "Caterpie", BaseExperience: 39}
	ceruledge := pokeapi.PokemonDetails{Name: "Ceruledge", BaseExperience: 263}
	blissey := pokeapi.PokemonDetails{Name: "Blissey", BaseExperience: 635}

	type args struct {
		poke pokeapi.PokemonDetails
		roll float64
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Captura caterpie", args{caterpie, 0.864}, true},
		{"Não Captura Caterpie", args{caterpie, 0.865}, false},
		{"Captura Ceruledge", args{ceruledge, 0.376}, true},
		{"Não Captura Ceruledge", args{ceruledge, 0.377}, false},
		{"Captura Blissey", args{blissey, 0.099}, true},
		{"Não Captura Blissey", args{blissey, 0.101}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			mockRng := mockRng{tt.args.roll}
			if got := ThrowPokeball(tt.args.poke, mockRng); got != tt.want {
				t.Errorf("ThrowPokeball() = %v, want %v", got, tt.want)
			}
		})
	}
}
