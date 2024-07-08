package model

type Pet struct {
	ID    uint   `gorm:"primary_key"`
	Name  string `json:"name"`
	Breed string `json:"breed"`
	Age   int    `json:"age"`
	Type  string `json:"type"`
}
