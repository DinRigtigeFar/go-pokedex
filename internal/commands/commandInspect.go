package commands

import (
	"fmt"
)

func CommandInspect(cfg *Config) error {
	if len(cfg.Pokedex) == 0 {
		fmt.Println("Your Pokedex is empty! Use the 'catch' command to catch some Pokemon first.")
		return nil
	}

	dat, ok := cfg.Pokedex[cfg.Pokemon]
	if !ok {
		fmt.Printf("%s is not in your Pokedex! Use the 'catch' command to catch it first.\n", cfg.Pokemon)
		return nil
	}

	fmt.Printf("Name: %s\n", dat.Name)
	fmt.Printf("Height: %d\n", dat.Height)
	fmt.Printf("Weight: %d\n", dat.Weight)
	fmt.Println("Stats:")
	for _, stat := range dat.Stats {
		fmt.Printf("  -%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, typ := range dat.Types {
		fmt.Printf("  - %s\n", typ.Type.Name)
	}
	return nil
}