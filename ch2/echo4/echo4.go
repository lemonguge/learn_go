package main

import (
	"flag"
	"fmt"
	"strings"
)

// 命令行标志参数，默认值，描述信息
var n = flag.Bool("n", false, "omit trailing newline")
var sep = flag.String("s", " ", "separator")

func main() {
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}
	// go run ch2/echo4/echo4.go a bc def
	// output: a bc def
	// go run ch2/echo4/echo4.go -s / a bc def
	// output: a/bc/def
	// go run ch2/echo4/echo4.go -n a bc def
	// output: a bc def%

	// go run ch2/echo4/echo4.go -h
	// -h 或 -help 打印所有标志参数的名字、默认值和描述信息
}
