package main

import (
	"bytes"
	"fmt"
	"strconv"
)

// intsToString is like fmt.Sprint(values) but adds commas.
func intsToString(values []int) string {
	var buf bytes.Buffer
	buf.WriteByte('[')
	for i, v := range values {
		if i > 0 {
			buf.WriteString(", ")
		}
		fmt.Fprintf(&buf, "%d", v)
	}
	buf.WriteByte(']')
	return buf.String()
}

func main() {
	fmt.Println(intsToString([]int{1, 2, 3})) // "[1, 2, 3]"

	x := 123
	// 将一个整数转为字符串的两种方法
	y := fmt.Sprintf("%d", x)
	fmt.Println(y, strconv.Itoa(x)) // "123 123"

	// 用不同的进制来格式化数字
	s := fmt.Sprintf("%b", x)                   // "x=1111011"
	fmt.Println(s)                              // 1111011
	fmt.Println(strconv.FormatInt(int64(x), 2)) // 1111011
	fmt.Printf("%b\n", x)                       // 1111011

	i, _ := strconv.Atoi("123")             // x is an int
	j, _ := strconv.ParseInt("123", 10, 64) // base 10, up to 64 bits
	fmt.Println(i, j)
}
