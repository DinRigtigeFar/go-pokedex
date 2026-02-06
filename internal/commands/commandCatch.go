package commands

import (
	"fmt"
	"math/rand"
	// "encoding/json"
	// "io"
	// "net/http"
)

func CommandCatch(cfg *Config) error {
	resp, err := cfg.PokeapiClient.LocationCatch(cfg.Pokemon)
	if err != nil {
		return err
	}
	
	found, err := cfg.PokeapiClient.CheckIfFoundInCurrentLocation(resp.LocationAreaEncounters, cfg.ExploreLocation)
	
	if err != nil {
		return err
	}

	if !found {
		fmt.Printf("%s is not found in the current location: %s\n", cfg.Pokemon, cfg.ExploreLocation)
		fmt.Printf("Use the 'explore' command to list the Pokemon at this location!\n")
		return nil
	}

	//fmt.Println(resp) // debug

	fmt.Printf("Throwing a Pokeball at %s...\n", cfg.Pokemon)
	if catch := calculateCatch(resp.BaseExperience); catch {
		fmt.Printf("%s was caught!\n", cfg.Pokemon)
		cfg.Pokedex[cfg.Pokemon] = resp
	} else {
		fmt.Printf("%s escaped!\n", cfg.Pokemon)
	}
	return nil
}

func calculateCatch(baseExp int) bool {
	first := rand.Float64() / float64(baseExp) * 100
	second := 1 / (first * first)
	compare := rand.Float64() * 100
	return second < compare
}
