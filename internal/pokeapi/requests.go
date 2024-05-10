package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func UnmarshalAreaNames(dat []byte) (LocationAreaNames, error) {
	locationAreaNames := LocationAreaNames{}
	err := json.Unmarshal(dat, &locationAreaNames)
	if err != nil {
		return LocationAreaNames{}, err
	}
	return locationAreaNames, nil
}

func UnmarshalArea(dat []byte) (LocationArea, error) {
	locationArea := LocationArea{}
	err := json.Unmarshal(dat, &locationArea)
	if err != nil {
		return LocationArea{}, err
	}
	return locationArea, nil
}

func UnmarshalPokemon(dat []byte) (PokemonInfo, error) {
	Pokemon := PokemonInfo{}
	err := json.Unmarshal(dat, &Pokemon)
	if err != nil {
		return PokemonInfo{}, err
	}
	return Pokemon, nil

}

func (c *Client) ListLocationAreas(pageURL *string) (LocationAreaNames, error) {
	var fullURL string
	if pageURL != nil {
		fullURL = *pageURL
	} else {
		endpoint := "/location-area"
		fullURL = baseURL + endpoint
	}

	dat, ok := c.cache.Get(fullURL)
	if ok {
		return UnmarshalAreaNames(dat)
	}

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return LocationAreaNames{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreaNames{}, err
	}

	if resp.StatusCode > 399 {
		return LocationAreaNames{}, fmt.Errorf("bad status code: %v", resp.StatusCode)
	}

	dat, err = io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreaNames{}, err
	}

	locationAreaNames, err := UnmarshalAreaNames(dat)
	if err != nil {
		return LocationAreaNames{}, err
	}

	c.cache.Add(fullURL, dat)

	return locationAreaNames, nil
}

func (c *Client) GetLocationArea(AreaName string) (LocationArea, error) {
	endpoint := "/location-area/" + AreaName

	fullURL := baseURL + endpoint

	dat, ok := c.cache.Get(fullURL)
	if ok {
		return UnmarshalArea(dat)
	}

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return LocationArea{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationArea{}, err
	}

	if resp.StatusCode > 399 {
		return LocationArea{}, fmt.Errorf("bad status code: %v", resp.StatusCode)
	}

	dat, err = io.ReadAll(resp.Body)
	if err != nil {
		return LocationArea{}, err
	}

	locationArea, err := UnmarshalArea(dat)
	if err != nil {
		return LocationArea{}, err
	}

	c.cache.Add(fullURL, dat)

	return locationArea, nil
}

func (c *Client) GetPokemonInfo(PokemonName string) (PokemonInfo, error) {

	endpoint := "/pokemon/" + PokemonName
	fullUrl := baseURL + endpoint

	req, err := http.NewRequest("GET", fullUrl, nil)
	if err != nil {
		return PokemonInfo{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return PokemonInfo{}, err
	}

	if resp.StatusCode > 399 {
		return PokemonInfo{}, fmt.Errorf("bad status code: %v", resp.StatusCode)
	}

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return PokemonInfo{}, err
	}

	Pokemon, err := UnmarshalPokemon(dat)
	if err != nil {
		return PokemonInfo{}, err
	}

	c.cache.Add(fullUrl, dat)
	return Pokemon, nil

}
