package pokeapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/Dhinihan/pokedexcli/internal/pokeapicache"
	"github.com/google/go-cmp/cmp"
)

func locationsToBytes(locs []Location) ([]byte, error) {
	response := LocationResponse{
		Results: locs,
	}
	data, err := json.Marshal(response)
	if err != nil {
		return []byte{}, fmt.Errorf("Error parsing locations to json: %w", err)
	}
	return data, nil
}

func locationDetailsToBytes(pl []Pokemon, loc Location) ([]byte, error) {
	encounters := []PokemonEncounter{}
	for _, p := range pl {
		encounters = append(encounters, PokemonEncounter{p})
	}
	response := LocationDetailsResponse{loc.Name, loc.URL, encounters}

	data, err := json.Marshal(response)
	if err != nil {
		return []byte{}, fmt.Errorf("Error parsing location detail to json: %w", err)
	}
	return data, nil
}

func TestGetLocationDetails(t *testing.T) {
	pList1 := []Pokemon{}
	pList2 := []Pokemon{{"Bulbasauro", "url1"}}
	pList3 := []Pokemon{{"Charmander", "url4"}, {"Squirtle", "url7"}}
	loc1 := Location{"local1", "url7"}
	loc2 := Location{"local2", "url8"}
	loc3 := Location{"local3", "url9"}
	type testCase struct {
		inLoc    Location
		expPList []Pokemon
	}
	testCases := []testCase{
		{loc1, pList1}, {loc2, pList2}, {loc3, pList3},
	}
	for _, tc := range testCases {
		t.Run(tc.inLoc.Name, func(tt *testing.T) {
			tt.Parallel()
			hf := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				expPath := "/location-area/" + tc.inLoc.Name
				if r.URL.Path != expPath {
					tt.Errorf(
						"Chamou %s ao invés de %s",
						r.URL.Path,
						expPath,
					)
				}
				data, err := locationDetailsToBytes(tc.expPList, tc.inLoc)
				if err != nil {
					fmt.Printf("Warning: %s", err.Error())
				}
				_, err = w.Write(data)
				if err != nil {
					fmt.Printf("Warning: %s", err.Error())
				}
			})
			ts := httptest.NewServer(hf)
			tt.Cleanup(func() {
				ts.Close()
			})

			pa := &Api{
				ts.Client(),
				ts.URL,
				pokeapicache.NewCache(1 * time.Second),
			}

			details, err := pa.GetLocationDetails(tc.inLoc.Name)
			if err != nil {
				tt.Errorf("UnexpectedError: %s", err.Error())
			}
			if len(details.PokemonEncounters) != len(tc.expPList) {
				tt.Errorf(
					"Expected %d encounters, received %d",
					len(details.PokemonEncounters),
					len(tc.expPList),
				)
			}
			for i, p := range details.PokemonEncounters {
				if diff := cmp.Diff(p.Pokemon, tc.expPList[i]); diff != "" {
					tt.Errorf("mismatch pokémon:\n%s", diff)
				}
			}
		})
	}
}
func TestGetLocations(t *testing.T) {
	loc1 := Location{
		Name: "Teste",
		URL:  "Teste",
	}
	loc2 := Location{
		Name: "Teste2",
		URL:  "Teste",
	}
	type testCase struct {
		limit  int
		offset int
		expLoc []Location
	}
	testCases := []testCase{
		{20, 1, []Location{loc2}},
		{20, 0, []Location{loc1, loc2}},
		{20, 2, []Location{}},
		{01, 0, []Location{loc1}},
	}

	for index, tc := range testCases {
		name := "limit: %d Offset: %d"
		t.Run(fmt.Sprintf(name, tc.limit, tc.offset), func(tt *testing.T) {
			tt.Parallel()
			hf := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if r.URL.Path != "/location-area" {
					tt.Errorf(
						"test %d, Chamou %s ao invés de /location-area",
						index,
						r.URL.Path,
					)
				}
				if r.URL.Query().Get("limit") != fmt.Sprintf("%d", tc.limit) {
					tt.Errorf(
						"Expected limit=%d, received %s",
						tc.limit,
						r.URL.Query().Get("limit"),
					)
				}
				if r.URL.Query().Get("offset") != fmt.Sprintf("%d", tc.offset) {
					tt.Errorf(
						"Expected offset=%d, received %s",
						tc.offset,
						r.URL.Query().Get("offset"),
					)
				}
				data, err := locationsToBytes(tc.expLoc)
				if err != nil {
					fmt.Printf("Warning: %s", err.Error())
				}
				_, err = w.Write(data)
				if err != nil {
					fmt.Printf("Warning: %s", err.Error())
				}
			})
			ts := httptest.NewServer(hf)
			tt.Cleanup(func() {
				ts.Close()
			})

			pa := &Api{
				ts.Client(),
				ts.URL,
				pokeapicache.NewCache(1 * time.Second),
			}

			locs, err := pa.GetLocation(tc.limit, tc.offset)
			if err != nil {
				tt.Errorf("UnexpectedError: %s", err.Error())
			}

			if len(locs) != len(tc.expLoc) {
				tt.Errorf(
					"test %d, received %d locations, expected %d",
					index,
					len(locs),
					len(tc.expLoc),
				)
			}
			for i, l := range locs {
				if diff := cmp.Diff(l, tc.expLoc[i]); diff != "" {
					tt.Errorf("test %d: mismatch locations:\n%s", index, diff)
				}
			}
		})
	}
}
