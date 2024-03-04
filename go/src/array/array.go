package main

import "fmt"

func main() {
	var arr [10]int // 声明数组  arr3 := [...]int{1, 2, 3}

	for i, v := range arr {
		fmt.Println("i: ", i, " v: ", v)
		// v = i
		// arr[i] = i
	}

	for i, v := range arr {
		fmt.Println("i: ", i, " v: ", v)
	}

	arr = append(arr, 1)

	var slice []int // 声明的是切片
	slice = append(slice, 1)
}
