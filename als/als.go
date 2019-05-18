package als

import (
	"fmt"

	"github.com/timkaye11/goRecommend/ALS"
	"github.com/tomjamescn/recommend_cf_als/data"
)

//计算模型，并生成
//为了可读性，生成的模型是一个json文件
//go中最好使用gob生成二进制序列化数据
func als() {
	ratingsData, userMaxIndex, movieMaxIndex := data.GetRatings()
	fmt.Println(len(ratingsData))
	Q := ALS.MakeRatingMatrix(ratingsData, userMaxIndex+1, movieMaxIndex+1)
	n_factors := 5
	n_iterations := 10
	lambda := 0.01

	Qhat, _ := ALS.Train(Q, n_factors, n_iterations, lambda)
	fmt.Println(Qhat)

}
