package pokeapi

import (
	"io"
	"log"
	"net/http"
)

func GetLocations(url string) ([]byte, error) {
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
	// fmt.Printf("%s", body)

	return body, nil
}
