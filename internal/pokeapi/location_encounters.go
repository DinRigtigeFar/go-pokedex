package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
	"errors"
	// "fmt"
)

// LocatioEncounters -
func (c *Client) LocationEncounters(area string) (RespEncounters, error) {
	
	url := baseURL + "/location-area/"
	if area == "" {
		return RespEncounters{}, errors.New("No location area specified")
	}
	url = url + area
	
	// Check cache first
	if c.cache != nil {
		if cachedData, found := c.cache.Get(url); found {
			locationsResp := RespEncounters{}
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
		return RespEncounters{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespEncounters{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespEncounters{}, err
	}

	// Store in cache
	if c.cache != nil {
		c.cache.Add(url, dat)
	}

	encountersResp := RespEncounters{}
	err = json.Unmarshal(dat, &encountersResp)
	if err != nil {
		return RespEncounters{}, err
	}

	return encountersResp, nil
}
