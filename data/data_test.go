package data

import (
	"testing"
)

func TestReadCSV(t *testing.T) {
	readCSV("ratings", func(record []string) {
		//fmt.Println(record)
	})
}

func TestGetRatings(t *testing.T) {

	movieIdToIndexMap, _, _, _ := GetMovieInfo()
	ratingsData, userMaxIndex, movieMaxIndex := GetRatings()
	if userMaxIndex != 609 {
		t.Error("userMaxIndex error!")
		return
	}
	if movieMaxIndex != 9741 {
		t.Error("movieMaxIndex error!")
		return
	}

	if ratingsData[0] != 4 {
		t.Error("ratingsData userIndex:0 movieId:1 error!")
		return
	}

	if ratingsData[movieIdToIndexMap[3147]] != 5 {
		t.Error("ratingsData userIndex:0 movieId:3147 error!")
		return
	}

	if ratingsData[(movieMaxIndex+1)*(4-1)+movieIdToIndexMap[2204]] != 5 {
		t.Error("ratingsData userIndex:3 movieId:2204 error!")
		return
	}

	if ratingsData[(movieMaxIndex+1)*(89-1)+movieIdToIndexMap[88785]] != 2 {
		t.Error("ratingsData userIndex:88 movieId:88785 error!")
		return
	}
}
