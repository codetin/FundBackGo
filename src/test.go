package main

import (
	"fmt"
	"time"
)

type userInfo struct {
	uid    int64
	name   string
	level  int
	rebate float32
}

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s, i)

	}
}

func readRebateList() (rebateArray []float32) {
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

func main() {
	/* 	go say("world")
	   	go say("hello test")
	   	time.Sleep(500 * time.Millisecond)
	   	//time.Until(time.Now())
	   	println("main")
	   	fmt.Printf("%s", "test")
	   	//say("ok") */

	//array := []int{1, 2, 3, 2, 1, 3}

	/* 	array := []int{1, 1, 1, 2, 2, 2, 3, 1, 2, 3, 3}
	   	fmt.Printf("%v\n", array)
	   	cutOff(array) */

	users := []userInfo{{1, "V1用户", 1, 0.0}, {2, "V2用户", 2, 0.0}}

	rebateArray := readRebateList()
	//设置用户等级对应的返利比例
	for i := range users {
		users[i].rebate = rebateArray[users[i].level]
	}
	fmt.Printf("\n%v", users)

	var usersLevel []int
	for i := range users {
		usersLevel = append(usersLevel, users[i].level)
		//fmt.Printf("%v", usersLevel)
	}
	var result []int
	result = cutOff(usersLevel)
	fmt.Printf("%v", result)

	/* 	r := gin.Default()
	   	r.GET("/ping", func(c *gin.Context) {
	   		c.JSON(200, gin.H{
	   			"message": "pong",
	   		})
	   	})
	   	r.Run() // listen and serve on 0.0.0.0:8080 */

}
