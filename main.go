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

	// get port dynamically for railway
	port := os.Getenv("PORT")

	//local configuration
	if port == "" {
		port = "8080"
	}

	router.Run(":" + port)
}

func homePageHandler(c *gin.Context) {
	// Render the home page template with the input form
	c.HTML(http.StatusOK, "index.html", nil)
}

func calculateHandler(c *gin.Context) {
	// Get the list of movies from the form submission
	movies := c.PostForm("userinput")

	// Process the movies
	movieList, seriesCount := movieutil.ProcessMovies(movies)

	// Calculate the total price
	totalPrice := movieutil.CalculateMoviePrices(movieList, seriesCount)

	// Return the result to the user
	c.JSON(http.StatusOK, gin.H{
		"totalPrice": totalPrice,
	})
}
