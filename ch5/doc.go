package main

import (
	"fmt"
	"image"
	"math"
	"net/http"
	"time"

	"golang.org/x/net/html"
)

// 函数声明包括函数名、形式参数列表、返回值列表（可省略）以及函数体。
// func name(parameter-list) (result-list) {
//     body
// }

// 形式参数列表描述了函数的参数名以及参数类型。这些参数作为局部变量，其值由参数调用者提供。
// 如果一组形参或返回值有相同的类型，我们不必为每个形参都写出参数类型。
// 返回值列表描述了函数返回值的变量名以及类型。如果函数返回一个无名变量或者没有返回值，返回值列表的括号是可以省略的。
// 如果一个函数声明不包括返回值列表，那么函数体执行完毕后，不会返回任何值。
// 如果一个函数声明包含返回值列表，该函数必须以 return 语句结尾，除非函数明显无法运行到结尾处。
// 返回值也可以像形式参数一样被命名。在这种情况下，每个返回值被声明成一个局部变量，并根据该返回值的类型，将其初始化为该类型的零值。

// x 和 y 是形参名，返回了一个 float64 类型的值。
func hypot(x, y float64) float64 {
	return math.Sqrt(x*x + y*y)
}

// 如果两个函数形式参数列表和返回值列表中的变量类型一一对应，那么这两个函数被认为有相同的类型或签名。
// 形参和返回值的变量名不影响函数签名，也不影响它们是否可以以省略参数类型的形式表示。
// 每一次函数调用都必须按照声明顺序为所有参数提供实参（参数值）；在函数体中，函数的形参作为局部变量，被初始化为调用者提供的值。
// 实参通过值的方式传递，因此函数的形参是实参的拷贝。对形参进行修改不会影响实参。
// 如果实参包括引用类型（指针、切片、字典、函数、通道），实参可能会由于函数的间接引用被修改。

// 没有函数体的函数声明，表示该函数不是以 Go 实现的。这样的声明定义了函数签名。
// Go 语言使用可变栈，栈的大小按需增加（初始时很小）。在使用递归时不必考虑溢出和安全问题。

func add(x int, y int) int   { return x + y }
func sub(x, y int) (z int)   { z = x - y; return }
func first(x int, _ int) int { return x }
func zero(int, int) int      { return 0 }

// 准确的变量名可以传达函数返回值的含义。尤其在返回值的类型都相同时，就像下面这样：
func Size(rect image.Rectangle) (width, height int)     { return }
func Split(path string) (dir, file string)              { return }
func HourMinSec(t time.Time) (hour, minute, second int) { return }

// 如果一个函数所有的返回值都有显式的变量名，那么该函数的 return 语句可以省略操作数。这称之为 bare return。
func CountWordsAndImages(url string) (words, images int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		// 等价于 return 0, 0, err
		return
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		err = fmt.Errorf("parsing HTML: %s", err)
		// 等价于 return 0, 0, err
		return
	}
	words, images = countWordsAndImages(doc)
	return
	// 等价于 return words, images, nil
}
func countWordsAndImages(n *html.Node) (words, images int) { return }

func main() {
	// 3 和 4 是调用时的传入的实参
	fmt.Println(hypot(3, 4)) // 5

	fmt.Printf("%T\n", add)   // func(int, int) int
	fmt.Printf("%T\n", sub)   // func(int, int) int
	fmt.Printf("%T\n", first) // func(int, int) int
	fmt.Printf("%T\n", zero)  // func(int, int) int

	fmt.Println(math.Sin(math.Pi / 6))
}
