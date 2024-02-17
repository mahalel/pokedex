package main

import (
	"encoding/json"
	"fmt"

	"github.com/mahalel/pokedex/internal/pokeapi"
)

var poke_api_next string = "https://pokeapi.co/api/v2/location"
var poke_api_prev string

func commandMap() error {
	if poke_api_next == "" {
		fmt.Println("No more locations available, please go backwards.")
		return nil
	}

	res, err := pokeapi.GetLocations(poke_api_next)
	if err != nil {
		return err
	}

	cnf := Config{}
	err = json.Unmarshal(res, &cnf)
	if err != nil {
		return err
	}

	for _, i := range cnf.Results {
		fmt.Println(i.Name)
	}

	if cnf.Next != nil {
		poke_api_next = *cnf.Next
	} else {
		poke_api_next = ""
	}

	if cnf.Previous != nil {
		poke_api_prev = *cnf.Previous
	} else {
		poke_api_prev = ""
	}

	return nil
}

func commandMapBack() error {
	if poke_api_prev == "" {
		fmt.Println("No more locations available, please go forward.")
		return nil
	}

	res, err := pokeapi.GetLocations(poke_api_prev)
	if err != nil {
		return err
	}

	cnf := Config{}
	err = json.Unmarshal(res, &cnf)
	if err != nil {
		return err
	}

	for _, i := range cnf.Results {
		fmt.Println(i.Name)
	}

	poke_api_next = poke_api_prev

	if cnf.Previous != nil {
		poke_api_prev = *cnf.Previous
	} else {
		poke_api_prev = ""
	}

	return nil
}

type Config struct {
	Count    int       `json:"count"`
	Next     *string   `json:"next"`
	Previous *string   `json:"previous"`
	Results  []Results `json:"results"`
}
type Results struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
