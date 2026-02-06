package main

import (
	"time"
	"github.com/DinRigtigeFar/pokedexcli/internal/pokeapi"
	"github.com/DinRigtigeFar/pokedexcli/internal/commands"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second, 10 * time.Second)
	cfg := &commands.Config{
		PokeapiClient: pokeClient,
	}

	startRepl(cfg)
}














// package main

// import (
// 	"bufio"
// 	// "encoding/json"
// 	"fmt"
// 	"io"
// 	"net/http"
// 	"os"
// 	// "go-pokedex/internal/pokeapi"
// )

// type cliCommand struct {
// 	name        string
// 	description string
// 	callback    func() error
// 	config      []struct {
// 		Next     string `json:"next"`
// 		Previous string `json:"previous"`
// 	} `json:"config"`
// }



// func main() {
// 	commands := map[string]cliCommand{
// 		"exit": {
// 			name:        "exit",
// 			description: "Exit the Pokedex",
// 			callback:    commandExit,
// 		},
// 		"help": {
// 			name:        "help",
// 			description: "Display available commands",
// 			callback:    commandHelp,
// 		},
// 		"map": {
// 			name:        "map",
// 			description: "Display the map",
// 			callback:    commandMap,
// 		},
// 	}

// 	res, err := http.Get("/location-area/")
// 	if err != nil {
// 		fmt.Println("Error fetching data:", err)
// 		return
// 	}
// 	defer res.Body.Close()
// 	_, err = io.ReadAll(res.Body)
// 	if err != nil {
// 		fmt.Println("Error reading response body:", err)
// 		return
// 	}
// 	// err = json.Unmarshal(dat, &config)
// 	// if err != nil {
// 	// 	fmt.Println("Error decoding JSON:", err)
// 	// 	return
// 	// }


// 	scanner := bufio.NewScanner(os.Stdin)

// 	for {
// 		fmt.Print("Pokedex > ")
// 		scanner.Scan()
// 		input := scanner.Text()
// 		words := CleanInput(input)
// 		if len(words) != 0 {
// 			if command := commands[words[0]]; command.callback != nil {
// 				command.callback()
// 			} else {
// 				fmt.Printf("Unknown command")
// 			}
// 			// fmt.Println("Your command was:", words[0])
// 		} else {
// 			fmt.Println("No command entered.")
// 		}

// 	}
// }





