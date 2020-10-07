package array

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
	"unicode/utf8"
)

// UseSlice 使用slice
func UseSlice() {
	mouths := []string{
		"", "Jan", "Feb", "Mar", "Apr", "May", "June", "July",
		"Aug", "Sep", "Oct", "Nov", "Dec",
	}
	q1 := mouths[1:4]
	summer := mouths[6:9] //
	// invalid := mouths[:20]
	// fmt.Println(invalid)
	extendSlice := summer[:4]
	fmt.Printf("%v\t%d\t%d\n", q1, len(q1), cap(q1))
	fmt.Printf("%v\t%d\t%d\n", summer, len(summer), cap(summer))
	fmt.Printf("%v\t%d\t%d\n", extendSlice, len(extendSlice), cap(extendSlice))
	var arr []int
	var p *int
	for i := 0; i < 10; i++ {

		arr = appendInt(arr, i)
		p = &arr[0]
		fmt.Printf("%v, %d, %d, %p\n", arr, len(arr), cap(arr), p)
	}
	statisticUtf8()
}

func appendInt(x []int, y ...int) []int {
	curLen := len(x)
	afterLen := curLen + len(y)
	curCap := cap(x)
	var z []int
	if afterLen <= curCap {
		z = x[:afterLen] // 扩展长度
	} else {
		growCap := 2 * curLen
		if afterLen > growCap {
			growCap = afterLen
		}
		z = make([]int, afterLen, growCap)
		copy(z, x)
	}
	copy(z[curLen:], y)
	// z[curLen] = y
	return z
}

func remove(x []int, i int) []int {
	copy(x[i:], x[i+1:])
	return x[:len(x)-1]
}

func statisticUtf8() {
	counter := make(map[rune]int)
	var utf8Len [utf8.UTFMax + 1]int
	invalid := 0
	reader := bufio.NewReader(os.Stdin)
	for {
		r, size, err := reader.ReadRune()
		if err != nil {
			fmt.Println(err)
			break
		}
		if r == unicode.ReplacementChar && size == 1 {
			invalid++
		}
		counter[r]++
		utf8Len[size]++
	}
	for k, v := range counter {
		fmt.Printf("%q\t%d\n", k, v)
	}
	for i, v := range utf8Len {
		fmt.Printf("%d\t%d\n", i, v)
	}
	fmt.Printf("invalid %d\n", invalid)
}
