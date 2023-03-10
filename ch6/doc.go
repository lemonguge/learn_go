package main

import (
	"fmt"
	"time"
)

// 在函数声明时，在其名字之前放上一个变量，即是一个方法。
// 这个附加的参数会将该函数附加到这种类型上，即相当于为这种类型定义了一个独占的方法。

// Go 语言只有一种控制可见性的手段：大写首字母的标识符会从定义它们的包中被导出，小写字母的则不会。
// 最小的封装单元是 package，而不是像其它语言一样的类型。封装提供了三方面的优点：
// 1、提高代码的可复用行；
// 2、隐藏实现的细节，可以防止调用方依赖那些可能变化的具体实现；
// 3、阻止了外部调用方对对象内部的值任意地进行修改。

func main() {
	const day = 24 * time.Hour
	// Seconds 方法和 time.Duration 类型关联
	fmt.Println(day.Seconds()) // 86400

}
