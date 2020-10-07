package usefunc

import (
	"fmt"
	"os"
)

// UseFunc 使用func
func UseFunc() {
	var f func(int, int) int
	f = add
	fmt.Printf("%T , 1, 2 : %d", f, f(1, 2))
	f1 := squares()
	fmt.Println(f1())

	fmt.Println(f1())
	fmt.Println(f1())
	topoSort(prereqs)
	fmt.Println("---------------")
	guangduSort(prereqs)
	rmDir()
}

func add(i, j int) int {
	return i + j
}
func squares() func() int {
	x := 0
	return func() int {
		x++
		return x * x
	}
}

var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},
	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},
	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

func topoSort(m map[string][]string) {
	seen := make(map[string]bool)
	var visit func([]string)
	visit = func(values []string) {
		for _, v := range values {
			if !seen[v] {
				seen[v] = true
				visit(m[v])
				fmt.Println(v)
			}
		}
	}
	var keys []string
	for k := range m {
		keys = append(keys, k)
	}
	visit(keys)
}

func guangduSort(m map[string][]string) {
	var keys []string
	for k := range m {
		keys = append(keys, k)
	}
	seen := map[string]bool{}
	for len(keys) > 0 {
		top := keys[0]
		keys = keys[1:]
		if seen[top] {
			continue
		}
		if pres, ok := m[top]; ok {
			keys = append(keys, pres...)
		}
		seen[top] = true
		fmt.Println(top)
	}
}

func rmDir() {
	tempDir := []string{"1", "2"}
	var dirs []func()
	for _, v := range tempDir {
		dir := v //这里是必须的, 闭包会引用变量
		os.MkdirAll(v, 0755)
		dirs = append(dirs, func() {
			os.RemoveAll(dir)
		})
	}
	for _, dir := range dirs {
		dir()
	}
}
