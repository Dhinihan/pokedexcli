package pokeapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"testing"
	"time"

	"github.com/Dhinihan/pokedexcli/internal/pokeapicache"
)

func isSlice(v any) bool {
	return reflect.TypeOf(v).Kind() == reflect.Slice
}

func mockServer(th *testing.T, expPath string, jsonable any) *Api {
	th.Helper()
	hf := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		u, err := url.Parse(expPath)

		if err != nil {
			fmt.Printf("Warning: %s", err.Error())
		}

		if r.URL.Path != u.Path {
			th.Errorf(
				"Chamou %s ao invés de %s",
				r.URL.Path,
				expPath,
			)
		}

		if u.Query().Encode() != "" && r.URL.Query().Encode() != u.Query().Encode() {
			th.Errorf(
				"Chamou %s ao invés de %s",
				r.URL.Query().Encode(),
				u.Query().Encode(),
			)
		}

		data, err := json.Marshal(jsonable)

		if err != nil {
			fmt.Printf("Warning: %s", err.Error())
		}
		_, err = w.Write(data)
		if err != nil {
			fmt.Printf("Warning: %s", err.Error())
		}
	})
	ts := httptest.NewServer(hf)
	th.Cleanup(func() {
		ts.Close()
	})

	return &Api{
		ts.Client(),
		ts.URL,
		pokeapicache.NewCache(1 * time.Second),
	}
}
