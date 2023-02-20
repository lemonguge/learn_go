package main

import "fmt"

// 通常来说，不应该对 panic 异常做任何处理，但有时可以从异常中恢复，在程序崩溃前做一些操作。

// 如果在 deferred 函数中调用了内置函数 recover，当定义该 defer 语句的函数发生了 panic 异常，recover 会使程序从 panic 中恢复，并返回 panic value。
// 导致 panic 异常的函数不会继续运行，但能正常返回。在未发生 panic 时调用 recover，recover 会返回 nil。

// 不加区分的恢复所有的 panic 异常，不是可取的做法！把对 panic 的处理都集中在一个包下，有助于简化对复杂和不可以预料问题的处理。
// 作为被广泛遵守的规范，你不应该试图去恢复其他包引起的 panic。公有的 API 应该将函数的运行失败作为 error 返回，而不是 panic。

func divide(x, y int) int {
	defer func() {
		if p := recover(); p != nil {
			fmt.Printf("internal error: %v", p)
		}
	}()
	return x / y
}

func main() {
	fmt.Println(divide(10, 3))
	fmt.Println(divide(10, 0))
	fmt.Println(divide(10, 7))

	isZero(-1)
	isZero(0)
	isZero(1)
}

func isZero(i int) {
	type bailout struct{}
	defer func() {
		switch p := recover(); p {
		case nil:
			fmt.Println("value is 0") // no panic
		case bailout{}:
			fmt.Println("less than 0")
		default:
			fmt.Println(p)
		}
	}()
	if i < 0 {
		panic(bailout{})
	}
	if i > 0 {
		panic("greater than 0")
	}
}
