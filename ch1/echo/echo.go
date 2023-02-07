package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(len(os.Args))
	// 变量会在声明时直接初始化，字符串类型被隐式地赋予 ""
	var s, sep string
	// i++ 是自增语句，不是表达式，所以 j=i++ 非法，而且 ++ 和 -- 都只能放在变量名后面
	for i := 0; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	// os.Args[0] 是命令本身的名字，其它的元素则是程序启动时传给它的参数
	fmt.Println(s)
}
