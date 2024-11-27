package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"superhero-search/cache"
	"superhero-search/models"
)

// SearchHandler handles both superhero and movie search queries.
func SearchHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("query")
	if query == "" {
		http.Error(w, "Query parameter is required", http.StatusBadRequest)
		return
	}

	// Handle 'type:movie' query parameter
	if strings.Contains(query, "type:movie") {
		query = strings.Replace(query, " type:movie", "", -1) // Clean the query to search for the movie only
		movieData, err := searchMovies(query)
		if err != nil {
			log.Println("Error fetching movie data:", err)
			http.Error(w, "Failed to fetch movie data", http.StatusInternalServerError)
			return
		}
		// Return the movie data as JSON
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(movieData)
		return
	}

	// Search for superhero data
	superheroData, err := searchSuperheroes(query)
	if err != nil {
		log.Println("Error fetching superhero data:", err)
		http.Error(w, "Failed to fetch superhero data", http.StatusInternalServerError)
		return
	}

	// Return superhero data as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(superheroData)
}

// searchMovies fetches movie data from OMDB API.
func searchMovies(query string) ([]models.Movie, error) {
	// First check Redis cache
	cachedMovies, err := cache.GetCache(query)
	if err == nil && cachedMovies != "" {
		// If data exists in cache, return it
		var movies []models.Movie
		err := json.Unmarshal([]byte(cachedMovies), &movies)
		if err == nil {
			return movies, nil
		}
	}

	// If no cached data, fetch from OMDB API
	url := fmt.Sprintf("http://www.omdbapi.com/?apikey=a3e8c90a&s=%s", query)
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode != http.StatusOK {
		log.Println("Error fetching movie data from OMDB:", err)
		return nil, fmt.Errorf("failed to fetch movie data")
	}
	defer resp.Body.Close()

	var apiResponse map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&apiResponse); err != nil {
		log.Println("Error decoding OMDB API response:", err)
		return nil, err
	}

	searchResults, ok := apiResponse["Search"].([]interface{})
	if !ok || len(searchResults) == 0 {
		log.Println("No movies found for query:", query)
		return nil, fmt.Errorf("no movies found")
	}

	// Process up to 5 movies from the search results
	var movies []models.Movie
	for i, result := range searchResults {
		if i >= 15 { // Limit to 5 movies
			break
		}
		movieMap := result.(map[string]interface{})
		movie := models.Movie{
			Title:       movieMap["Title"].(string),
			ReleaseYear: movieMap["Year"].(string),
			Genre:       "Not available", // Add detailed API calls for more data
			Rating:      "N/A",
			Poster:      movieMap["Poster"].(string),
		}
		movies = append(movies, movie)
	}

	// Cache the results before returning
	movieDataJSON, _ := json.Marshal(movies)
	cache.SetCache(query, string(movieDataJSON))

	return movies, nil
}

// searchSuperheroes fetches superhero data from the Superhero API.
func searchSuperheroes(query string) ([]models.Superhero, error) {
	// First check Redis cache
	cachedHeroes, err := cache.GetCache(query)
	if err == nil && cachedHeroes != "" {
		// If data exists in cache, return it
		var superheroes []models.Superhero
		err := json.Unmarshal([]byte(cachedHeroes), &superheroes)
		if err == nil {
			return superheroes, nil
		}
	}

	// If no cached data, fetch from Superhero API
	apiURL := fmt.Sprintf("https://superheroapi.com/api/2806e8bc5db394883f56c9e18bd846fe/search/%s", query)
	resp, err := http.Get(apiURL)
	if err != nil || resp.StatusCode != http.StatusOK {
		log.Println("Error calling Superhero API:", err)
		return nil, fmt.Errorf("failed to fetch data from Superhero API")
	}
	defer resp.Body.Close()

	// Parse Superhero API response
	var apiResponse map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&apiResponse); err != nil {
		log.Println("Error parsing API response:", err)
		return nil, err
	}

	results, ok := apiResponse["results"].([]interface{})
	if !ok || len(results) == 0 {
		return nil, fmt.Errorf("no results found in Superhero API response")
	}

	// Process the results into a list of superheroes
	var heroes []models.Superhero
	for _, result := range results {
		data := result.(map[string]interface{})
		powers := []string{}

		if powerstats, ok := data["powerstats"].(map[string]interface{}); ok {
			for key, value := range powerstats {
				powers = append(powers, fmt.Sprintf("%s: %v", key, value))
			}
		}

		// Get superhero image URL
		imageURL := ""
		if image, ok := data["image"].(map[string]interface{}); ok {
			imageURL = image["url"].(string)
		}

		heroes = append(heroes, models.Superhero{
			Name:   data["name"].(string),
			Powers: powers,
			Movies: []string{"Movies not provided by the API"}, // Placeholder for movie data
			Image:  imageURL,                                   // Added image URL
		})
	}

	// Cache the results before returning
	heroDataJSON, _ := json.Marshal(heroes)
	cache.SetCache(query, string(heroDataJSON))

	return heroes, nil
}
