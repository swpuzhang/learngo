package array

import "fmt"

// UseArray  使用数组
func UseArray() {
	var arr [3]int
	for i := range arr {
		arr[i] = i
	}
	fmt.Printf("%v\n", arr)
	fmt.Printf("%T\n", arr)
	var arr2 [3]int = [3]int{0, 1} //指定元素不够长度,默认为0
	fmt.Println(arr2)
	arr3 := [...]int{0, 1, 2}
	fmt.Println(arr3)
	num := 10
	arr4 := [...]int{0: 0, 10: num}
	fmt.Println(arr4)
	arr5 := arr3[0:2] //返回一个切片, 共享一个底层数组
	fmt.Printf("%+v, %d, %d\n", arr5, len(arr5), cap(arr5))
}
