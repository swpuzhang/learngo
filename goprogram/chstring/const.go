package chstring

import "fmt"

// TestStringConst 测试字符串不变性
func TestStringConst() {
	s1 := "1234"
	s2 := s1
	s1 += ", 5678"
	//s1[0] = "0"  //不能改变
	fmt.Println("s1:" + s1)
	fmt.Println("s2:" + s2)
	s3 := s1[6:]
	fmt.Println("s3:" + s3)
}
