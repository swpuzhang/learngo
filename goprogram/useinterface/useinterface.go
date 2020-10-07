package useinterface

import (
	"flag"
	"fmt"
	"net/http"
	"sort"
	"time"
)

var period = flag.Duration("period", 1*time.Second, "sleep period")

// UseInterface 使用接口
func UseInterface() {
	useFlag()
	useSort()
	useHTTP()
}
func useFlag() {
	flag.Parse()
	time.Sleep(*period)
	fmt.Printf("sleep %d\n", *period)
}

func useSort() {
	values := []int{3, 1, 2, 4}
	fmt.Println(sort.IntsAreSorted(values))
	sort.Ints(values)
	fmt.Println(values)
	fmt.Println(sort.IntsAreSorted(values))
	sort.Sort((sort.Reverse(sort.IntSlice(values))))
	fmt.Println(values)
}

type database map[string]float64

func (db *database) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/list":
		for k, v := range *db {
			fmt.Fprintf(w, "%s, %.2f\n", k, v)
		}
	case "/price":
		item := req.URL.Query().Get("item")
		price, ok := (*db)[item]
		if !ok {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		fmt.Fprintf(w, "%.2f", price)
	default:
		w.WriteHeader(http.StatusNotFound)
	}

}

func useHTTP() {
	db := &database{"yifu": 1.1, "wazi": 0.1}
	http.ListenAndServe("127.0.0.1:8080", db)
}
