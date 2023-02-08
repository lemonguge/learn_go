package main

import (
	"fmt"
	"unicode/utf8"
)

// 一个字符串是一个不可改变的字节序列，通常被解释为采用 UTF8 编码的 Unicode 码点（rune）序列

const GoUsage = `Go is a tool for managing Go source code.

Usage:
    go command [arguments]
...`

func main() {
	s := "Hi, 你好"
	// 对于非 ASCII 字符的 UTF8 编码会要两个或多个字节
	fmt.Println(len(s))
	// s[i] 返回第 i 个字节的字节值，并不一定是字符串的第 i 个字符
	fmt.Println(s[0])

	// s[i:j] 基于原始的 s 字符串的第 i 个字节开始到第 j 个字节（并不包含 j 本身）生成一个新字符串
	// 不管 i 还是 j 都可能被忽略，当它们被忽略时将采用 0 作为开始位置，采用 len(s) 作为结束的位置
	fmt.Println(s[:])
	fmt.Println(s[4:])

	// s[0] = 'L' // compile error

	// + 操作符将两个字符串连接构造一个新字符串
	fmt.Println("Hello, " + s[4:])

	// 字符串可以用 == 和 < 进行比较，通过逐个字节比较完成的，因此结果是字符串自然编码的顺序。

	s = "left foot"
	t := s
	s += ", right foot"
	fmt.Println(s) // left foot, right foot
	fmt.Println(t) // left foot

	// 在一个双引号包含的字符串面值中，可以用以反斜杠 \ 开头的转义序列插入任意的数据。
	// 可以通过十六进制或八进制转义在字符串面值中包含任意的字节。
	// 一个十六进制的转义形式是 \xhh，其中两个 h 表示十六进制数字（大写或小写都可以）。
	// 一个八进制的转义形式是 \ooo，包含三个八进制的 o 数字（0到7），但是不能超过 \377（一个字节的范围，十进制为 255）。

	// 一个原生的字符串面值形式是 `...`，使用反引号代替双引号。在原生的字符串面值中，没有转义操作。
	// 在原生字符串面值内部是无法直接写 ` 字符的，可以用八进制或十六进制转义或 + "`" 连接字符串常量完成。

	fmt.Println(GoUsage)

	// ASCII 字符集：美国信息交换标准代码，使用 7bit 来表示 128 个字符。
	// Unicode 提供了涵盖全球各种语言的单一字符集，每个符号都分配一个唯一的 Unicode 码点。
	// Unicode 码点对应 Go 语言中的 rune 整数类型（rune 是 int32 等价类型）
	// 可以将一个符文序列表示为一个 int32 序列。这种编码方式叫 UTF-32 或 UCS-4，每个 Unicode 码点都使用同样大小的 32bit 来表示。简单统一，但是会浪费很多存储空间。

	// UTF8 是一个将 Unicode 码点编码为字节序列的变长编码，使用 1 到 4 个字节来表示每个 Unicode 码点，现在已经是 Unicode 的标准。

	// 包含 13 个字节，以 UTF8 形式编码，但是只对应 9 个 Unicode 字符
	s = "Hello, 世界"
	fmt.Println(len(s))                    // 13
	fmt.Println(utf8.RuneCountInString(s)) // 9
	fmt.Println(s)
	for i := 0; i < len(s); {
		// 返回一个 r 和长度，r 对应字符本身，长度对应 r 采用 UTF8 编码后的编码字节数目。
		r, size := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%d\t%c\n", i, r)
		i += size
	}
	fmt.Println("-------")
	// range 循环在处理字符串的时候，会自动隐式解码 UTF8 字符串。
	for i, r := range "Hello, 世界" {
		fmt.Printf("%d\t%q\n", i, r)
	}

	fmt.Println(string(rune(0x4eac))) // 京
	// 如果对应码点的字符是无效的，则用 \uFFFD 无效字符作为替换
	fmt.Println(string(rune(1234567))) // "?"

	s = "abc"
	// 分配了一个新的字节数组用于保存字符串数据的拷贝
	b := []byte(s)
	s2 := string(b)
	fmt.Println("s2:", s2)
}
