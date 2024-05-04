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
