package main

import (
	"elmer-almeida/GoMovieCatalog/helpers"
	"elmer-almeida/GoMovieCatalog/utils"
	"fmt"
	"log"
	"os"
)

var MAX_RATING int = 5

func main() {

	helpers.ShowIntro()

	var movies map[string]utils.MovieItem = make(map[string]utils.MovieItem)

	dirPath := "/Volumes/T7 SSD - Data/Media/Movies"

	dirResults, err := os.ReadDir(dirPath)
	if err != nil {
		log.Fatalf("Something went wrong trying to access `%s`. Error: %v\n", dirPath, err)
	}

	// skip .DS_Store
	dirResults = dirResults[1:]

	for _, movieItem := range dirResults {
		movieName := movieItem.Name()
		title, year, err := helpers.GetMovieNameAndYear(movieName)
		rating := helpers.GetRating()
		if err != nil {
			log.Fatalf("Something went wrong with extracting name and year. Error: %s\n", err)
		}

		// Get key in the form of lowercase
		// " " replaced with "-"
		key := helpers.GetKeyForMovie(title)

		var movie utils.MovieItem
		movies[key] = *movie.SetMovie(title, year, rating)
	}

	// Print each item
	helpers.ListMovies(movies)

	helpers.Divider()

	helpers.PrintNumberOfMovies(helpers.NumberOfMovies(movies))

	helpers.Divider()

	helpers.Separator()

	helpers.Divider()

	// Get all movies that are 5 stars
	ratingFilter := MAX_RATING

	for ratingFilter >= 1 {
		filteredMovies := helpers.FilterRatings(movies, ratingFilter)

		fmt.Printf("Showing all %d star movies: \n", ratingFilter)
		for index, movie := range filteredMovies {
			fmt.Printf("%2d) %s (%s) - %s (%d)\n", index+1, movie.Title, movie.Year, helpers.ShowRatingStars(movie.Rating), movie.Rating)
		}
		fmt.Printf("\nThere are %d number of %d star movies.\n", len(filteredMovies), ratingFilter)

		helpers.Divider()
		helpers.SingleSeparator()
		helpers.Divider()
		ratingFilter--
	}

}
