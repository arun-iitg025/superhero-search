package models

type Movie struct {
	Title       string `json:"title"`
	ReleaseYear string `json:"release_year"`
	Genre       string `json:"genre"`
	Rating      string `json:"rating"`
	Poster      string `json:"poster"`
}
