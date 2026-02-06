package commands

import (
	"fmt"
)
// TODO: add pagination support by adding parameter after "explore" command to dynamically fetch encounters for location
// TODO: Sanatize input
// TODO: Implement caching
// TODO: Refactor codebase to improve readability and maintainability by separating "command" functions into their own package
func CommandExplore(cfg *Config) error {
	resp, err := cfg.PokeapiClient.LocationEncounters(cfg.ExploreLocation)
	if err != nil {
		return err
	}
	fmt.Printf("Exploring %s...\n", resp.LocationName)
	fmt.Println("Found Pokemon:")
	for _, encounter := range resp.Encounters {
		fmt.Printf("- %s\n", encounter.Pokemon.Name)
	}
	return nil
}
