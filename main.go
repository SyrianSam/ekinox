package main

import (
	"fmt"
	"log"

	"ekinox/fileutil"

	"github.com/gin-gonic/gin"

	"ekinox/movieutil"
)

func main() {

	// Initialize the gin router
	router := gin.Default()

	// Define API endpoints and their handlers
	router.GET("/movies", getMoviesHandler)
	router.POST("/movies", processMoviesHandler)

	// Run the server
	router.Run(":8080")

	// Read the file
	filePath := "panier.txt" // Replace with the actual file path
	content, err := fileutil.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Error reading file: %s", err)
	}

	// Process the file
	movies, seriesCount := movieutil.ProcessMovies(content)

	// Calculate the final price for each movie
	fmt.Println(movieutil.CalculateMoviePrices(movies, seriesCount))
	// movieutil.PrintTotalPrice(movies)
	// Print the result
	// movieutil.PrintMovies(movies)
}
