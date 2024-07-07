package model

type Pet struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Species string `json:"species"`
	Age     int    `json:"age"`
	Adopted bool   `json:"adopted"`
}
