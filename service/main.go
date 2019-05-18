package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"

	. "github.com/skelterjohn/go.matrix"
	"github.com/timkaye11/goRecommend/ALS"
	"github.com/tomjamescn/recommend_cf_als/als"
	"github.com/tomjamescn/recommend_cf_als/data"
)

var modelMatrix *DenseMatrix
var movieIdToIndexMap map[int]int
var movieIndexToIdMap map[int]int
var movieIdToNameMap map[int]string

func loadModel() {
	var model als.Model

	movieIdToIndexMap, movieIndexToIdMap, movieIdToNameMap, _ = data.GetMovieInfo()

	filePath := fmt.Sprintf("%s/src/github.com/tomjamescn/recommend_cf_als/als/model.json", os.Getenv("GOPATH"))

	fmt.Println(filePath)

	modelData, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(modelData, &model)
	if err != nil {
		panic(err)
	}
	modelMatrix = ALS.MakeRatingMatrix(model.Data, model.Rows, model.Cols)
}

func main() {
	//als.BuildModel()

	loadModel()

	http.HandleFunc("/recommend_cf_asl", func(w http.ResponseWriter, r *http.Request) {
		movieIdStr := r.URL.Query().Get("movie_id")
		movieId, err := strconv.Atoi(movieIdStr)
		if err != nil {
			w.Write([]byte("movie_id参数错误"))
			return
		}

		if movieIndex, ok := movieIdToIndexMap[movieId]; !ok {
			w.Write([]byte("movie_id不是合法值"))
			return
		} else {
			//喜欢这个电影的人
			var likeMovieUserInfo string = "喜欢此电影的人有（排名不分先后）:\n"
			var likeMovieUserIndexList []int
			userRating := modelMatrix.GetColVector(movieIndex).Array()
			//fmt.Println(userRating)
			for userIndex, score := range userRating {
				if score >= 2.5 {
					likeMovieUserIndexList = append(likeMovieUserIndexList, userIndex)
					likeMovieUserInfo += fmt.Sprintf(" %d", userIndex+1)
				}
			}

			w.Write([]byte(likeMovieUserInfo))
			w.Write([]byte("\n\n\n"))

			var otherLikeMovieInfo string = "喜欢此电影的人还喜欢哪些电影（排名不分先后）:\n"
			var otherLikeMovieIdMap = map[int]bool{}
			//喜欢这个电影的人还喜欢哪些电影
			for _, userIndex := range likeMovieUserIndexList {
				movieRating := modelMatrix.GetRowVector(userIndex).Array()
				for movieIndex, score := range movieRating {
					if score > 2.5 {
						otherLikeMovieIdMap[movieIndexToIdMap[movieIndex]] = true
					}
				}
			}

			for movieId, _ := range otherLikeMovieIdMap {
				otherLikeMovieInfo += movieIdToNameMap[movieId] + "\n"
			}

			w.Write([]byte(otherLikeMovieInfo))
		}

	})

	http.ListenAndServe(":9999", nil)
}
