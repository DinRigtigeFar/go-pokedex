package commands

import (
	"github.com/DinRigtigeFar/pokedexcli/internal/pokeapi"
)
type Config struct {
	PokeapiClient    pokeapi.Client
	NextLocationsURL *string
	PrevLocationsURL *string
	ExploreLocation  string
	Pokemon          string
	Pokedex          map[string]pokeapi.Pokemon
}


type cliCommand struct {
	name        string
	description string
	Callback    func(*Config) error
}

func GetCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			Callback:    CommandHelp,
		},
		"map": {
			name:        "map",
			description: "Get the next page of locations",
			Callback:    CommandMapf,
		},
		"mapb": {
			name:        "mapb",
			description: "Get the previous page of locations",
			Callback:    CommandMapb,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			Callback:    CommandExit,
		},
		"explore": {
			name:        "explore",
			description: "Explore the location",
			Callback:    CommandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Try to catch a Pokemon", // in the current location", TODO: add current location to config and use it in the description
			Callback:    CommandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Inspect Pokedex entry for Pokemon", // in the current location", TODO: add current location to config and use it in the description
			Callback:    CommandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "List all caught Pokemon",
			Callback:    CommandPokedex,
		},
	}
}