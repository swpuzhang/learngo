package ch1

import (
	"fmt"
	"io/ioutil"
	"strings"
	"testing"
)

func TestDump3(t *testing.T) {
	counter := map[string]int{}
	files := []string{"echo3.go"}
	for _, filename := range files {
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Println(err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			counter[line]++
		}
	}
	for line, n := range counter {
		fmt.Printf("%d\t%s\n", n, line)
	}
}
