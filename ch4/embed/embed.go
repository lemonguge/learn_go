package main

import "fmt"

type Point struct {
	X, Y int
}

type Circle struct {
	Center Point
	Radius int
}

type Wheel struct {
	Circle Circle
	Spokes int
}

type Circle2 struct {
	// 匿名成员：只声明一个成员对应的数据类型而不指名成员的名字，匿名成员的数据类型必须是命名的类型或指向一个命名的类型的指针。
	// 因为匿名成员也有一个隐式的名字，因此不能同时包含两个类型相同的匿名成员，这会导致名字冲突。
	// 因为成员的名字是由其类型隐式地决定的，所以匿名成员也有可见性的规则约束。
	// 匿名成员并不要求是结构体类型，任何命名的类型都可以作为结构体的匿名成员。
	Point
	Radius int
}

type Wheel2 struct {
	Circle2
	Spokes int
}

func main() {
	var w Wheel
	// 访问每个成员将变得繁琐
	w.Circle.Center.X = 8
	w.Circle.Center.Y = 6
	w.Circle.Radius = 5
	w.Spokes = 20
	fmt.Println(w)

	// 语法糖：匿名嵌入的特性，可以直接访问叶子属性而不需要给出完整的路径。
	// 简短的点运算符语法可以用于选择匿名成员嵌套的成员，也可以用于访问它们的方法。
	// 外层的结构体不仅仅是获得了匿名成员类型的所有成员，而且也获得了该类型导出的全部方法。
	var w2 Wheel2
	w2.X = 8      // equivalent to w2.Circle2.Point.X = 8
	w2.Y = 6      // equivalent to w2.Circle2.Point.Y = 6
	w2.Radius = 5 // equivalent to w2.Circle2.Radius = 5
	w2.Spokes = 20

	// 不幸的是，结构体字面值并没有简短表示匿名成员的语法， 因此下面的语句都不能编译通过。
	// w2 = Wheel2{8, 6, 5, 20}                       // compile error: unknown fields
	// w2 = Wheel2{X: 8, Y: 6, Radius: 5, Spokes: 20} // compile error: unknown fields

	// 下面的两种语法彼此是等价的
	w2 = Wheel2{Circle2{Point{8, 6}, 5}, 20}
	fmt.Printf("%#v\n", w2)
	w2 = Wheel2{
		Circle2: Circle2{
			Point:  Point{X: 8, Y: 6},
			Radius: 5,
		},
		Spokes: 20, // NOTE: trailing comma necessary here (and at Radius)
	}
}
