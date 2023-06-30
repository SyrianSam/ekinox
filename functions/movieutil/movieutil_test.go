package movieutil

import (
	"fmt"
	"reflect"
	"testing"
)

func TestProcessMovies(t *testing.T) {
	// Test case 1
	content := "Back to the Future 1\nOther Movie 1\n"
	expectedMovies := []*Movie{
		{Name: "Back to the Future 1", Price: backToTheFuturePrice},
		{Name: "Other Movie 1", Price: otherMoviePrice},
	}
	expectedSeriesCount := 1

	movies, seriesCount := ProcessMovies(content)
	fmt.Println(movies)

	if !reflect.DeepEqual(movies, expectedMovies) {
		t.Errorf("Unexpected movies. Expected: %v, Got: %v", expectedMovies, movies)
	}

	if seriesCount != expectedSeriesCount {
		t.Errorf("Unexpected seriesCount. Expected: %d, Got: %d", expectedSeriesCount, seriesCount)
	}

	// Test case 2
	content = "Other Movie 1\nOther Movie 2\nOther Movie 3\n"
	expectedMovies = []*Movie{
		{Name: "Other Movie 1", Price: otherMoviePrice},
		{Name: "Other Movie 2", Price: otherMoviePrice},
		{Name: "Other Movie 3", Price: otherMoviePrice},
	}
	expectedSeriesCount = 0

	movies, seriesCount = ProcessMovies(content)

	if !reflect.DeepEqual(movies, expectedMovies) {
		t.Errorf("Unexpected movies. Expected: %v, Got: %v", expectedMovies, movies)
	}

	if seriesCount != expectedSeriesCount {
		t.Errorf("Unexpected seriesCount. Expected: %d, Got: %d", expectedSeriesCount, seriesCount)
	}
}
