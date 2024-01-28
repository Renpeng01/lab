package main

import "fmt"

func main() {
	aa := []int{1}
	slice(aa)
	fmt.Println("aa")
	printSlice(aa)

	fmt.Println("bb")
	bb := make([]int, 0, 20)
	bb = append(bb, 1)
	slice(bb)
	printSlice(bb)

}

// 如果方案是对切片做读写，并想要生效返回，必须有slice返回值，否则在调用处无效
func slice(a []int) {
	for i := 0; i < 10; i++ {
		a = append(a, i)
	}
}

func printSlice(a []int) {
	for i, v := range a {
		fmt.Println("i:", i, " v: ", v)
	}
}
