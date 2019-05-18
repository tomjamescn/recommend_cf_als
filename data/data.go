package data

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
)

//封装读取csv的功能
func readCSV(filename string, f func([]string)) {
	filePath := fmt.Sprintf("%s/src/github.com/tomjamescn/recommend_cf_als/data/%s.csv", os.Getenv("GOPATH"), filename)
	fmt.Println(filePath)
	csvFile, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer csvFile.Close()

	r := csv.NewReader(csvFile)

	isHead := true
	for {
		record, err := r.Read()
		if err == io.EOF {
			return
		}
		if err != nil {
			panic(err)
		}
		if isHead {
			isHead = false
			continue
		}
		f(record)
	}
}

func GetRatings() (ratingsData []float64, userMaxIndex, movieMaxIndex int) {
	ratingsData = make([]float64, 0)
	minUserId := 9999999
	maxUserId := 0

	movieIdToIndexMap, _, _, movieMaxIndex := GetMovieInfo()

	lastUserId := 0
	var userRating []float64

	readCSV("ratings", func(record []string) {
		userIdStr := record[0]
		if userIdStr == "" {
			return
		}
		movieIdStr := record[1]
		ratingStr := record[2]

		userId, err := strconv.Atoi(userIdStr)
		if err != nil {
			panic(err)
		}
		movieId, err := strconv.Atoi(movieIdStr)
		if err != nil {
			panic(err)
		}
		rating, err := strconv.ParseFloat(ratingStr, 64)
		if err != nil {
			panic(err)
		}

		if userId > lastUserId {
			if lastUserId != 0 {
				//上一个用户的所有rating信息已经得到了，加入到最终结果中
				for _, rating := range userRating {
					ratingsData = append(ratingsData, rating)
				}
			}
			//是新的一个用户的信息
			userRating = make([]float64, movieMaxIndex+1)
		}

		//更新userRating
		userRating[movieIdToIndexMap[movieId]] = rating

		if minUserId > userId {
			minUserId = userId
		}
		if maxUserId < userId {
			maxUserId = userId
		}

		lastUserId = userId
	})

	userMaxIndex = maxUserId - 1

	for _, rating := range userRating {
		ratingsData = append(ratingsData, rating)
	}
	return
}

func GetMovieInfo() (idToIndexMap map[int]int, indexToIdMap map[int]int, idToNameMap map[int]string, movieMaxIndex int) {
	movieMaxIndex = 0
	idToIndexMap = make(map[int]int, 0)
	indexToIdMap = make(map[int]int, 0)
	idToNameMap = make(map[int]string, 0)
	index := 0
	readCSV("movies", func(record []string) {
		movieIdStr := record[0]
		if movieIdStr == "" {
			return
		}
		movieName := record[1]
		movieId, err := strconv.Atoi(movieIdStr)
		if err != nil {
			panic(err)
		}
		if index > movieMaxIndex {
			movieMaxIndex = index
		}
		idToIndexMap[movieId] = index
		indexToIdMap[index] = movieId
		idToNameMap[movieId] = movieName
		index++
	})

	return
}
