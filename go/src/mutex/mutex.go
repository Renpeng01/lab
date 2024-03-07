package main

import "sync"

func main() {
	var mutex sync.RWMutex
	mutex.RLock()
	defer mutex.RUnlock()

	mutex.Lock()
	defer mutex.Unlock()

}
