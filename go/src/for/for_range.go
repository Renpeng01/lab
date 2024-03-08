package main

import "fmt"

func main() {

	arr := []int{1, 2, 3}
	i := 1

	// 循环只遍历了原始切片中的三个元素，在遍历切片时追加的元素壁不会增加循环的执行次数
	for _, v := range arr {
		arr = append(arr, v)
		fmt.Printf("do %+v\n", i)
		i++
	}
	fmt.Println("done and exit")
	fmt.Println(arr)

	for range arr {

	}
}
