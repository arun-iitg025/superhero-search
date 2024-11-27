package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const omdbAPIBase = "http://www.omdbapi.com/?apikey=a3e8c90a"

func FetchMovieData(movieTitle string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s&t=%s", omdbAPIBase, movieTitle)
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
