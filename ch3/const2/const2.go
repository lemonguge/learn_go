package main

import (
	"fmt"
	"math"
	"net"
	"time"
)

// 常量表达式的值在编译期计算，而不是在运行期。每种常量的潜在类型都是基础类型：boolean、string 或数字。

// 一个常量的声明也可以包含一个类型和一个值，但是如果没有显式指明类型，那么将从右边的表达式推断类型。
const pi = 3.14

// 批量声明多个常量
const (
	e   = 2.7183
	pi2 = 3.1415
)

const (
	_   = 1 << (10 * iota)
	KiB // 1024
	MiB // 1048576
	GiB // 1073741824
)

func main() {
	fmt.Println(pi, pi2, e)

	const noDelay time.Duration = 0
	const timeout = 5 * time.Minute
	fmt.Printf("%T %[1]v\n", noDelay)     // "time.Duration 0"
	fmt.Printf("%T %[1]v\n", timeout)     // "time.Duration 5m0s"
	fmt.Printf("%T %[1]v\n", time.Minute) // "time.Duration 1m0s"

	// 如果是批量声明的常量，除了第一个外其它的常量右边的初始化表达式都可以省略，如果省略初始化表达式则表示使用前面常量的初始化表达式写法，对应的常量类型也一样的。
	const (
		a = 1
		b
		c = 2
		d
	)
	fmt.Println(a, b, c, d) // "1 1 2 2"

	// 常量声明可以使用 iota 常量生成器初始化，用于生成一组以相似规则初始化的常量。
	// 在一个 const 声明语句中，在第一个声明的常量所在的行，iota 将会被置为 0，然后在每一个有常量声明的行加一。
	fmt.Println(KiB, MiB, GiB)
	// 例如：time 包 Weekday 和 net 包的 Flags
	fmt.Println(time.Sunday, time.Monday)
	fmt.Println(net.FlagUp, net.FlagBroadcast)

	// 编译器为这些没有明确基础类型的数字常量提供比基础类型更高精度的算术运算
	// 无类型的布尔型、无类型的整数、无类型的字符、无类型的浮点数、无类型的复数、无类型的字符串
	fmt.Println(GiB / MiB) // "1024"

	var x float32 = math.Pi
	var y float64 = math.Pi
	var z complex128 = math.Pi
	fmt.Println(x, y, z)

	var f float64 = 212
	fmt.Println((f - 32) * 5 / 9) // "100"; (f - 32) * 5 is a float64
	// fmt.Println(5 / 9 * (f - 32))     // "0";   5/9 is an untyped integer, 0
	fmt.Println(5.0 / 9.0 * (f - 32)) // "100"; 5.0/9.0 is an untyped float

	// 当一个无类型的常量被赋值给一个有明确类型的变量时，无类型的常量将会被隐式转换为对应的类型，如果转换合法的话。
	// 对于一个没有显式类型的变量声明（包括简短变量声明），常量的形式将隐式决定变量的默认类型
}
