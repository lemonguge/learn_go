package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// 内置函数 make 创建空 map
	counts := make(map[string]int)
	// 从程序的标准输入中读取内容
	input := bufio.NewScanner(os.Stdin)
	// 读入下一行，并移除行末的换行符
	for input.Scan() {
		s := input.Text()
		if s == "" {
			break
		}
		counts[s]++
	}
	//  for 循环和 if 语句一样，条件两边也不加括号 ()
	for line, n := range counts {
		if n > 1 {
			// %d 十进制整数，%s 字符串，制表符 \t，换行符 \n
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
