package main

type Game struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Difficulty  int    `json:"difficulty"`
	MinPlayer   int    `json:"min-player"`
	MaxPlayer   int    `json:"max-player"`
}
