package main

import "fmt"

func main() {
	// 指针是变量对应类型值在内存中的地址，通过指针可以直接读或更新对应变量的值。
	// 并不是每一个值都会有一个内存地址，但是对于每一个变量必然有对应的内存地址。

	// 如果用“var x int”声明语句声明一个 x 变量
	// 那么 &x 表达式（取 x 变量的内存地址）将产生一个指向该整数变量的指针
	// 如果指针名字为 p，那么可以说“p 指针指向变量 x”，或者说“p 指针保存了 x 变量的内存地址”
	// *p 表达式读取指针指向的变量的值，这里为 int 类型的值
	// 因为 *p 对应一个变量，可以出现在赋值语句的左边，表示更新指针所指向的变量的值
	// 指针对应的数据类型是 *int，指针被称之为“指向 int 类型的指针”

	x := 1
	p := &x         // p, of type *int, points to x
	fmt.Println(p)  // 内存地址
	fmt.Println(*p) // "1"
	*p = 2          // equivalent to x = 2
	fmt.Println(x)  // "2"

	// 任何类型的指针的零值都是 nil，如果p指向某个有效变量，那么p != nil测试为 true
	var m, n int
	fmt.Println(m == 0, m == n, &m == &n) // true true false

	var f = f()
	fmt.Println(f)
	fmt.Println(*f)

	v := 1
	incr(&v)
	fmt.Println(incr(&v) == v) // v is 3

	leo := person{10, "leo"}
	// 复合类型传递的是其本身以及里面的值的拷贝
	modify(leo)
	fmt.Println(leo) // {10 leo}

	jim := person{10, "jim"}
	// 引用类型传递的是一个指向底层数据的指针
	modify2(&jim)
	fmt.Println(jim) // {20 jim}
}

func f() *int {
	v := 1
	return &v
}

func incr(p *int) int {
	*p++ // 非常重要：只是增加 p 指向的变量的值，并不改变 p 指针
	return *p
}

func modify(p person) {
	p.age = p.age + 10
}
func modify2(p *person) {
	p.age = p.age + 10
}

type person struct {
	age  int
	name string
}
