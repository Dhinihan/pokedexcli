package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/Dhinihan/pokedexcli/internal/pokeapicache"
)

const realUrl = "https://pokeapi.co/api/v2/"

type Api struct {
	client  *http.Client
	baseUrl string
	cache   *pokeapicache.Cache
}

func NewPokeapi() *Api {
	return &Api{
		client:  http.DefaultClient,
		baseUrl: realUrl,
		cache:   pokeapicache.NewCache(1 * time.Hour),
	}
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
