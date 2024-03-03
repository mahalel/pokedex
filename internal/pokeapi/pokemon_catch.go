package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// GetPokemonStats -
func (c *Client) GetPokemonStats(pokemonName string) (RespPokemonStats, error) {
	url := fmt.Sprintf("%s/pokemon/%s", baseURL, pokemonName)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespPokemonStats{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespPokemonStats{}, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespPokemonStats{}, err
	}

	pokemonStatsResp := RespPokemonStats{}
	err = json.Unmarshal(dat, &pokemonStatsResp)
	if err != nil {
		return RespPokemonStats{}, err
	}

	return pokemonStatsResp, nil
}
