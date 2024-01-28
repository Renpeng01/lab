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
