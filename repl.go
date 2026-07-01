package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/ajananias/pokedex/internal/pokeapi"
	"github.com/ajananias/pokedex/internal/pokecache"
)

func startLoop(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		userInput := scanner.Text()
		cleanUserInput := cleanInput(userInput)

		if len(cleanUserInput) == 0 {
			continue
		}
		if len(cleanUserInput) > 2 {
			fmt.Println("Invalid input. Use only one parameter after the command.")
			continue
		}

		optionalParameter := ""
		if len(cleanUserInput) == 2 {
			optionalParameter = cleanUserInput[1]
		}
		firstWord := cleanUserInput[0]
		command, exists := getCommands()[firstWord]
		if exists {
			err := command.callback(cfg, optionalParameter)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

func cleanInput(text string) []string {
	cleanList := []string{}
	for word := range strings.FieldsSeq(text) {
		cleanList = append(cleanList, strings.ToLower(word))
	}
	return cleanList
}

type cliCommand struct {
	name        string
	description string
	callback    func(cfg *config, parameters string) error
}
type config struct {
	pokeapiClient    pokeapi.Client
	pokeCache        pokecache.Cache
	nextLocationsURL *string
	prevLocationsURL *string
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
			description: "Display the names of the next 20 location areas in the Pokemon world",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Display the names of the previous 20 location areas in the Pokemon world",
			callback:    commandMapBack,
		},
		"explore": {
			name:        "explore",
			description: "Display all pokemon located in the specified location area.",
			callback:    commandExplore,
		},
	}
}
