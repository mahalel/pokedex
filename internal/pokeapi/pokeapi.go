package pokeapi

import (
	"io"
	"log"
	"net/http"
)

const poke_api_location string = "https://pokeapi.co/api/v2/location"

func getUrl() []byte {
	res, err := http.Get(poke_api_location)
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

	return body
}

func Unmarshal(data []byte, v interface{}) error {

}
