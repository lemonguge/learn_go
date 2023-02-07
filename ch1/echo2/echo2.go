package main

import (
	"fmt"
	"os"
)

func main() {
	// 短变量声明，只能用在函数内部
	s, sep := "", ""
	// 每次循环迭代，range 产生一对值：索引以及在该索引处的元素值
	// 不允许使用无用的局部变量，会导致编译错误，用空标识符“_”表示变量名
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}
