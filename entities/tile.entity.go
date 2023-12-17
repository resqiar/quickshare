package entities

import "time"

type Tile struct {
	ID          int
	Name        string
	Description string
	Color       string
	Width       float32
	Length      float32
	Height      float32
	Weight      float32
	Price       int
	Quantity    int

	CreatedAt time.Time
	UpdatedAt time.Time
}
