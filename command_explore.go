package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type Pokemon struct {
	Name string
	Url  string
}
type PokemonEncounters struct {
	Pokemon Pokemon `json:"pokemon"`
}
type AreaPage struct {
	PokemonEncounters []PokemonEncounters `json:"pokemon_encounters"`
}

func commandExplore(c *config, parameters string) error {
	url := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%s/", parameters)
	var body []byte
	cache, exists := c.pokeCache.Get(url)
	if exists {
		body = cache
	} else {
		res, err := http.Get(url)
		if err != nil {
			return errors.New("Couldn't fetch data from GET request:" + err.Error())
		}
		body, err = io.ReadAll(res.Body)
		res.Body.Close()
		if res.StatusCode > 299 {
			return errors.New("Unable to find location area:\nStatus Code: " + fmt.Sprint(res.StatusCode) + "\nBody: " + string(body))
		}
		if err != nil {
			return errors.New("Error reading response body" + err.Error())
		}
		c.pokeCache.Add(url, body)
	}

	var areaPage AreaPage
	if err := json.Unmarshal(body, &areaPage); err != nil {
		return errors.New("error unmarshalling JSON: " + err.Error())
	}
	fmt.Println("Exploring " + parameters + "...\nFound Pokemon:")

	for i := 0; i < len(areaPage.PokemonEncounters); i++ {
		fmt.Println("- " + areaPage.PokemonEncounters[i].Pokemon.Name)
	}
	return nil
}
