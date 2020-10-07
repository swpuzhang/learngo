package usegorout

import (
	"fmt"
	"time"
)

// UseGo 使用go协程
func UseGo() {

	ret := fib(40)
	fmt.Printf("\r%d", ret)

	go EchoServer()
	for i := 0; i < 1; i++ {
		go Netcat("127.0.0.1:8080")
	}
	func() {
		for {

			time.Sleep(100 * time.Millisecond)
		}
	}()
}

func fib(n int) int {
	if n < 2 {
		return n
	}
	return fib(n-1) + fib(n-2)
}
