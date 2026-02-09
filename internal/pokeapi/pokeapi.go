package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/Dhinihan/pokedexcli/internal/pokeapicache"
)

const realUrl = "https://pokeapi.co/api/v2/"

type Location struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type Api struct {
	client  *http.Client
	baseUrl string
	cache   *pokeapicache.Cache
}
type LocationResponse struct {
	Count    int        `json:"count"`
	Next     string     `json:"next"`
	Previous any        `json:"previous"`
	Results  []Location `json:"results"`
}

type Pokemon struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type PokemonEncounter struct {
	Pokemon Pokemon `json:"pokemon"`
}

type LocationDetailsResponse struct {
	Name              string             `json:"name"`
	URL               string             `json:"url"`
	PokemonEncounters []PokemonEncounter `json:"pokemon_encounters"`
}

func NewPokeapi() *Api {
	return &Api{
		client:  http.DefaultClient,
		baseUrl: realUrl,
		cache:   pokeapicache.NewCache(1 * time.Hour),
	}
}

type queryParams map[string]string

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

func getRequest[T any](p *Api, path string, qp string, dataContainer *T) error {
	fullUrl := p.baseUrl + "/" + path + qp

	var data []byte
	data, ok := p.cache.Get(fullUrl)
	if !ok {
		request, err := http.NewRequest(http.MethodGet, fullUrl, nil)
		if err != nil {
			return fmt.Errorf("Erro ao montar request: %w", err)
		}
		res, err := p.client.Do(request)

		if err != nil {
			return fmt.Errorf("Erro ao chamar GET %s: %w", fullUrl, err)
		}
		defer res.Body.Close()
		data, err = io.ReadAll(res.Body)
		if err != nil {
			return fmt.Errorf("Erro ao processar GET %s: %w", fullUrl, err)
		}
		p.cache.Add(fullUrl, data)

	}
	if err := json.Unmarshal(data, dataContainer); err != nil {
		return fmt.Errorf("Erro ao decodificar json %s: %w", data, err)
	}
	return nil
}

func (qp queryParams) String() string {
	if len(qp) == 0 {
		return ""
	}
	pairs := []string{}
	for chave, valor := range qp {
		pairs = append(pairs, chave+"="+valor)
	}
	return "?" + strings.Join(pairs, "&")
}
