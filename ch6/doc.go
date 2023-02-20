package main

import (
	"fmt"
	"time"
)

// 在函数声明时，在其名字之前放上一个变量，即是一个方法。
// 这个附加的参数会将该函数附加到这种类型上，即相当于为这种类型定义了一个独占的方法。

func main() {
	const day = 24 * time.Hour
	// Seconds 方法和 time.Duration 类型关联
	fmt.Println(day.Seconds()) // 86400

}
