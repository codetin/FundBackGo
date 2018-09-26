package main

import (
	"fmt"
	"math"
)

type userInfo struct {
	uid    int64
	name   string
	level  int
	rebate float64
}

func round(f float64, n int) float64 {
	n10 := math.Pow10(n)
	return math.Trunc((f+0.5/n10)*n10) / n10
}

var rebateArray []float64 = make([]float64, 20)

func readRebateList(rebateArray []float64) {
	rebateArray[1] = 0.06 //体验
	rebateArray[2] = 0.07 //会员
	rebateArray[3] = 0.08 //V1
	rebateArray[4] = 0.09 //V2
	rebateArray[5] = 0.1  //V3
	rebateArray[6] = 0.1  //合伙人
	rebateArray[7] = 0.1  //股东
	rebateArray[8] = 0.1  //董事 */
	return
}

//按照规则裁剪数组
func cutOff(s []int) (cutOffArray []int) {
	var currentMaxLevel int
	var count int //计算平级遇到的次数
	for i, e := range s {
		//第一个元素不进行比较判断
		if i == 0 {
			currentMaxLevel = e
			fmt.Printf("\n首个分红index:%d-value:%d", i, e)
			cutOffArray = append(cutOffArray, i)
			continue
		}
		if e > currentMaxLevel {
			currentMaxLevel = e
			count = 0
			fmt.Printf("\n级差分红index:%d-value:%d", i, e)
			cutOffArray = append(cutOffArray, i)
			continue
		}
		if e == currentMaxLevel {
			if count == 0 {
				fmt.Printf("\n平级分红index:%d-value:%d", i, e)
				cutOffArray = append(cutOffArray, i)
			}
			count++
			continue
		}
	}
	fmt.Printf("\n需要分红数组位置%v", cutOffArray)
	return
}

//按照规则修改返点数据
func userRebate(s []userInfo) (cutOffArray []userInfo) {
	//var currentMaxRebate float64
	var currentMaxLevel int
	var count int //计算平级遇到的次数
	for i, e := range s {
		//首个分红
		if i == 0 {
			//fmt.Printf("\n首个分红index:%d-value:%f", i, e.rebate)
			cutOffArray = append(cutOffArray, e)
			currentMaxLevel = e.level
			continue
		}
		//级差分红
		if e.level > currentMaxLevel {
			count = 0
			//currentMaxRebate = e.rebate
			//fmt.Printf("\n级差分红index:%d-value:%f", i, e.rebate)
			e.rebate = round(e.rebate-rebateArray[currentMaxLevel], 5)
			cutOffArray = append(cutOffArray, e)
			currentMaxLevel = e.level
			continue
		}
		//平级分红
		if e.level == currentMaxLevel {
			if count == 0 {
				//fmt.Printf("\n平级分红index:%d-value:%f", i, e.rebate)
				e.rebate = round(cutOffArray[len(cutOffArray)-1].rebate*0.1, 5)
				cutOffArray = append(cutOffArray, e)
			}
			count++
			continue
		}
	}
	return
}

func main() {

	users := []userInfo{{1, "V1用户", 1, 0.0}, {2, "V2用户", 2, 0.0}, {3, "V2用户", 2, 0.0}, {4, "V2用户", 2, 0.0}, {5, "V3用户", 3, 0.0}, {6, "V3用户", 3, 0.0}, {7, "V1用户", 1, 0.0}, {8, "V4用户", 4, 0.0}, {9, "V5用户", 5, 0.0}}
	//读取最新的返利配置

	readRebateList(rebateArray)
	//设置用户等级对应的返利比例
	for i := range users {
		users[i].rebate = rebateArray[users[i].level]
	}

	var result []userInfo
	result = userRebate(users)
	fmt.Printf("\n%v", users)
	fmt.Printf("\n%v", result)

}
