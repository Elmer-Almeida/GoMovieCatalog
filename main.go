package main

import (
	"elmer-almeida/GoMovieCatalog/helpers"
	"elmer-almeida/GoMovieCatalog/utils"
	"fmt"
	"log"
	"os"
)

var MAX_RATING int = 5
var URL string = "http://www.omdbapi.com/"
var movies map[string]utils.MovieItem

func main() {

	// http://www.omdbapi.com/?apikey=<API_KEY>
	os.Setenv("API_KEY", "ec582a81")

	helpers.ShowIntro()

	movies = make(map[string]utils.MovieItem)

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
		rating := helpers.GetRating() // generate rating between 1 & 5
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

	// Show all movies filtered by rating
	helpers.ShowFilteredRatings(MAX_RATING, movies)

	jsonData := `{ "Title": "Weapons", "Year": "2025", "Rated": "R", "Released": "08 Aug 2025", "Runtime": "128 min", "Genre": "Horror, Mystery", "Director": "Zach Cregger", "Writer": "Zach Cregger", "Actors": "Julia Garner, Josh Brolin, Alden Ehrenreich", "Plot": "When all but one child from the same class mysteriously vanish on the same night at exactly the same time, a community is left questioning who or what is behind their disappearance.", "Language": "English", "Country": "United States", "Awards": "1 nomination total", "Poster": "https://m.media-amazon.com/images/M/MV5BNTBhNWJjZWItYzY3NS00M2NkLThmOWYtYTlmNzBmN2UxZWFjXkEyXkFqcGc@._V1_SX300.jpg", "Ratings": [ { "Source": "Internet Movie Database", "Value": "7.6/10" }, { "Source": "Rotten Tomatoes", "Value": "94%" }, { "Source": "Metacritic", "Value": "81/100" } ], "Metascore": "81", "imdbRating": "7.6", "imdbVotes": "159,990", "imdbID": "tt26581740", "Type": "movie", "DVD": "N/A", "BoxOffice": "$150,200,554", "Production": "N/A", "Website": "N/A", "Response": "True" }`

	// http://www.omdbapi.com/?apikey=ec582a81&t=weapons&y=2025
	testItem := movies["weapons"]
	title, year := testItem.GetMovieShortDetails()
	UrlString := fmt.Sprintf("%s?apikey=%s&t=%s&y=%s", URL, os.Getenv("API_KEY"), title, year)
	fmt.Println("The URL is: ", UrlString)
	// response, err := http.Get(UrlString)
	// if err != nil {
	// 	log.Fatalf("Unable to get a response from URL. Error: %v\n", err)
	// }
	// data, err := io.ReadAll(response.Body)
	// if err != nil {
	// 	log.Fatalf("Unable to get request body. Error: %v\n", err)
	// }
	fmt.Println("Data: ", string(jsonData))
}
