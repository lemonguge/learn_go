package main

import (
	"fmt"
	"log"
	"time"
)

func bigSlowOperation() {
	// trace 返回的函数值会在 bigSlowOperation 函数退出时被调用
	defer trace("bigSlowOperation")() // don't forget the extra parentheses
	// ...lots of work...
	time.Sleep(3 * time.Second) // simulate slow operation by sleeping
}

func trace(msg string) func() {
	start := time.Now()
	log.Printf("enter %s", msg)
	return func() { log.Printf("exit %s (%s)", msg, time.Since(start)) }
}

func main() {
	fmt.Println(double(4)) // 8
	fmt.Println(triple(4)) // 12

	bigSlowOperation()
}

func triple(x int) (result int) {
	// 被延迟执行的匿名函数甚至可以修改函数返回给调用者的返回值
	defer func() { result += x }()
	return double(x)
}

func double(x int) (result int) {
	// defer 语句中的函数会在 return 语句更新返回值变量后再执行，
	// 在函数中定义的匿名函数可以访问该函数包括返回值变量在内的所有变量，
	// 所以对匿名函数采用 defer 机制，可以使其观察函数的返回值。
	// 对于有许多 return 语句的函数而言，这个技巧很有用。
	defer func() { fmt.Printf("double(%d) = %d\n", x, result) }()
	return x + x
}
