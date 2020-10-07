package usefunc

import (
	"fmt"
	"runtime"
)

// UsePanic 使用panic
func UsePanic() {
	var buf [4098]byte
	defer func() {
		n := runtime.Stack(buf[:], false)
		fmt.Printf("%s", buf[:n])
	}()
	defer func() {
		if p := recover(); p != nil {
			fmt.Println(p)
		}
	}()
	f(3)
}

// defer后注册先执行
func f(x int) {
	fmt.Printf("f(%d)\n", x+0/x)
	defer func() {
		fmt.Printf("%d\n", x)
	}()
	f(x - 1)
}
