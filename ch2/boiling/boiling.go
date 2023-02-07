package main

// 包声明语句之后是 import 语句导入依赖的其它包
import "fmt"

// 四种类型的声明语句：var、const、type 和 func，分别对应变量、常量、类型和函数

func main() {
	// 函数内部的名字则必须先声明之后才能使用
	var f = boilingF
	c := (f - 32) * 5 / 9
	fmt.Printf("boiling point = %g°F or %g°C\n", f, c)
}

// 包一级的各种类型的声明语句的顺序无关紧要
const boilingF = 212.0
