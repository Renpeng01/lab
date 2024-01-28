package main

import "fmt"

func main() {
	go aysncDo()
	recNilChannel()
	fmt.Println("sendNilChannel")
}

// 对一个nil的channel读写会阻塞
func sendNilChannel() {
	var a chan int
	a <- 1
}

func recNilChannel() {
	var a chan int
	<-a
}

func initChannel() {
	a := make(chan int, 0)
	a <- 1
}

func aysncDo() {
	for {

	}
}
