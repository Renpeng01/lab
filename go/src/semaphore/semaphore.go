package main

import (
	"fmt"
	"log"
	"sync"
	"time"

	"golang.org/x/sync/semaphore"
)

func main() {

	sema := semaphore.NewWeighted(4)
	// ctx := context.Background()
	tasks := []int{1, 2, 3, 4}

	var wg sync.WaitGroup
	for _, v := range tasks {
		if !sema.TryAcquire(1) {
			log.Printf("Failed to acquire semaphore:")
			break
		}
		wg.Add(1)
		go func(v int) {
			fmt.Printf("start: %+v \n", v)
			defer wg.Done()
			defer sema.Release(1)
			time.Sleep(3 * time.Second)
			fmt.Printf("end: %+v \n", v)

		}(v)
	}

	wg.Wait()
	fmt.Println("success")
}
