package utils

import (
	"fmt"
)

type MovieItem struct {
	Title  string
	Year   string
	Rating int
}

func (m *MovieItem) GetMovieItem() *MovieItem {
	return m
}

func (m *MovieItem) GetMovieShortDetails() (string, string) {
	return m.Title, m.Year
}

func (m *MovieItem) GetMovieFullDetails() (string, string, int) {
	return m.Title, m.Year, m.Rating
}

func (m *MovieItem) SetMovie(title, year string, rating int) *MovieItem {
	m.Title = title
	m.Year = year
	m.Rating = rating
	return m
}

func (m *MovieItem) SetRating(rating int) {
	m.Rating = rating
}

func (m *MovieItem) PrintMovieItem() {
	fmt.Printf("%s (%s) - %d\n", m.Title, m.Year, m.Rating)
}
