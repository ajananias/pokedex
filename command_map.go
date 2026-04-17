package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Result struct {
	Name string
	Url string
}

type Page struct {
	Count int
	Next *string
	Previous *string
	Results []Result
}

func commandMap(c *config) error {
	url := ""
	if c.nextLocationsURL == nil {
		url = "https://pokeapi.co/api/v2/location-area/?offset=0&limit=20"
	} else {
		url = *c.nextLocationsURL
	}
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
	}
	// unmarshal the JSON of the page
	var page Page
	if err := json.Unmarshal(body, &page); err != nil {
		log.Fatalf("Error unmarshalling JSON: %v", err)
	}
	// store the locations in the config struct
	c.prevLocationsURL = page.Previous
	c.nextLocationsURL = page.Next

	for i := 0; i < len(page.Results); i++ {
		fmt.Println(page.Results[i].Name)
	}

	return nil
}
