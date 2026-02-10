package pokeapi

import (
	"fmt"
	"net/url"
)

type Location struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type LocationResponse struct {
	Count    int        `json:"count"`
	Next     string     `json:"next"`
	Previous any        `json:"previous"`
	Results  []Location `json:"results"`
}

type LocationDetailsResponse struct {
	Name              string             `json:"name"`
	URL               string             `json:"url"`
	PokemonEncounters []PokemonEncounter `json:"pokemon_encounters"`
}

func (p *Api) GetLocation(limit int, offset int) ([]Location, error) {
	params := url.Values{}
	params.Add("limit", fmt.Sprintf("%d", limit))
	params.Add("offset", fmt.Sprintf("%d", offset))
	var lr LocationResponse
	err := getRequest(p, "location-area", "?"+params.Encode(), &lr)
	if err != nil {
		return []Location{}, fmt.Errorf("Error getting location: %w", err)
	}

	return lr.Results, nil
}

func (p *Api) GetLocationDetails(name string) (LocationDetailsResponse, error) {
	var ldr LocationDetailsResponse
	err := getRequest(p, fmt.Sprintf("location-area/%s", name), "", &ldr)

	if err != nil {
		return LocationDetailsResponse{}, fmt.Errorf("Error getting location: %w", err)
	}
	return ldr, err
}
