package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocationAreas(pageURL *string) (LocationAreaNames, error) {
	var fullURL string
	if pageURL != nil {
		fullURL = *pageURL
	} else {
		endpoint := "/location-area"
		fullURL = baseURL + endpoint
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

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreaNames{}, err
	}

	locationAreaNames := LocationAreaNames{}
	err = json.Unmarshal(dat, &locationAreaNames)
	if err != nil {
		return LocationAreaNames{}, err
	}

	return locationAreaNames, nil
}
