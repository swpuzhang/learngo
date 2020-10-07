package ch1

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"testing"
)

func TestDump2(t *testing.T) {
	counter := map[string]int{}
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counter)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "%s", err)
				continue
			}
			countLines(f, counter)
			f.Close()
		}
	}
	for line, n := range counter {
		fmt.Printf("%d\t%s\n", n, line)
	}
}

func countLines(r io.Reader, counter map[string]int) {
	input := bufio.NewScanner(r)
	for input.Scan() {
		counter[input.Text()]++
	}
}
