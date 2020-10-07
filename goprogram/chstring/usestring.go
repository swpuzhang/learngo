package chstring

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

func basename(s string) string {
	index := strings.LastIndex(s, "/")
	s = s[index+1:]
	if index = strings.LastIndex(s, "."); index >= 0 {
		s = s[:index]
	}
	return s
}

// UseString 使用string
func UseString() {
	s1 := "/a/b/c.go"
	fmt.Printf("%s basename is %s\n", s1, basename(s1))
	fmt.Println(useBuffer())
	useConvert()
}

func useBuffer() string {
	var buff bytes.Buffer
	buff.WriteByte('[')
	num := []int{1, 2, 3}
	for i, v := range num {
		if i > 0 {
			buff.WriteString(", ")
		}
		fmt.Fprintf(&buff, "%d", v)
	}
	buff.WriteByte(']')
	return buff.String()
}

func useConvert() {
	num := 123
	strNum := strconv.Itoa(num)
	strNum2 := fmt.Sprintf("%d", num)
	strNum3 := strconv.FormatInt(int64(num), 10)
	num1, _ := strconv.Atoi(strNum)
	num2, _ := strconv.ParseInt(strNum2, 10, 64)
	fmt.Println(strNum3, num1, num2) //Println默认使用空格
}
