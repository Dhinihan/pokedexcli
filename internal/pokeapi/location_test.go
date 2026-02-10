package pokeapi

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func makeLocationResponse(locs []Location) LocationResponse {
	return LocationResponse{0, "", "", locs}
}

func makeLocationDetailsResponse(pl []Pokemon, loc Location) LocationDetailsResponse {
	encounters := []PokemonEncounter{}
	for _, p := range pl {
		encounters = append(encounters, PokemonEncounter{p})
	}
	return LocationDetailsResponse{loc.Name, loc.URL, encounters}
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
			expRes := makeLocationDetailsResponse(tc.expPList, tc.inLoc)
			pa := mockServer(tt, "/location-area/"+tc.inLoc.Name, expRes)

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
			pa := mockServer(
				tt,
				fmt.Sprintf(
					"/location-area?limit=%d&offset=%d",
					tc.limit,
					tc.offset,
				),
				makeLocationResponse(tc.expLoc),
			)

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
