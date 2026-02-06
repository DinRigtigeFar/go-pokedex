package commands

import (
	"fmt"
)

func CommandPokedex(cfg *Config) error {
	if len(cfg.Pokedex) == 0 {
		fmt.Println("Your Pokedex is empty! Use the 'catch' command to catch some Pokemon first.")
		return nil
	}

	fmt.Println("Your Pokedex:")
	for name := range cfg.Pokedex {
		fmt.Printf("- %s\n", name)
	}
	return nil
}