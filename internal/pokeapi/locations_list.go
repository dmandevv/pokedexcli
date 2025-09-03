package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

type LocationList struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func (c Client) GetLocationsList(pageURL *string) (LocationList, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationList{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationList{}, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationList{}, err
	}

	var locations LocationList
	err = json.Unmarshal(data, &locations)
	if err != nil {
		return LocationList{}, err
	}

	return locations, nil
}
