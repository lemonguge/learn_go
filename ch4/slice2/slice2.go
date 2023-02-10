package main

import (
	"fmt"

	"lemonguge.cn/learn_go/ch4/rev"
)

// Slice（切片）代表变长的序列，序列中每个元素都有相同的类型。
// 一般写作 []T，其中 T 代表 slice 中元素的类型，和数组很像，只是没有固定长度而已。

// 一个 slice 是一个轻量级的数据结构，提供了访问数组子序列（或者全部）元素的功能，而且 slice 的底层确实引用一个数组对象。
// 一个 slice 由三个部分构成：指针、长度和容量。
// 指针指向第一个 slice 元素对应的底层数组元素的地址；
// 长度对应 slice 中元素的数目，长度不能超过容量；容量一般是从 slice 的开始位置到底层数据的结尾位置。
// 内置的 len 和 cap 函数分别返回 slice 的长度和容量。

func main() {
	// 第 0 个元素会被自动初始化为空字符串
	months := [...]string{1: "January", 2: "February", 3: "March", 4: "April", 5: "May", 6: "June",
		7: "July", 8: "August", 9: "September", 10: "October", 11: "November", 12: "December"}
	fmt.Println(months[1:])

	Q2 := months[4:7]
	summer := months[6:9]
	fmt.Println(Q2)          // ["April" "May" "June"]
	fmt.Println(summer)      // ["June" "July" "August"]
	fmt.Println(len(summer)) // 3
	fmt.Println(cap(summer)) // 7

	// fmt.Println(summer[:20]) // panic: out of range
	endlessSummer := summer[:5] // extend a slice (within capacity)
	fmt.Println(endlessSummer)  // "[June July August September October]"

	// 向函数传递 slice 将允许在函数内部修改底层数组的元素。
	a := [...]int{0, 1, 2, 3, 4, 5}
	rev.Reverse(a[:])
	fmt.Println(a) // "[5 4 3 2 1 0]"

	// 和数字的初始化差异：slice 并没有指明序列的长度。
	b := []int{0, 1, 2, 3, 4, 5}
	rev.Reverse(b)
	fmt.Println(b)

	// 和数组不同的是，slice 之间不能比较，因此不能使用 == 操作符来判断两个 slice 是否含有全部相等元素。
	// fmt.Println(a[:] == b) // compile error: cannot compare
	fmt.Println(equal(a[:], b)) // true
	// slice 只能和 nil 比较
	fmt.Println(b == nil) // false

	// 有两个原因。第一个原因，一个slice的元素是间接引用的，一个slice甚至可以包含自身（译注：当slice声明为[]interface{}时，slice的元素可以是自身）。虽然有很多办法处理这种情形，但是没有一个是简单有效的。
	// 第二个原因，因为slice的元素是间接引用的，一个固定的slice值（译注：指slice本身的值，不是元素的值）在不同的时刻可能包含不同的元素，因为底层数组的元素可能会被修改。

	var s []int
	fmt.Println(s == nil, len(s)) // true 0
	s = nil
	fmt.Println(s == nil, len(s)) // true 0
	// 类型转换
	s = []int(nil)
	fmt.Println(s == nil, len(s)) // true 0
	s = []int{}
	fmt.Println(s == nil, len(s)) // false 0
	// 测试一个 slice 是否是空的，使用 len(s) == 0 来判断，而不应该用 s == nil 来判断。

	// make([]T, len)
	p := make([]int, 3)
	fmt.Println(p)
	// make([]T, len, cap)
	p = make([]int, 3, 5) // same as make([]T, cap)[:len]
	fmt.Println(p)
	fmt.Println(make([]int, 5)[:3])

	// 内置的 append 函数用于向 slice 追加元素
	var runes []rune
	for _, r := range "Hello, 世界" {
		// 因为不能确认 append 调用是否导致了内存的重新分配，所以通常是将 append 返回的结果直接赋值给输入的 slice 变量
		runes = append(runes, r)
	}
	fmt.Printf("%q\n", runes) // ['H' 'e' 'l' 'l' 'o' ',' ' ' '世' '界']
	runes = []rune("Hello, 世界")
	fmt.Printf("%q\n", runes) // ['H' 'e' 'l' 'l' 'o' ',' ' ' '世' '界']
}

func equal(x, y []int) bool {
	if len(x) != len(y) {
		return false
	}
	for i := range x {
		if x[i] != y[i] {
			return false
		}
	}
	return true
}
