package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/dmandevv/pokedexcli/internal/pokeapi"
	"github.com/dmandevv/pokedexcli/internal/pokecache"
)

type config struct {
	pokeapiClient        pokeapi.Client
	nextLocationsURL     *string
	previousLocationsURL *string
}

func startRepl(cfg *config, cache *pokecache.Cache) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		cleanText := cleanInput(scanner.Text())
		if len(cleanText) == 0 {
			continue
		}

		command, exists := getCommands()[cleanText[0]]
		if exists {
			err := command.callback(cfg, cache)
			if err != nil {
				fmt.Println("Error:", err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
		}
		//fmt.Print("Your command was: ", cleanText[0], "\n")
	}
}

func cleanInput(text string) []string {
	lowered := strings.ToLower(text)
	trimmed := strings.Fields(lowered)
	return trimmed
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, *pokecache.Cache) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays next 20 areas",
			callback:    commandMapf,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays previous 20 areas",
			callback:    commandMapb,
		},
	}
}
