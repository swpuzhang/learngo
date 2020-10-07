package usefunc

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

// UseDefer 使用defer
func UseDefer() {
	bigSlowOperation()
	doubleValue(1)
	tripleValue(1)
}

func httpGet() {
	const url = "www.baidu.com"
	resp, err := http.Get("url")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	ct := resp.Header.Get("Content-Type")
	if ct != "text/html" && !strings.HasPrefix(ct, "text/html;") {
		fmt.Printf("%s has type %s, not text/html\n", url, ct)
		return
	}
}

func readFile(filename string) ([]byte, error) {
	f, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer f.Close()
	return ioutil.ReadAll(f)
}

func bigSlowOperation() {
	defer trace("bigSlowOperation")() // don't forget the extra parentheses
	// ...lots of work…
	time.Sleep(2 * time.Second) // simulate slow operation by sleeping
}
func trace(msg string) func() {
	start := time.Now()
	log.Printf("enter %s", msg)
	return func() {
		log.Printf("exit %s (%s)", msg, time.Since(start))
	}
}

func doubleValue(x int) (result int) {
	defer func() {
		fmt.Printf("doublevalue x:%d is %d\n", x, result)
	}()
	return x + x
}

// 还可以修改返回值
func tripleValue(x int) (result int) {
	defer func() {
		result += x
		fmt.Printf("triplevalue x:%d is %d\n", x, result)
		//return result + x //  defer 不能返回值
	}()
	return x + x
}
