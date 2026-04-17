package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/ajananias/pokedex/internal/pokeapi"
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

		firstWord := cleanUserInput[0]

		command, exists := getCommands()[firstWord]
		if exists {
			err := command.callback(cfg)
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
	callback    func(cfg *config) error
}
type config struct {
	pokeapiClient    pokeapi.Client
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
	}
}
