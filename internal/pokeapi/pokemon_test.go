package pokeapi

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestGetPokemonDetails(t *testing.T) {
	type testCase struct {
		name string
		poke PokemonDetails
	}
	testCases := []testCase{
		{"ceruledge", PokemonDetails{"ceruledge", "url1", 321}},
		{"armarouge", PokemonDetails{"armarouge", "url2", 123}},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(tt *testing.T) {
			tt.Parallel()

			api := mockServer(tt, "/pokemon/"+tc.name, tc.poke)

			details, err := api.GetPokemonDetails(tc.name)
			if err != nil {
				tt.Errorf("UnexpectedError: %s", err.Error())
			}
			if diff := cmp.Diff(details, tc.poke); diff != "" {
				tt.Errorf("mismatch pokémon:\n%s", diff)
			}
		})
	}
}
