package main

import (
	"fmt"
	"math"
)

// Go 语言将数据类型分为四类：基础类型、复合类型、引用类型和接口类型。
// 基础类型：数字、字符串和布尔型
// 复合类型：数组和结构体
// 引用类型：指针、切片、字典、函数、通道。对任一引用类型数据的修改都会影响所有该引用的拷贝。

// Go 语言中关于算术运算、逻辑运算和比较运算的二元运算符，它们按照优先级递减的顺序排列：
// *      /      %      <<       >>     &       &^
// +      -      |      ^
// ==     !=     <      <=       >      >=
// &&
// ||

// &      位运算 AND
// |      位运算 OR
// ^      位运算 XOR
// &^     位清空（AND NOT）
// <<     左移
// >>     右移

// 在同一个优先级，使用左优先结合规则，但是使用括号可以明确优先顺序，例如 mask & (1 << 28)

func main() {
	// % 取模运算符的符号和被取模数的符号总是一致的
	fmt.Println(-5 % 3)  // -2
	fmt.Println(-5 % -3) // -2
	fmt.Println(5 % -3)  // 2

	// / 除法运算符依赖于操作数是否全为整数
	fmt.Println(5.0 / 4.0) // 1.25
	fmt.Println(5 / 4)     // 1

	// 当计算结果溢出时，超出的高位 bit 位部分将被丢弃。
	var u uint8 = 255
	fmt.Println(u, u+1, u*u) // "255 0 1"
	u = 0
	// 无符号运算不会产生 -1
	fmt.Println(u, u-1) // "0 255"

	var i int8 = 127
	fmt.Println(i, i+1, i*i) // "127 -128 1"

	// 对于整数，+x 是 0+x 的简写，-x 则是 0-x 的简写；对于浮点数和复数，+x 就是 x，-x 则是 x 的负数。

	var x uint8 = 1<<1 | 1<<5
	var y uint8 = 1<<1 | 1<<2
	fmt.Printf("%08b\n", x) // "00100010", the set {1, 5}
	fmt.Printf("%08b\n", y) // "00000110", the set {1, 2}

	fmt.Printf("%08b\n", x&y)  // "00000010", the intersection {1}
	fmt.Printf("%08b\n", x|y)  // "00100110", the union {1, 2, 5}
	fmt.Printf("%08b\n", x^y)  // "00100100", the symmetric difference {2, 5}
	fmt.Printf("%08b\n", x&^y) // "00100000", the difference {5}
	fmt.Printf("%08b\n", x<<1) // "01000100", the set {2, 6}
	fmt.Printf("%08b\n", x>>1) // "00010001", the set {0, 4}

	// 左移运算用零填充右边空缺的 bit 位，无符号数的右移运算也是用 0 填充左边空缺的 bit 位，但是有符号数的右移运算会用符号位的值填充左边空缺的 bit 位。

	// 算术和逻辑运算的二元操作中必须是相同的类型，将一个值从一种类型转化为另一种类型需要显式的转换。
	var apples int32 = 1
	var oranges int16 = 2
	// var compote int = apples + oranges // compile error
	var compote = int(apples) + int(oranges)
	fmt.Println("compote:", compote)

	// 许多整数之间的相互转换并不会改变数值，只是告诉编译器如何解释这个值；
	// 但是对于将一个大尺寸的整数类型转为一个小尺寸的整数类型，或者是将一个浮点数转为整数，可能会改变数值或丢失精度
	f := 1.99
	fmt.Println(int(f)) // 1

	// 任何大小的整数字面值都可以用以 0 开始的八进制格式书写，例如 0666；
	// 或用以 0x 或 0X 开头的十六进制格式书写（不区分大小写），例如 0xdeadbeef
	o := 0666
	fmt.Println(o) // 438

	// 字符面值通过一对单引号直接包含对应字符
	c := 'a'
	fmt.Println(c) // 97
	c = 'A'
	fmt.Println(c) // 65

	// % 之后的 [1] 副词告诉 Printf 函数再次使用第一个操作数，字符使用 %c 参数打印，或者是用 %q 参数打印带单引号
	fmt.Printf("%d %[1]c %[1]q\n", c)

	// nan == nan is false
	nan := math.NaN()
	fmt.Println(nan) // false false false

	// 布尔值可以和 &&（AND）和 ||（OR）操作符结合，并且有短路行为
}
