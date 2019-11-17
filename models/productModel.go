package models

type Product struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	FilmDirector string `json:"filmDirector"`
	Year         string `json:"year"`
}
