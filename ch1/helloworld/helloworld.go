// main 包比较特殊。它定义了一个独立可执行的程序，而不是一个库
package main

import "fmt"

// 没有类层次结构，甚至可以说没有类；仅仅通过组合（而不是继承）简单的对象来构建复杂的对象

// 在 main 包里的 main 函数也很特殊，它是整个程序执行时的入口
func main() {
	// 不需要在语句或者声明的末尾添加分号，除非一行上有多条语句
	fmt.Println("Hello world")
}