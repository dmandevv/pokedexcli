package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

type EncountersList struct {
	Name       string `json:"name"`
	URL        string `json:"url"`
	Encounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

func (c Client) GetAreaPokemonList(areaName string) (EncountersList, error) {
	url := baseURL + "/location-area/" + areaName

	if data, found := c.cache.Get(areaName); found {
		encounters_list := EncountersList{}
		err := json.Unmarshal(data, &encounters_list)
		if err != nil {
			return encounters_list, err
		}
		return encounters_list, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return EncountersList{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return EncountersList{}, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return EncountersList{}, err
	}

	encounters_list := EncountersList{}
	err = json.Unmarshal(data, &encounters_list)
	if err != nil {
		return EncountersList{}, err
	}

	c.cache.Add(areaName, data)
	return encounters_list, nil
}
