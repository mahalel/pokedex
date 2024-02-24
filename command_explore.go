package main

import (
	"errors"
	"fmt"
)

// commandExplore -
func commandExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("no location area given")
	}
	locationAreaName := args[0]

	exploreAreaResp, err := cfg.pokeapiClient.GetLocationArea(locationAreaName)
	if err != nil {
		return err
	}
	for pokemon := range exploreAreaResp.PokemonEncounters {
		pokemonName := exploreAreaResp.PokemonEncounters[pokemon].Pokemon.Name
		fmt.Printf("- %s\n", pokemonName)
	}
	return nil
}
