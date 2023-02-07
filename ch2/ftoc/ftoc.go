package main

import "fmt"

func main() {
	const freezingF, boilingF = 32.0, 212.0
	fmt.Printf("%g°F = %g°C\n", freezingF, fToC(freezingF))
	fmt.Printf("%g°F = %g°C\n", boilingF, fToC(boilingF))
}

// 名字的开头字母的大小写决定了名字在包外的可见性，大写字母开头的名字可以被外部的包访问
func fToC(f float64) float64 {
	return (f - 32) * 5 / 9
}
