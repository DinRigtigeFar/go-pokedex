package commands

import (
	"fmt"
	"errors"
)

func CommandMapf(cfg *Config) error {
	locationsresp, err := cfg.PokeapiClient.ListLocations(cfg.NextLocationsURL)
	if err != nil {
		return err
	}
	cfg.NextLocationsURL = locationsresp.Next
	cfg.PrevLocationsURL = locationsresp.Previous
	
	fmt.Println()
	fmt.Println("Locations:")
	for _, location := range locationsresp.Results {
		fmt.Println(location.Name)
	}
	return nil	
}

func CommandMapb(cfg *Config) error {
	if cfg.PrevLocationsURL == nil {
		return errors.New("you're on the first page")
	}
	
	locationsresp, err := cfg.PokeapiClient.ListLocations(cfg.PrevLocationsURL)
	if err != nil {
		return err
	}
	cfg.NextLocationsURL = locationsresp.Next
	cfg.PrevLocationsURL = locationsresp.Previous
	
	fmt.Println()
	fmt.Println("Locations:")
	for _, location := range locationsresp.Results {
		fmt.Println(location.Name)
	}
	return nil	
}
