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
		param := ""
		if len(cleanText) > 1 {
			param = cleanText[1]
		}
		command, exists := getCommands()[cleanText[0]]
		if exists {
			err := command.callback(cfg, param)
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
	callback    func(*config, string) error
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
	}
}
