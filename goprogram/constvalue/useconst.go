package constvalue

import "fmt"

const pi = 3.14

const (
	e   = 2.71828182845904523536028747135266249775724709369995957496696763
	pi2 = 3.14159265358979323846264338327950288419716939937510582097494459
)

const (
	flag1 = 1 << iota //1 << 0
	flag2             //1 << 1
	flag3             //1 << 2
	flag4             //1 << 3
)

const (
	_ = 1 << (10 * iota)
	kb
	mb
	gb
	tb
)

// UseConst 使用const
func UseConst() {
	fmt.Println(e)
	fmt.Println(pi2)
	fmt.Printf("%b\n", flag1)
	fmt.Printf("%b\n", flag2)
	fmt.Printf("%b\n", flag3)
	fmt.Printf("%b\n", flag4)

	num := 10
	f := 10.1
	r := float64(num) / f //必须要强制转换
	fmt.Println(r)
	const num1 = 10
	const num2 = 10.1
	r2 := num1 / num2
	fmt.Println(r2)

	var ff float64 = 3 //3是无类型常量 隐式转换成float
	ff = 2
	ff = 'a' //'a'是无类型常量, 隐式转换成float64
	fmt.Println(ff)

}
