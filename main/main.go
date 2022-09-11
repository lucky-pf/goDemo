package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	array := [80]int{}
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len(array); i++ {
		temp := rand.Intn(100 + 1)
		array[i] = temp
	}
	fmt.Printf("原数组为：\n%v\n排序后为：\n", array)
	// 冒泡排序
	for i := 1; i < len(array); i++ {
		flag := true
		for j := 0; j < len(array)-i; j++ {
			if array[j] > array[j+1] {
				temp := array[j]
				array[j] = array[j+1]
				array[j+1] = temp
				flag = false
			}
		}
		if flag {
			break
		}
	}
	fmt.Println(array)
}
