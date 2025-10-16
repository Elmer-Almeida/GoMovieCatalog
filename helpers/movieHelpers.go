package helpers

import (
	"elmer-almeida/GoMovieCatalog/utils"
	"fmt"
	"math/rand/v2"
	"regexp"
	"strings"
)

func GetMovieNameAndYear(movieName string) (title, year string, err error) {
	re, err := regexp.Compile(`^(.*?)\s*\((\d{4}(?:-\d{4})?)\)$`)
	if err != nil {
		return "", "", fmt.Errorf("Failed to compile regex: %v", err)
	}
	matches := re.FindStringSubmatch(movieName)
	if len(matches) != 3 {
		return "", "", fmt.Errorf("Input format invalid: %s", movieName)
	}

	// Extract name and year
	title = matches[1]
	year = matches[2]

	return title, year, nil
}

func GetKeyForMovie(title string) string {
	return strings.ToLower(strings.ReplaceAll(title, " ", "-"))
}

func GetRating() int {
	return rand.IntN(6-1) + 1
}

func ListMovies(movies map[string]utils.MovieItem) {
	for _, movie := range movies {
		fmt.Printf("%s (%s) - %s (%d)\n",
			movie.Title,
			movie.Year,
			ShowRatingStars(movie.Rating),
			movie.Rating,
		)
	}
}

func NumberOfMovies(movies map[string]utils.MovieItem) int {
	return len(movies)
}

func PrintNumberOfMovies(numMovies int) {
	switch numMovies {
	case 0:
		fmt.Println("There are no movies in your collection.")
	case 1:
		fmt.Printf("There is %d movie in your collection.\n", numMovies)
	default:
		fmt.Printf("There are %d movies in your collection.\n", numMovies)
	}
}

func ShowRatingStars(rating int) string {
	var stars string
	for range rating {
		stars += " "
	}
	return stars
}

func FilterRatings(movies map[string]utils.MovieItem, rating int) (moviesBasedOnRating []utils.MovieItem) {
	for _, movie := range movies {
		if movie.Rating == rating {
			moviesBasedOnRating = append(moviesBasedOnRating, movie)
		}
	}
	return
}
