package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

// ListLocations -
func (c *Client) ListLocations(pageURL *string) (RespShallowLocations, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}
	// Check cache first
	if c.cache != nil {
		if cachedData, found := c.cache.Get(url); found {
			locationsResp := RespShallowLocations{}
			err := json.Unmarshal(cachedData, &locationsResp)
			if err == nil {
				return locationsResp, nil
			}
			// If unmarshalling fails, proceed to fetch from API
		}
	}

	// Fetch from API
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespShallowLocations{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespShallowLocations{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespShallowLocations{}, err
	}

	// Store in cache
	if c.cache != nil {
		c.cache.Add(url, dat)
	}

	locationsResp := RespShallowLocations{}
	err = json.Unmarshal(dat, &locationsResp)
	if err != nil {
		return RespShallowLocations{}, err
	}

	return locationsResp, nil
}
