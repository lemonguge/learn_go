// The sum program demonstrates a variadic function.
package main

import "fmt"

// 在函数体中，vals 被看作是类型为 []int 的切片。
func sum(vals ...int) int {
	total := 0
	for _, val := range vals {
		total += val
	}
	return total
}

func f(...int) {}
func g([]int)  {}

func main() {
	fmt.Println(sum())           //  0
	fmt.Println(sum(3))          //  3
	fmt.Println(sum(1, 2, 3, 4)) //  10
	// 在上面的代码中，调用者隐式的创建一个数组，并将原始参数复制到数组中，再把数组的一个切片作为参数传给被调用函数。

	values := []int{1, 2, 3, 4}
	// 如果原始参数已经是切片类型，只需在最后一个参数后加上省略符。
	fmt.Println(sum(values...)) // 10

	fmt.Printf("%T\n", f) // func(...int)
	fmt.Printf("%T\n", g) // func([]int)
}
