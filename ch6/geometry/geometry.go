package main

import (
	"fmt"
	"math"
)

type Point struct{ X, Y float64 }

// 参数 p 叫做方法的接收器（receiver），建议是可以使用其类型的第一个字母
func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

// 由于方法和字段都是在同一命名空间，所以如果我们在这里声明一个 X 方法的话，编译器会报错
// func (p Point) X() {} // compile error: field and method with the same name

func Distance(p, q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func main() {
	p := Point{1, 2}
	q := Point{4, 6}
	// 包级别的函数调用
	fmt.Println(Distance(p, q)) // 5
	// Point 类下声明的方法调用
	fmt.Println(p.Distance(q)) // 5

	// p.Distance 的表达式叫做选择器，也会被用来选择一个 struct 类型的字段，比如 p.X。

	perim := Path{
		{1, 1},
		{5, 1},
		{5, 4},
		{1, 1},
	}
	fmt.Println(perim.Distance()) // 12
}

// 可以给同一个包内的任意命名类型定义方法，只要这个命名类型的底层类型不是指针或者 interface。
type Path []Point

// 因为每种类型都有其方法的命名空间，我们在用 Distance 这个名字的时候，不同的 Distance 调用指向了不同类型里的 Distance 方法。
func (path Path) Distance() float64 {
	sum := 0.0
	for i := range path {
		if i > 0 {
			sum += path[i-1].Distance(path[i])
		}
	}
	return sum
}
