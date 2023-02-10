package main

import (
	"fmt"
	"os"
)

// var 变量名字 类型 = 表达式

// 其中“类型”或“= 表达式”两个部分可以省略其中的一个。
// 如果省略的是类型信息，那么将根据初始化表达式来推导变量的类型信息；
// 如果初始化表达式被省略，那么将用零值初始化该变量。

// 数值类型变量对应的零值是 0，布尔类型变量对应的零值是 false，字符串类型对应的零值是 ""；
// 接口或引用类型（包括 slice、指针、map、chan 和函数）变量对应的零值是 nil；
// 数组或结构体等聚合类型对应的零值是每个元素或字段都是对应该类型的零值。

func main() {
	var s string
	fmt.Println(s) // ""

	// 可以在一个声明语句中同时声明一组变量，或用一组初始化表达式声明并初始化一组变量。
	var i, j, k int                  // int, int, int
	var b, f, s2 = true, 2.3, "four" // bool, float64, string
	fmt.Println(i + j + k)
	fmt.Println("b:", b)
	fmt.Println("f:", f)
	fmt.Println("s2:", s2)

	// 一组变量也可以通过调用一个函数，由函数返回的多个返回值初始化
	var data, err = os.ReadFile("doc/dup.txt")
	if err == nil {
		fmt.Println("data len:", len(data))
	} else {
		fmt.Println("open file error")
	}

	// 简短变量声明：在函数内部用于声明和初始化局部变量，形式为“名字 := 表达式”，变量的类型根据表达式来自动推导
	pi := 3.14
	fmt.Println("pi:", pi)

	// 简短变量声明语句也可以用来声明和初始化一组变量
	p, q := 3, 7
	fmt.Printf("p:%d, q:%d\n", p, q)
	// “:=”是一个变量声明语句，而“=”是一个变量赋值操作
	p, q = q, p // 元组赋值，交换 p 和 q 的值（在赋值之前，赋值语句右边的所有表达式将会先进行求值，然后再统一更新左边对应变量的值）
	fmt.Printf("p:%d, q:%d\n", p, q)

	// 有些表达式会产生多个值，当这样一个函数调用出现在元组赋值右边的表达式中时，左边变量的数目必须和右边一致

	// 简短变量声明语句也可以用函数的返回值来声明和初始化变量，对已经声明过的变量就只有赋值行为，但至少要声明一个新的变量
	data2, err := os.ReadFile("doc/dup.txt")
	if err == nil {
		fmt.Println("data len:", len(data2))
	} else {
		fmt.Println("open file error")
	}

	// 隐式地对 slice 的每个元素进行赋值操作（隐式赋值）
	medals := []string{"gold", "silver", "bronze"}
	fmt.Println(medals)                     // [gold silver bronze]
	fmt.Println(len(medals) == cap(medals)) // true

	// 如果将指向短生命周期对象的指针保存到具有长生命周期的对象中，特别是保存到全局变量时，会阻止对短生命周期对象的垃圾回收
}
