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

func readRebateList() (rebateList map[int]float32) {
	rebateList = make(map[int]float32)
	rebateList[1] = 0.06 //体验
	rebateList[2] = 0.07 //会员
	rebateList[3] = 0.08 //V1
	rebateList[4] = 0.09 //V2
	rebateList[5] = 0.1  //V3
	rebateList[6] = 0.1  //合伙人
	rebateList[7] = 0.1  //股东
	rebateList[8] = 0.1  //董事
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

	rebateArray := [...]float32{1: 0.06, 2: 0.07, 3: 0.08, 4: 0.09, 5: 0.10, 6: 0.10, 7: 0.10, 8: 0.10}

	for i := range users {
		users[i].rebate = rebateArray[users[i].level]
	}
	fmt.Printf("\n%v", users)
	/* 	var rebateList map[int]float32
	   	rebateList[1] = 0.06f //体验
	   	rebateList[2] = 0.07 //会员
	   	rebateList[3] = 0.08 //V1
	   	rebateList[4] = 0.09 //V2
	   	rebateList[5] = 0.1  //V3
	   	rebateList[6] = 0.1  //合伙人
	   	rebateList[7] = 0.1  //股东
	   	rebateList[8] = 0.1  //董事 */

	var usersLevel []int
	for i := range users {
		usersLevel = append(usersLevel, users[i].level)
		//fmt.Printf("%v", usersLevel)
	}
	var result []int
	result = cutOff(usersLevel)
	fmt.Printf("%v", result)

	var rebateList map[int]float32
	rebateList = readRebateList()
	fmt.Printf("%v", rebateList)

	/* 	r := gin.Default()
	   	r.GET("/ping", func(c *gin.Context) {
	   		c.JSON(200, gin.H{
	   			"message": "pong",
	   		})
	   	})
	   	r.Run() // listen and serve on 0.0.0.0:8080 */

}
