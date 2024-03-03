package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// GetLocationArea -
func (c *Client) GetLocationArea(locationName string) (RespLocationArea, error) {
	url := fmt.Sprintf("%s/location-area/%s", baseURL, locationName)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespLocationArea{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespLocationArea{}, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespLocationArea{}, err
	}

	areaExploreResp := RespLocationArea{}
	err = json.Unmarshal(dat, &areaExploreResp)
	if err != nil {
		return RespLocationArea{}, err
	}

	return areaExploreResp, nil
}
