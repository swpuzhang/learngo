package ch1

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func Fetch(url string) {
	start := time.Now()
	res, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	data, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(len(data) != 0)
	secs := time.Since(start).Seconds()
	fmt.Printf("%.2fs %7d %s", secs, len(data), url)
}

// ExampleFetch 必须在最后加//OutPut:
func ExampleFetch() {
	urls := []string{"http://www.baidu.com"}
	for _, url := range urls {
		Fetch(url)
	}
	// Output: true
}
