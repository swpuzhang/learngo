package ch1

import (
	"bufio"
	"fmt"
	"os"
	"testing"
)

//Dump1 example Dump1
func TestDump1(t *testing.T) {
	counter := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		counter[input.Text()]++
	}
	for line, n := range counter {
		fmt.Printf("%d\t%s\n", n, line)
	}
}
