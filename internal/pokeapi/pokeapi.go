package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

const realUrl = "https://pokeapi.co/api/v2/"

type Location struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type pokeapi struct {
	client  *http.Client
	baseUrl string
}

func NewPokeapi() *pokeapi {
	return &pokeapi{
		client:  http.DefaultClient,
		baseUrl: realUrl,
	}
}

type queryParams map[string]string

func (p *pokeapi) GetLocation(limit int, offset int) ([]Location, error) {
	params := queryParams{
		"limit":  fmt.Sprintf("%d", limit),
		"offset": fmt.Sprintf("%d", offset),
	}
	var locations []Location
	err := getRequest(p, "location-area", params, &locations)
	if err != nil {
		return []Location{}, fmt.Errorf("Error getting location: %w", err)
	}

	return locations, nil
}

func getRequest[T any](p *pokeapi, path string, qp queryParams, dataContainer *[]T) error {
	fullUrl := p.baseUrl + "/" + path + qp.String()

	request, err := http.NewRequest(http.MethodGet, fullUrl, nil)
	if err != nil {
		return fmt.Errorf("Erro ao montar request: %w", err)
	}
	res, err := p.client.Do(request)

	if err != nil {
		return fmt.Errorf("Erro ao chamar GET %s: %w", fullUrl, err)
	}
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("Erro ao processar GET %s: %w", fullUrl, err)
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
