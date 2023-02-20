package main

import "fmt"

// Go 语言中接口类型的独特之处在于它是满足隐式实现的。也就是说，我们没有必要对于给定的具体类型定义所有满足的接口类型。
// 可以让你创建一个新的接口类型满足已经存在的具体类型却不会去改变这些类型的定义，当我们使用的类型来自于不受我们控制的包时这种设计尤其有用。

// 接口类型是一种抽象的类型，只会表现出它们的方法。

func main() {
	// io.Writer 接口类型，File 结构体类型
	fmt.Printf("%d: %s\n", 1, "hello")

	c := ByteCounter(0)
	fmt.Fprintln(&c, "hello") // 1: hello
	fmt.Println(c)            // 6
	fmt.Println(&c)           // ByteCounter(6)
}

type ByteCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p)) // convert int to ByteCounter
	return len(p), nil
}

func (c *ByteCounter) String() string {
	return fmt.Sprintf("ByteCounter(%d)", *c)
}
