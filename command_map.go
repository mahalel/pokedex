package main

func commandMap() error {
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
