package ch3

import "fmt"

func UseString() {
	s := "hello中国"
	fmt.Printf("len:%[1]d\n", len(s)) //len 返回字节数
	fmt.Printf("at:%c", s[5])         //下标返回某个字节
}
