package main

import (
	"fmt"
)

// Go 会在编译时捕获很多错误，但有些错误只能在运行时检查，如数组访问越界、空指针引用等。这些运行时错误会引起 panic 异常。

// 当 panic 异常发生时，程序会中断运行，并立即执行在该 goroutine 中被延迟的函数（defer 机制）。
// 随后，程序崩溃并输出日志信息。日志信息包括 panic value（通常是某种错误信息）和函数调用的堆栈跟踪信息。

// 由于 panic 会引起程序的崩溃，因此 panic 一般用于严重错误，如程序内部的逻辑不一致。
// 对于大部分漏洞，我们应该使用 Go 提供的错误机制，而不是 panic，尽量避免程序的崩溃。

// 在 Go 的 panic机制中，延迟函数的调用在释放堆栈信息之前。

func main() {
	// 不是所有的 panic 异常都来自运行时，直接调用内置的 panic 函数也会引发 panic 异常；
	// panic 函数接受任何值作为参数。当某些不应该发生的场景发生时，我们就应该调用 panic。
	panic(fmt.Sprintf("invalid suit %q", "Joker"))
}
