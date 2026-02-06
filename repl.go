package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/DinRigtigeFar/pokedexcli/internal/commands"
	"github.com/DinRigtigeFar/pokedexcli/internal/pokeapi"
)

type config struct {
	pokeapiClient    pokeapi.Client
	nextLocationsURL *string
	prevLocationsURL *string
	exploreLocation  string
	pokemon          string
	pokedex          map[string]pokeapi.Pokemon
}

func startRepl(cfg *commands.Config) {
	cfg.Pokedex = make(map[string]pokeapi.Pokemon)
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]

		command, exists := commands.GetCommands()[commandName]
		if exists {
			switch commandName {
			case "explore":
				switch {
				case len(words) >= 2 && words[1] == ".":
					// do nothing, just use the current location
				case len(words) >= 2:
					cfg.ExploreLocation = words[1]
				case cfg.ExploreLocation != "":
					fmt.Printf("Please specify a location to explore.\nCurrrent location is: %s\n", cfg.ExploreLocation)
				default:
					fmt.Println("Please specify a location to explore")
					continue
				}
			case "catch":
				switch {
				case len(words) >= 2:
					cfg.Pokemon = words[1]
				default:
					if cfg.ExploreLocation != "" {
						fmt.Printf("Please specify a pokemon to catch.\nIt must be found in the current location: %s\nUse the 'explore' command to list the Pokemon at this location!\n", cfg.ExploreLocation)
					} else {
						fmt.Println("Use the 'explore' command to set a location before trying to catch a pokemon")
					}
					continue
				}
			case "inspect":
				switch {
				case len(words) >= 2:
					cfg.Pokemon = words[1]
				default:
					fmt.Println("Please specify a previously caught pokemon to inspect.\nUse the 'pokedex' command to list all caught Pokemon!")
					continue
				}

			}
			err := command.Callback(cfg)
			if err != nil {
				switch {
				case err.Error() == "No location area specified":
					fmt.Println("Please specify a location to explore")
				case err.Error() == "you're on the first page":
					fmt.Println("You're on the first page of locations")
				case commandName == "explore" && strings.Contains(err.Error(), "invalid character"):
					cfg.ExploreLocation = ""
					fmt.Println("Invalid location name. Please try again.")
				default:
					fmt.Printf("Error executing command: %s\n", err)
				}
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

// type CliCommand struct {
// 	name        string
// 	description string
// 	callback    func(*config) error
// }

// func getCommands() map[string]CliCommand {
// 	return map[string]CliCommand{
// 		"help": {
// 			name:        "help",
// 			description: "Displays a help message",
// 			callback:    commands.CommandHelp,
// 		},
// 		"map": {
// 			name:        "map",
// 			description: "Get the next page of locations",
// 			callback:    commands.CommandMapf,
// 		},
// 		"mapb": {
// 			name:        "mapb",
// 			description: "Get the previous page of locations",
// 			callback:    commands.CommandMapb,
// 		},
// 		"exit": {
// 			name:        "exit",
// 			description: "Exit the Pokedex",
// 			callback:    commands.CommandExit,
// 		},
// 		"explore": {
// 			name:        "explore",
// 			description: "Explore the location",
// 			callback:    commands.CommandExplore,
// 		},
// 		"catch": {
// 			name:        "catch",
// 			description: "Try to catch a Pokemon", // in the current location", TODO: add current location to config and use it in the description
// 			callback:    commands.CommandCatch,
// 		},
// 		"inspect": {
// 			name:        "inspect",
// 			description: "Inspect Pokedex entry for Pokemon", // in the current location", TODO: add current location to config and use it in the description
// 			callback:    commands.CommandInspect,
// 		},
// 		"pokedex": {
// 			name:        "pokedex",
// 			description: "List all caught Pokemon",
// 			callback:    commands.CommandPokedex,
// 		},
// 	}
// }
