package main

import (
	"crypto/sha256"
	"fmt"
)

// 数组和结构体是聚合类型，它们的值由许多元素或成员字段的值组成。
// 数组是由同构的元素组成，每个数组元素都是完全相同的类型；结构体则是由异构的元素组成的。
// 数组和结构体都是有固定内存大小的数据结构。相比之下，slice 和 map 则是动态的数据结构，将根据需要动态增长。

// 货币
type Currency int

const (
	USD Currency = iota // 美元
	EUR                 // 欧元
	GBP                 // 英镑
	RMB                 // 人民币
)

func main() {
	// 数组的每个元素都被初始化为元素类型对应的零值，对于数字类型来说就是 0
	var a [3]int      // array of 3 integers
	fmt.Println(a[0]) // print the first element
	// 内置的 len 函数将返回数组中元素的个数。
	fmt.Println(a[len(a)-1]) // print the last element, a[2]

	// Print the indices and elements.
	for i, v := range a {
		fmt.Printf("%d %d\n", i, v)
	}

	// Print the elements only.
	for _, v := range a {
		fmt.Printf("%d\n", v)
	}

	// 也可以使用数组字面值语法用一组值来初始化数组
	var q [3]int = [3]int{1, 2, 3}
	var r [3]int = [3]int{1, 2}
	fmt.Println(q)    // [1 2 3]
	fmt.Println(r[2]) // 0

	// 在数组字面值中，如果在数组的长度位置出现的是“...”省略号，则表示数组的长度是根据初始化值的个数来计算。
	p := [...]int{1, 2, 3}
	fmt.Printf("%T\n", p) // [3]int

	// 数组的长度必须是常量表达式（需在编译阶段确定），作为数组类型的一个组成部分，[3]int 和 [4]int 是两种不同的数组类型。
	// p = [4]int{1, 2, 3}   // compile error

	// 上面的形式是直接提供顺序初始化值序列，也可以指定一个索引和对应值列表的方式初始化。
	symbol := [...]string{RMB: "￥", USD: "$", GBP: "￡", EUR: "€"}
	fmt.Println(RMB, symbol[RMB]) // "3 ￥"

	// 没用到的索引可以省略，未指定初始值的元素将用零值初始化。
	c := [...]int{99: -1}
	fmt.Println(len(c))

	// 数组类型是可以相互比较的，只有当两个数组的所有元素都是相等的时候数组才是相等的。
	fmt.Println(p == q, q == r) // true false

	s := sha256.Sum256([]byte("x"))
	s2 := sha256.Sum256([]byte("X"))
	// %x 以十六进制的格式打印数组或 slice 全部的元素
	fmt.Printf("%x\n%x\n%t\n%T\n", s, s2, s == s2, s)

	// 函数的每个调用参数将会被赋值给函数内部的参数变量，所以函数参数变量接收的是一个复制的副本，并不是原始调用的变量。
	// 导致传递大的数组类型将是低效的，并且对数组参数的任何的修改都是发生在复制的数组上，并不能直接修改调用时原始的数组变量。
	// 可以显式地传入一个数组指针，那样的话函数通过指针对数组的任何修改都可以直接反馈到调用者。
	m := [...]int{1, 2, 3, 4}
	fmt.Println(m)
	zero(&m)
	fmt.Println(m)
}

// 给 [4]int 类型的数组清零
func zero(ptr *[4]int) {
	*ptr = [4]int{}
}
