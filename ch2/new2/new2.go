package main

import "fmt"

func main() {
	// 另一个创建变量的方法是调用内建的 new 函数
	// new(T) 将创建一个 T 类型的匿名变量，初始化为 T 类型的零值，然后返回变量地址，返回的指针类型为 *T。

	p := new(int)   // p, *int 类型, 指向匿名的 int 变量
	fmt.Println(*p) // "0"
	*p = 2          // 设置 int 匿名变量的值为 2
	fmt.Println(*p) // "2"

	fmt.Println(*newInt())

	m := new(int)
	n := new(int)
	fmt.Println(m == n) // false

	// new 函数使用通常相对比较少，因为对于结构体来说，直接用字面量语法创建新变量的方法会更灵活

	// new 只是一个预定义的函数，它并不是一个关键字
	// 如果在函数内定义了 new 变量名，则无法使用内置的 new 函数
}

func newInt() *int {
	var dummy int
	return &dummy
	// 等价与 return new(int)
}
