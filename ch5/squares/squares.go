package main

import "fmt"

func squares() func() int {
	// 在函数中定义的内部函数可以引用该函数的变量
	// 变量的生命周期不由它的作用域决定：squares 返回后，变量 x 仍然隐式的存在于 f 中。
	var x int
	// 返回一个匿名函数，每次被调用时都会返回下一个数的平方。
	return func() int {
		x++
		return x * x
	}
}
func main() {
	f := squares()
	fmt.Println(f()) // 1
	fmt.Println(f()) // 4
	fmt.Println(f()) // 9
	fmt.Println(f()) // 16
}
