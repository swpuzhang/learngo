package ch1

import (
	"fmt"
	"os"
	"strings"
)

// Echo3 example Echo3
func Echo3() {
	fmt.Println(strings.Join(os.Args[1:], ","))
}
