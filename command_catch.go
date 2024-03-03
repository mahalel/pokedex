package main

import (
	"errors"
	"fmt"
	"math"
	"math/rand"
)

// catchProbability -
func catchProbability(baseExperience, k float64) float64 {
	return 1 / (1 + math.Exp(-k*baseExperience))
}

// commandCatch -
func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("no pokemon given")
	}
	pokemonName := args[0]

	pokeDex := map[string]Pokemon{}

	catchPokemonResp, err := cfg.pokeapiClient.GetPokemonStats(pokemonName)
	if err != nil {
		return err
	}
	// Base experience of the Pokémon
	baseExperience := float64(catchPokemonResp.BaseExperience)
	// Constant 'k' determines the steepness of the curve
	k := 0.005

	// Calculate catch probability
	probability := catchProbability(baseExperience, k)
	// Generate a random number between 0 and 1
	random := rand.Float64()

	// Check if the Pokémon is caught
	if random >= probability {
		fmt.Println("Congratulations! You caught the Pokémon!")
		pokeDex[catchPokemonResp.Name] = Pokemon{
			Name:   catchPokemonResp.Name,
			Height: catchPokemonResp.Height,
			Weight: catchPokemonResp.Weight,
			Stats:  catchPokemonResp.Stats,
			Types:  catchPokemonResp.Types,
		}
	} else {
		fmt.Println("Oh no! The Pokémon escaped!")
	}
	return nil
}

type Pokemon struct {
	Name   string `json:"name"`
	Height int    `json:"height"`
	Weight int    `json:"weight"`
	Stats  []struct {
		BaseStat int `json:"base_stat"`
		Effort   int `json:"effort"`
		Stat     struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		} `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Slot int `json:"slot"`
		Type struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		} `json:"type"`
	} `json:"types"`
}
