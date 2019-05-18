package als

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/timkaye11/goRecommend/ALS"
	"github.com/tomjamescn/recommend_cf_als/data"
)

type Model struct {
	Rows int
	Cols int
	Data []float64
}

//计算模型，并生成
//为了可读性，生成的模型是一个json文件
//go中最好使用gob生成二进制序列化数据
func BuildModel() {
	fmt.Println("构建模型开始时间:", time.Now())
	ratingsData, userMaxIndex, movieMaxIndex := data.GetRatings()
	rows := userMaxIndex + 1
	cols := movieMaxIndex + 1
	/*
		rows = 100
		cols = 1000
	*/
	Q := ALS.MakeRatingMatrix(ratingsData, rows, cols)
	n_factors := 5
	n_iterations := 10
	lambda := 0.01

	Qhat, _ := ALS.Train(Q, n_factors, n_iterations, lambda)
	//fmt.Println(Qhat)
	fmt.Println("构建模型结束时间:", time.Now())

	m := &Model{
		Rows: rows,
		Cols: cols,
		Data: Qhat.Array(),
	}
	jsonData, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}
	ioutil.WriteFile("/tmp/model.json", jsonData, 0777)
}
