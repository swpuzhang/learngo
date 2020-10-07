package usemethod

import (
	"bytes"
	"fmt"
)

const (
	// BITSWIDTH 一个字位宽
	bitsWidth = 64
)

// IntSet 比特位集合
type IntSet struct {
	words []int64
}

//UseSet 使用set
func UseSet() {
	var set IntSet
	set.Add(1)
	set.Add(64)
	fmt.Println(set.Has(1))
	fmt.Printf("%s\n", set.String())
}

// Has 是否存在参数Bit
func (set *IntSet) Has(i int) bool {
	word, bit := i/bitsWidth, i%bitsWidth
	return word < len(set.words) && set.words[word]&(1<<bit) != 0
}

// Add 添加比特位
func (set *IntSet) Add(i int) {
	word, bit := i/bitsWidth, i%bitsWidth
	for word >= len(set.words) {
		set.words = append(set.words, 0)
	}
	set.words[word] |= 1 << bit
}

// String 输出
func (set *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteString("{")
	for i, v := range set.words {
		if v == 0 {
			continue
		}
		for j := 0; j < bitsWidth; j++ {
			if v&(1<<j) != 0 {
				buf.WriteByte(' ')
				fmt.Fprintf(&buf, "%d", i*bitsWidth+j)
			}

		}
	}
	buf.WriteByte('}')
	return buf.String()
}
