package models

// Superhero represents a superhero object with details like name, powers, and image
type Superhero struct {
	Name   string   `json:"name" bson:"name"`
	Powers []string `json:"powers" bson:"powers"`
	Movies []string `json:"movies" bson:"movies"`
	Image  string   `json:"image" bson:"image"` // Add the Image field to store superhero's image URL
}
