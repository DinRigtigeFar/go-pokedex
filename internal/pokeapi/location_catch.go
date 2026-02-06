package pokeapi

import (
	"encoding/json"
	// "fmt"
	"io"
	"net/http"
	// "errors"
	// "fmt"
	// "math/rand"
	// "fmt"
)

func (c *Client) LocationCatch(pokemon string) (Pokemon, error) {
	url := baseURL + "/pokemon/" + pokemon

	if c.cache != nil {
		if cachedData, found := c.cache.Get(url); found {
			locationsResp := Pokemon{}
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
		return Pokemon{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return Pokemon{}, err
	}

	// Store in cache
	if c.cache != nil {
		c.cache.Add(url, dat)
	}

	catchResp := Pokemon{}
	err = json.Unmarshal(dat, &catchResp)
	if err != nil {
		return Pokemon{}, err
	}

	return catchResp, nil
}

func (c *Client) CheckIfFoundInCurrentLocation(foundLocationsURL, currentLocation string) (bool, error) {
	if c.cache != nil {
		if cachedData, found := c.cache.Get(foundLocationsURL); found {
			foundAtLocationsResp := RespFoundAtLocations{}
			err := json.Unmarshal(cachedData, &foundAtLocationsResp)
			if err == nil {
				for _, loc := range foundAtLocationsResp {
					if loc.LocationArea.Name == currentLocation {
						return true, nil
					}
				}
			}
			// If unmarshalling fails, proceed to fetch from API
		}
	}
	
	
	req, err := http.NewRequest("GET", foundLocationsURL, nil)
	if err != nil {
		return false, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}

	// Store in cache
	if c.cache != nil {
		c.cache.Add(foundLocationsURL, dat)
	}

	foundAtLocationsResp := RespFoundAtLocations{}

	err = json.Unmarshal(dat, &foundAtLocationsResp)
	if err != nil {
		return false, err
	}
	for _, loc := range foundAtLocationsResp {
		if loc.LocationArea.Name == currentLocation {
			return true, nil
		}
	}
	return false, nil
}