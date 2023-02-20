package main

import (
	"fmt"
	"math"
)

// 在每一个合法的方法调用表达式中，也就是下面三种情况里的任意一种情况都是可以的：
// 1、接收器的实际参数和其形式参数是相同的类型，比如两者都是类型 T 或者都是类型 *T；
// 2、接收器实参是类型 T，但接收器形参是类型 *T，编译器会隐式地为我们取变量的地址；
// 3、接收器实参是类型 *T，形参是类型 T，编译器会隐式地为我们取到指针指向的实际变量。

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

	(&q).ScaleBy(0.5)
	fmt.Println(q) // {2 3}
	// 语法糖，编译器会隐式地帮我们用 &q 去调用 ScaleBy 这个方法。
	// 只适用于“变量”，包括 struct 里的字段比如 p.X，以及 array 和 slice 内的元素比如 perim[0]。
	q.ScaleBy(2)
	fmt.Println(q) // {4 6}
	// 不能通过一个无法取到地址的接收器来调用指针方法，比如临时变量的内存地址就无法获取得到
	// Point{1, 2}.ScaleBy(2) // compile error: cannot call pointer method ScaleBy on Point

	var i *IntList
	fmt.Println(i == nil, i.Sum()) // true 0
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

// 当调用一个函数时，会对其每一个参数值进行拷贝，如果一个函数需要更新一个变量，或者函数的其中一个参数实在太大，我们希望能够避免进行这种默认的拷贝，这种情况下我们就需要用到指针了。
// 对应到我们这里用来更新接收器的对象的方法，当这个接受者变量本身比较大时，我们就可以用其指针而不是对象来声明方法
func (p *Point) ScaleBy(factor float64) {
	// 方法的名字是 (*Point).ScaleBy，这里的括号是必须的。

	(*p).X *= factor
	// p.X *= factor
	// 语法糖，编译器在这里也会给我们隐式地插入 * 这个操作符
	p.Y *= factor
}

// 一般会约定如果 Point 这个类有一个指针作为接收器的方法，那么所有 Point 的方法都必须有一个指针接收器，即使是那些并不需要这个指针接收器的函数。

// 只有类型（Point）和指向他们的指针(*Point)，才可能是出现在接收器声明里的两种接收器。
// 为了避免歧义，在声明方法时，如果一个类型名本身是一个指针的话，是不允许其出现在接收器中的。
type P *int

// func (P) f() {} // compile error: invalid receiver type

// 注意：nil 也是一个合法的接收器类型，比如对于 map 或者 slice 来说，nil 是合法的零值。
type IntList struct {
	Value int
	Tail  *IntList
}

func (list *IntList) Sum() int {
	if list == nil {
		return 0
	}
	return list.Value + list.Tail.Sum()
}
