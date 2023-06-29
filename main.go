package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	// "ekinox/fileutil"

	"ekinox/movieutil"
)

func main() {

	// Initialize the Gin router
	router := gin.Default()

	// Define the route for the home page
	router.GET("/", homePageHandler)

	// Define the route for the form submission
	router.POST("/calculate", calculateHandler)

	// Run the server
	router.Run(":8080")

	// // Read the file
	// filePath := "panier.txt" // Replace with the actual file path
	// content, err := fileutil.ReadFile(filePath)
	// if err != nil {
	// 	log.Fatalf("Error reading file: %s", err)
	// }

	// // Process the file
	// movies, seriesCount := movieutil.ProcessMovies(content)

	// // Calculate the final price for each movie
	// fmt.Println(movieutil.CalculateMoviePrices(movies, seriesCount))
	// // movieutil.PrintTotalPrice(movies)
	// // Print the result
	// // movieutil.PrintMovies(movies)
}

func homePageHandler(c *gin.Context) {
	// Render the home page template with the input form
	c.HTML(http.StatusOK, "index.html", nil)
}

func calculateHandler(c *gin.Context) {
	// Get the list of movies from the form submission
	movies := c.PostForm("movies")

	// Process the movies and calculate the price
	movieList, seriesCount := movieutil.ProcessMovies(movies)
	movieutil.CalculateMoviePrices(movieList, seriesCount)

	// Calculate the total price
	totalPrice := 0.0
	for _, movie := range movieList {
		totalPrice += movie.Price
	}

	// Return the result to the user
	c.JSON(http.StatusOK, gin.H{
		"totalPrice": totalPrice,
	})
}
