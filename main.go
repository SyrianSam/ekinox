package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"os"

	"ekinox/functions/movieutil"
)

func main() {

	// Initialize the Gin router
	router := gin.Default()

	router.Static("/favicon.ico", "./favicon.ico")
	router.Static("/styles", "./styles")
	router.Static("/images", "./images")
	router.Static("/javascript", "./javascript")
	router.LoadHTMLGlob("templates/*")
	// Define the route for the home page
	router.GET("/", homePageHandler)

	// Define the route for the form submission
	router.POST("/calculate", calculateHandler)

	port := os.Getenv("PORT")
	// Run the server

	if port == "" {
		port = "8080"
	}

	router.Run(":" + port)
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
	movies := c.PostForm("userinput")

	// Process the movies and calculate the price
	movieList, seriesCount := movieutil.ProcessMovies(movies)
	// movieutil.CalculateMoviePrices(movieList, seriesCount)

	// Calculate the total price
	totalPrice := movieutil.CalculateMoviePrices(movieList, seriesCount)
	// for _, movie := range movieList {
	// 	totalPrice += movie.Price
	// }

	// Return the result to the user
	c.JSON(http.StatusOK, gin.H{
		"totalPrice": totalPrice,
	})
}
