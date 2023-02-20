package main

import (
	"fmt"
	"image/color"
	"math"
)

type Point struct{ X, Y float64 }

func (p *Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}

type ColoredPoint struct {
	Point
	Color color.RGBA
}

type ColoredPoint2 struct {
	*Point
	Color color.RGBA
}

// 参考 ch4/embed/embed.go 匿名嵌入
func main() {
	var cp ColoredPoint
	cp.X = 1
	fmt.Println(cp.Point.X) // 1
	cp.Point.Y = 2
	fmt.Println(cp.Y) // 2

	red := color.RGBA{255, 0, 0, 255}
	blue := color.RGBA{0, 0, 255, 255}
	var p = ColoredPoint{Point{1, 1}, red}
	var q = ColoredPoint{Point{5, 4}, blue}

	// 简短的点运算符语法可以用于选择匿名成员嵌套的成员，也可以用于访问它们的方法（递归向下找）。
	// Distance 有一个参数是 Point 类型，但 q 并不是一个 Point 类，必须要显式地选择它。
	// p.Distance(q) // compile error: cannot use q (ColoredPoint) as Point
	fmt.Println(p.Distance(q.Point)) // 5

	p.ScaleBy(2)
	q.ScaleBy(2)
	fmt.Println(p.Distance(q.Point)) // 10

	p2 := ColoredPoint2{&Point{1, 1}, red}
	q2 := ColoredPoint2{&Point{5, 4}, blue}
	fmt.Println(p2.Distance(*q2.Point)) // 5
	q2.Point = p2.Point                 // p and q now share the same Point
	p2.ScaleBy(2)
	fmt.Println(*p2.Point, *q2.Point) // {2 2} {2 2}
}
