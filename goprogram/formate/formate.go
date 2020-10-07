package formate

import (
	"fmt"
)

type user struct {
	ID   int
	name string
}

// Formate example
func Formate() {
	u := user{ID: 1, name: "zhangyang"}
	fmt.Printf("%v\n", u)
	fmt.Printf("%T\n", u)
}
