package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func commandMapBack(c *config, parameters string) error {
	url := ""
	if c.prevLocationsURL == nil {
		fmt.Println("you're on the first page")
		return nil
	} else {
		url = *c.prevLocationsURL
	}

	var body []byte
	cache, exists := c.pokeCache.Get(url)
	if exists {
		body = cache
	} else {
		res, err := http.Get(url)
		if err != nil {
			log.Fatal(err)
		}
		body, err = io.ReadAll(res.Body)
		res.Body.Close()
		if res.StatusCode > 299 {
			log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
		}
		if err != nil {
			log.Fatal(err)
		}
		c.pokeCache.Add(url, body)
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
