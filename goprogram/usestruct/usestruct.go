package usestruct

import (
	"fmt"
)

type employee struct {
	ID   int
	name string
}

func employeeByID(id int) *employee {
	return &employee{
		ID:   id,
		name: "test",
	}
}
func employeeByID1(id int) employee {
	return employee{
		ID:   id,
		name: "test",
	}
}
func useEmployee() {
	employeeByID(1).name = "test1"
	// employeeByID1(1).name = "test1" // 返回值类型不能给里面的字段赋值
	fmt.Println(employeeByID1(1).name) // 但是可以使用字段
}

type tree struct {
	value       int
	left, right *tree
}

func add(root *tree, value int) *tree {
	if root == nil {
		root = new(tree)
		root.value = value
		return root
	}
	if value < root.value {
		root.left = add(root.left, value)
	} else {
		root.right = add(root.right, value)
	}
	return root
}

func sort(values []int) {
	var root *tree
	for _, v := range values {
		add(root, v)
	}

}

func appendValue(values []int, root *tree) []int {
	if root != nil {
		appendValue(values, root.left)
		values = append(values, root.value)
		appendValue(values, root.right)
	}
	return values
}
