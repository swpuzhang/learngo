package chstring

import (
	"fmt"
	"unicode/utf8"
)

func testUtf8() {
	s1 := "世界"
	fmt.Println(s1)
	s2 := '世'
	fmt.Println(s2) //19990 println打印码点
	fmt.Printf("%c\n", s2)
	fmt.Println('\u4e16')
	//fmt.Println('\xe4\xb8\x96') //单个字符不能使用多个十六进制表示
	s3 := "\xe4\xb8\x96\xe7\x95\x8c"
	fmt.Println(s3)
	fmt.Println("\u4e16\u754c")
	fmt.Println("\U00004e16\U0000754c") //32位码点需要使用\U大写的U
}

// HasPrefix 判断前缀
func HasPrefix(s, prefix string) bool {
	return len(s) >= len(prefix) && s[:len(prefix)] == prefix
}

// Useutf8 使用utf8
func Useutf8() {
	testUtf8()
	s := "hello, 世界"
	fmt.Println(len(s))
	fmt.Println(utf8.RuneCountInString(s))
	for i := 0; i < len(s); {
		c, size := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%c\t%d\n", c, size)
		i += size
	}

	// range 会自动解码utf8
	for i, c := range s {
		fmt.Printf("%q\t%d\t%d\n", c, c, i)
	}

	n := 0
	for range s {
		n++
	}
	fmt.Printf("字符数:%d\n", n)
	runeString()
}

func runeString() {
	s := "你好,世界"
	fmt.Printf("% x\n", s)
	r := []rune(s) //string 转rune
	fmt.Printf("% x\n", r)

	//rune转string
	s2 := string(r)
	fmt.Println(s2)
	//将一个整数调用string
	s3 := string(65)
	fmt.Println(s3)
	s4 := string(0x4eac)
	fmt.Println(s4)
}
