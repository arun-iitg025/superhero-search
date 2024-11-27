package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const superheroAPIBase = "https://superheroapi.com/api/2806e8bc5db394883f56c9e18bd846fe"

func FetchSuperheroData(query string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/search/%s", superheroAPIBase, query)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return result, nil
}
