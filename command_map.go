package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func commandMap() error {

	res, err := http.Get("http://www.google.com/robots.txt")
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	cnf := Config{}
	err = json.Unmarshal(body, &cnf)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(cnf)

	// make a call to an api and display the result
	// 20 locations at a time

	// if you're on the last page, next should print an error
	// at the end next = null
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
