package usemethod

import (
	"fmt"
	"math"
	"time"
)

// UsePoint 使用point
func UsePoint() {
	p := Point{1, 2}
	fmt.Println(p.Distance(Point{1, 2}))
	fmt.Println(p.Distance1(&Point{1, 2})) //可以使用非指针的接收器 会自动解引用或者去地址
	// var pp *Point = nil
	// pp.Distance1(&Point{1, 2}) // nil也可以作为接收器调用, 会自动解引用或者去地址, 但是会panic
}

// Point 点
type Point struct {
	x, y float64
}

// Pp 一个指针
type Pp *Point

// Distance 距离
func (p Point) Distance(q Point) float64 {
	return math.Hypot(p.x-q.x, p.y-q.y)
}

// Distance1 距离,  相同类型的普通对象和指针不饿能同时定义一个相同的方法
func (p *Point) Distance1(q *Point) float64 {
	return math.Hypot(p.x-q.x, p.y-q.y)
}

// Distance 距离,  相同类型的普通对象和指针不饿能同时定义一个相同的方法
// func (p Pp) Distance(q *Point) float64 { //使用type定义的指针不能作为接收器
// 	return math.Hypot(p.x-q.x, p.y-q.y)
// }

func (p Point) lantch() {
	fmt.Println("Lantch")
}

// 方法表达式
func useFuncExp() {
	p := Point{1, 2}
	f := p.Distance //将对象方法复制个一个对象, 这个对象叫做方法表达式, 绑定了p对象
	f(Point{1, 2})
	//将对象的方法直接复制给一个函数对象参数
	time.AfterFunc(time.Second, p.lantch)
	//否则药定义一个匿名函数
	time.AfterFunc(time.Second, func() {
		p.lantch()
	})

	//将类型的方法复制给对象
	f1 := Point.Distance
	f1(p, Point{1, 2})
	f2 := (*Point).Distance1 // 这里必须加指针
	f2(&p, &Point{1, 2})
}
