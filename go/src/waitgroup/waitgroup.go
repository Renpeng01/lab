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

	go func() {
		wg.Wait()
		fmt.Println("wait11111")
	}()

	// 不可以多次wait
	go func() {
		wg.Wait()
		fmt.Println("wait222")
	}()

	s := make(chan struct{})
	<-s
}
