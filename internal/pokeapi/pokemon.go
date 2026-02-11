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
	URL            string
	BaseExperience int `json:"base_experience"`
	Height         int `json:"height"`
	Stats          []struct {
		BaseStat int `json:"base_stat"`
		Stat     struct {
			Name string `json:"name"`
		} `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Type struct {
			Name string `json:"name"`
		} `json:"type"`
	} `json:"types"`
	Weight int `json:"weight"`
}

func NewPokemon(name string) PokemonDetails {
	return PokemonDetails{Name: name}
}

type ParsedStats struct {
	Hp             int
	Attack         int
	Defense        int
	SpecialAttack  int
	SpecialDefense int
	Speed          int
}

func (p *PokemonDetails) ParseStats() ParsedStats {
	parsed := ParsedStats{}

	for _, stat := range p.Stats {
		switch stat.Stat.Name {
		case "hp":
			parsed.Hp = stat.BaseStat
		case "attack":
			parsed.Attack = stat.BaseStat
		case "defense":
			parsed.Defense = stat.BaseStat
		case "special-attack":
			parsed.SpecialAttack = stat.BaseStat
		case "special-defense":
			parsed.SpecialDefense = stat.BaseStat
		case "speed":
			parsed.Speed = stat.BaseStat
		}
	}
	return parsed
}

func (p *Api) GetPokemonDetails(name string) (PokemonDetails, error) {
	var pd PokemonDetails
	err := getRequest(p, fmt.Sprintf("pokemon/%s", name), "", &pd)

	if err != nil {
		return PokemonDetails{}, fmt.Errorf("Error getting pokemon %s: %w", name, err)
	}
	return pd, err
}
