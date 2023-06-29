package movieutil

import (
	"fmt"
	"strings"
)

const (
	backToTheFuturePrice = 15
	otherMoviePrice      = 20
)

type Movie struct {
	Name  string
	Price float64
}

func ProcessMovies(content string) ([]*Movie, int) {
	rawlines := strings.Split(content, "\n")

	var lines []string

	for _, line := range rawlines {
		trimmedLine := strings.TrimSpace(line)
		if trimmedLine != "" {
			lines = append(lines, trimmedLine)
		}
	}

	movies := make(map[int]*Movie)
	distinctMovies := make(map[string]*Movie)
	encounteredMovies := make(map[string]bool)

	seriesCount := 0
	moviesCounter := 0

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if strings.Contains(line, "Back to the Future") {
			if !encounteredMovies[line] {
				encounteredMovies[line] = true
				distinctMovies[line] = &Movie{Name: line, Price: backToTheFuturePrice}
				movies[moviesCounter] = &Movie{Name: line, Price: backToTheFuturePrice}
				seriesCount++
				moviesCounter++
			} else {
				movies[moviesCounter] = &Movie{Name: line, Price: backToTheFuturePrice}
				moviesCounter++
			}
		} else {
			movies[moviesCounter] = &Movie{Name: line, Price: otherMoviePrice}
			moviesCounter++
		}
		fmt.Println("seriesCount:", seriesCount)
	}

	fmt.Println("Encountered Movies:")
	for movie, encountered := range encounteredMovies {
		fmt.Printf("Movie: %s, Encountered: %v\n", strings.TrimSpace(movie), encountered)
	}

	return convertToSlice(movies), seriesCount
}

// CalculateMoviePrices calculates the final price for each movie, considering the count of distinct series movies,
// and returns the total sum of movie prices after reduction.
func CalculateMoviePrices(movies []*Movie, seriesCount int) float64 {

	fmt.Println("Movies:")
	for _, movie := range movies {
		fmt.Printf("Movie: %s, Price: %.2f\n", movie.Name, movie.Price)
	}

	priceReduction := 0.0
	switch seriesCount {

	case 1:
		priceReduction = 0.0
	case 2:
		priceReduction = 0.10
	case 3:
		priceReduction = 0.20
	default:
		priceReduction = 0.20
	}

	totalPrice := 0.0
	for _, movie := range movies {
		if strings.Contains(movie.Name, "Back to the Future") {
			// fmt.Println(movie.Price)
			// fmt.Println(float64(movie.Price))
			fmt.Println(float64(movie.Price) * (1 - priceReduction))
			reducedPrice := float64(movie.Price) * (1 - priceReduction)
			totalPrice += reducedPrice
			movie.Price = reducedPrice
		} else {
			totalPrice += float64(movie.Price)
		}
	}

	return totalPrice
}

func PrintTotalPrice(movies []*Movie) {
	totalPrice := 0.0

	for _, movie := range movies {
		totalPrice += movie.Price
	}

	fmt.Printf("Total Price: %.2f\n", totalPrice)
}

// PrintMovies prints the movie names and prices
func PrintMovies(movies []*Movie) {
	for _, movie := range movies {
		fmt.Printf("Movie: %s, Price: %f\n", movie.Name, movie.Price)
	}
}

// convertToSlice converts a map of movies to a slice
func convertToSlice(moviesMap map[int]*Movie) []*Movie {
	moviesSlice := make([]*Movie, 0, len(moviesMap))
	for _, movie := range moviesMap {
		moviesSlice = append(moviesSlice, movie)
	}
	return moviesSlice
}
