package main

import (
	"fmt"
	"log"
	"os"
)

// 声明语句的作用域是指源代码中可以有效使用这个名字的范围，是一个编译时的属性。
// 变量的生命周期是指程序运行时变量存在的有效时间段，在此时间区域内它可以被程序的其他部分引用，是一个运行时的概念。

// 一个程序可能包含多个同名的声明，只要它们在不同的词法域就没有关系。例如可以声明一个局部变量，和包级的变量同名。
// 当编译器遇到一个名字引用时，它会对其定义进行查找，查找过程从最内层的词法域向全局的作用域进行。

func f() {}

var g = "g"

var cwd string

func init() {
	// 虽然 cwd 在外部已经声明过，但是 := 语句还是将 cwd 和 err 重新声明为新的局部变量
	cwd, err := os.Getwd() // NOTE: wrong!
	// var err error
	// cwd, err = os.Getwd()

	if err != nil {
		log.Fatalf("os.Getwd failed: %v", err)
	}
	log.Printf("Working directory = %s", cwd)
}

func main() {
	f()
	f := "f"
	fmt.Println(f) // "f"; local var f shadows package-level func f
	fmt.Println(g) // "g"; package-level var
	g := "f-g"
	fmt.Println(g) // "f-g"; function-level var

	// 未初始化
	fmt.Println(cwd)

	// 有三个不同的 x 变量，一个在函数体词法域，一个在 for 隐式的初始化词法域，一个在 for 循环体词法域
	x := "hello"
	for _, x := range x {
		x := x + 'A' - 'a'
		fmt.Printf("%c", x) // "HELLO" (one letter per iteration)
	}
}
