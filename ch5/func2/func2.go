package main

import (
	"fmt"
	"strings"
)

func square(n int) int     { return n * n }
func negative(n int) int   { return -n }
func product(m, n int) int { return m * n }

func addOne(r rune) rune { return r + 1 }

func main() {
	f := square
	fmt.Println(f(3)) // 9

	f = negative
	fmt.Println(f(3))     // -3
	fmt.Printf("%T\n", f) // func(int) int

	// f = product // compile error: can't assign func(int, int) int to func(int) int
	p := product
	fmt.Println(p(3, 7))

	// 函数类型的零值是 nil，函数值之间是不可比较的，不能将函数值作为 map 的 key。
	var f2 func(int) int
	fmt.Println(f2 == nil) // true
	// f2(3) // panic: runtime error
	f2 = square
	fmt.Println(f2(5)) // 25

	fmt.Println(strings.Map(addOne, "HAL-9000")) // IBM.:111
	fmt.Println(strings.Map(addOne, "VMS"))      // WNT
	fmt.Println(strings.Map(addOne, "Admix"))    // Benjy

	// 匿名函数
	fmt.Println(strings.Map(func(r rune) rune { return r + 1 }, "HAL-9000")) // IBM.:111

}
