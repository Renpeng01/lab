package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// go aysncDo()
	// recNilChannel()
	// fmt.Println("sendNilChannel")

	close_channle_send()
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

func close_channle_send() {

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGCONT)
	done := make(chan bool, 1)
	go func() {
		sig := <-sigs
		fmt.Println("signal：", sig)
		done <- true
	}()

	sendChan := make(chan struct{}) // 这里不能使用 chan<- 或者 <-chan，如果使用 <-chan，则不能close(chan), 如果使用chan<-, 则调用close后会panic，提示不能在已经close的chan写入
	for i := 0; i < 10; i++ {
		go func(m int) {
			fmt.Printf("send start %+v\n", m)
			<-sendChan
			fmt.Printf("send end %+v\n", m)
		}(i)
	}
	time.Sleep(5 * time.Second)
	close(sendChan)

	<-done

}

// func close_channle_rec() {
// 	sigs := make(chan os.Signal, 1)
// 	signal.Notify(sigs, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGCONT)
// 	done := make(chan bool, 1)
// 	go func() {
// 		sig := <-sigs
// 		fmt.Println("signal：", sig)
// 		done <- true
// 	}()

// 	recChan := make(<-chan struct{})
// 	for i := 0; i < 10; i++ {
// 		go func(m int) {
// 			fmt.Printf("rec start%+v\n", m)
// 			<-recChan
// 			fmt.Printf("rec end%+v\n", m)
// 		}(i)
// 	}
// 	time.Sleep(5 * time.Second)
// 	close(recChan) // 这里报错  invalid operation: cannot close receive-only channel recChan (variable of type <-chan struct{})

// 	<-done
// }
