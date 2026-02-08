package pokeapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

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

			pa := &pokeapi{
				ts.Client(),
				ts.URL,
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
