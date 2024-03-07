package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	var wg sync.WaitGroup

	// fmt.Println("wg:", wg)
	// wg.Add(1)
	// wg.Done()
	// wg.Wait()

	// copy := wg
	// fmt.Println("copy", copy)

	task := []int{1, 2, 3}

	for _, v := range task {
		wg.Add(1)
		go func(v int) {
			defer wg.Done()
			fmt.Printf("do %+v\n", v)
			time.Sleep(500 * time.Millisecond)
		}(v)
	}

	// 可以多次判断wait
	wg.Wait()
	wg.Wait()

	fmt.Println("111111")
}
