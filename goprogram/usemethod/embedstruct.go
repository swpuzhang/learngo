package usemethod

import (
	"fmt"
	"image/color"
	"sync"
)

// ColorPoint point
type ColorPoint struct {
	Point
	color color.RGBA
}

// UseEmbedStruct 使用内嵌结构体
func UseEmbedStruct() {
	// 两种初始化方式
	p := ColorPoint{Point{1, 2}, color.RGBA{1, 2, 3, 4}}
	pp := &ColorPoint{Point: Point{1, 2}, color: color.RGBA{1, 2, 3, 4}}
	// 内嵌结构体字段名就是类型名
	p.Distance(pp.Point)
	pp.Distance1(&p.Point)
	// 直接使用内嵌的成员
	fmt.Println(p.x, p.y)
}

//给匿名类型定义方法

var cache = struct {
	sync.Mutex
	mapping map[string]string
}{
	mapping: make(map[string]string),
}

func useCache(key string) {
	cache.Lock()
	defer cache.Unlock()
	fmt.Printf(cache.mapping[key])
}
