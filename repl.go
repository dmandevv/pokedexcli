package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/dmandevv/pokedexcli/internal/pokeapi"
)

type config struct {
	pokeapiClient        pokeapi.Client
	nextLocationsURL     *string
	previousLocationsURL *string
	pokedex              map[string]pokeapi.Pokemon
}

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		cleanText := cleanInput(scanner.Text())
		if len(cleanText) == 0 {
			continue
		}
		args := []string{}
		if len(cleanText) > 1 {
			args = cleanText[1:]
		}
		command, exists := getCommands()[cleanText[0]]
		if exists {
			err := command.callback(cfg, args...)
			if err != nil {
				fmt.Println("Error:", err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
		}
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
	callback    func(*config, ...string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
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
		"explore": {
			name:        "explore",
			description: "List all pokemon in a specific area",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch <pokemon-name>",
			description: "Attempt to catch a pokemon by name",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect <pokemon-name>",
			description: "Inspect a caught pokemon by name",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "List all caught pokemon",
			callback:    commandPokedex,
		},
	}
}
