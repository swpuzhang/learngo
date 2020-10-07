package useinterface

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

// UseTypeAssert 使用类型断言
func UseTypeAssert() {
	// 类型断言可以将接口类型转换成动态类型
	var w io.Writer
	w = os.Stdout
	_ = w.(*os.File)
	//_ = w.(*bytes.Buffer) // panic

	// 接口可以断言为另一个接口. 如果不满足接口值会panic
	_ = w.(io.ReadWriter)
	//w = nil
	//_ = w.(*os.File) // 任何nil值的断言都会panic
	_, ok := w.(*os.File)
	if ok {
		fmt.Println("类型断言成功")
	}
	_, ok = w.(*bytes.Buffer)
	if !ok {
		fmt.Println("类型断言失败")
	}
	errorAssert()
}

func errorAssert() {
	_, err := os.Open("no")
	if err != nil {
		fmt.Println(err)
		if pe, ok := err.(*os.PathError); ok {
			err = pe.Err
		}
	}
	fmt.Println(err)
}

// 判断一个接口是否有某个方法
func writeString(w io.Writer, s string) (n int, err error) {
	type stringWriter interface {
		WriteSstring(string) (int, error)
	}
	if sw, ok := w.(stringWriter); ok {
		return sw.WriteSstring(s)
	}
	return w.Write([]byte(s))
}

// 使用switch
// switch x := x.(type) {}
