package pokeapi

import "fmt"

type Pokemon struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type PokemonEncounter struct {
	Pokemon Pokemon `json:"pokemon"`
}

type PokemonDetails struct {
	Name           string `json:"name"`
	URL            string `json:"url"`
	BaseExperience int    `json:"base_experience"`
}

func (p *Api) GetPokemonDetails(name string) (PokemonDetails, error) {
	var pd PokemonDetails
	err := getRequest(p, fmt.Sprintf("pokemon/%s", name), "", &pd)

	if err != nil {
		return PokemonDetails{}, fmt.Errorf("Error getting pokemon %s: %w", name, err)
	}
	return pd, err
}
