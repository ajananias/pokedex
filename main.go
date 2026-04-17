package main

import (
	"time"

	"github.com/ajananias/pokedex/internal/pokeapi"
)

func main() {
	cfg := &config{
		pokeapiClient: pokeapi.NewClient(5 * time.Second),
	}
	startLoop(cfg)
}
