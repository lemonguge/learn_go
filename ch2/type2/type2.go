package main

import (
	"fmt"

	"lemonguge.cn/learn_go/ch2/tempconv"
)

func main() {
	// 一个字符串可以用来表示一个密码或者一个颜色的名称

	// type 类型名字 底层类型
	// 一个类型声明语句创建了一个新的类型名称，和现有类型具有相同的底层结构
	// 类型声明语句一般出现在包一级，如果新创建的类型名字的首字符大写，则在包外部也可以使用

	// 底层数据类型决定了内部结构和表达方式，也决定是否可以像底层类型一样对内置运算符的支持
	// 对于每一个类型 T，都有一个对应的类型转换操作 T(x)，用于将 x 转为 T 类型
	// 如果 T 是指针类型，需要用小括弧包装 T，比如：(*int)(0)

	fmt.Printf("%g\n", tempconv.BoilingC-tempconv.FreezingC) // "100" °C
	boilingF := tempconv.CToF(tempconv.BoilingC)
	fmt.Printf("%g\n", boilingF-tempconv.CToF(tempconv.FreezingC)) // "180" °F
	// fmt.Printf("%g\n", boilingF-tempconv.FreezingC) // compile error: type mismatch

	var c tempconv.Celsius
	var f tempconv.Fahrenheit
	fmt.Println(c == 0) // true
	fmt.Println(f >= 0) // true
	// fmt.Println(c == f) // compile error: type mismatch

	// 注意：Celsius(f) 是类型转换操作，并不会改变值，仅仅是改变值的类型而已
	// 只有当两个类型的底层基础类型相同时，才允许这种转型操作；
	// 或者是两者都是指向相同底层结构的指针类型，这些转换只改变类型而不会影响值本身。
	fmt.Println(c == tempconv.Celsius(f)) // true

	// 数值类型之间的转型也是允许的，并且在字符串和一些特定类型的 slice 之间也是可以转换的
	// 例如，将一个浮点数转为整数将丢弃小数部分。在任何情况下，运行时不会发生转换失败的错误（错误只会发生在编译阶段）

	c = tempconv.FToC(212.0)
	// 当使用 fmt 包的打印方法时，将会优先使用该类型对应的 String 方法返回的结果打印
	fmt.Println(c) // 100°C
}
