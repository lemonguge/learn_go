package main

import (
	"fmt"
	"sort"
)

// 哈希表是一个无序的 key/value 对的集合，其中所有的 key 都是不同的。
// 通过给定的 key 可以在常数时间复杂度内检索、更新或删除对应的 value。

// Go 语言中，一个 map 就是一个哈希表的引用，map 类型可以写为 map[K]V，其中 K: key 和 V: value。
// map 中所有的 key 都有相同的类型，所有的 value 也有着相同的类型，key 和 value 可以是不同的数据类型。
// 其中 K 对应的 key 必须是支持 == 比较运算符的数据类型，所以 map 可以通过测试 key 是否相等来判断是否已经存在。

// 虽然浮点数类型也是支持相等运算符比较的，但是将浮点数用做 key 类型则是一个坏的想法，可能出现的 NaN 和任何浮点数都不相等。value 数据类型则没有任何的限制。

func main() {
	// make 函数可以创建一个 map
	ages := make(map[string]int)
	ages["alice"] = 31
	ages["charlie"] = 34

	// 用 map 字面值的语法创建 map，同时还可以指定一些最初的 key/value
	ages = map[string]int{
		"alice":   31,
		"charlie": 34,
	}

	// map 中的元素通过 key 对应的下标语法访问：
	ages["alice"] = 32
	fmt.Println(ages["alice"]) // 32

	// delete 函数可以删除元素
	delete(ages, "alice")
	// 查找失败将返回 value 类型对应的零值
	fmt.Println(ages["alice"]) // 0

	// x += y 和 x++ 等简短赋值语法也可以用在 map 上
	ages["bob"] += 1
	fmt.Println(ages["bob"]) // 1
	ages["bob"]++
	fmt.Println(ages["bob"]) // 2

	// 禁止对 map 元素取址的原因是 map 可能随着元素数量的增长而重新分配更大的内存空间，从而可能导致之前的地址无效。

	// map 的迭代顺序是不确定的
	for name, age := range ages {
		fmt.Printf("%s\t%d\n", name, age)
	}

	// 如果要按顺序遍历 key/value 对，我们必须显式地对 key 进行排序
	names := make([]string, 0, len(ages))
	for name := range ages {
		names = append(names, name)
	}
	sort.Strings(names)
	// 使用 _ 空白标识符来忽略迭代 slice 时的索引
	for _, name := range names {
		fmt.Printf("%s\t%d\n", name, ages[name])
	}

	// map 类型的零值是 nil，也就是没有引用任何哈希表。
	var ages2 map[string]int
	fmt.Println(ages2 == nil)    // true
	fmt.Println(len(ages2) == 0) // true

	// 查找、删除、len 和 range 循环都可以安全工作在 nil 值的 map 上，和一个空的 map 类似。
	// 但是向一个 nil 值的 map 存入元素将导致一个 panic 异常，在向 map 存数据前必须先创建 map。
	// ages2["carol"] = 21 // panic: assignment to entry in nil map

	// 如果元素类型是一个数字，你可能需要区分一个已经存在的 0，和不存在而返回零值的 0。
	// 在这种场景下，map 的下标语法将产生两个值，第二个是一个布尔值，用于报告元素是否真的存在。
	_, ok := ages["alice"]
	if !ok {
		fmt.Println("alice is not a key in this map")
	}

	// 和 slice 一样，map 之间也不能进行相等比较，唯一的例外是和 nil 进行比较。
	ages2 = map[string]int{
		"bob":     2,
		"charlie": 34,
	}
	fmt.Println(equal(ages, ages2))

	// Go 语言中并没有提供一个 set 类型，可以用 map 实现类似 set 的功能。

	// 有时候我们需要一个 map 或 set 的 key 是 slice 类型，但是 map 的 key 必须是可比较的类型，但是 slice 并不满足这个条件。
	// 第一步，定义一个辅助函数 k，将 slice 转为 map 对应的 string 类型的 key，确保只有 x 和 y 相等时 k(x) == k(y) 才成立。
	// 然后创建一个 key 为 string 类型的 map，在每次对 map 操作时先用 k 辅助函数将 slice 转化为 string 类型。
}

func equal(x, y map[string]int) bool {
	if len(x) != len(y) {
		return false
	}
	for k, xv := range x {
		// 使用 !ok 来区分元素不存在，与元素存在但为 0 的场景。
		if yv, ok := y[k]; !ok || yv != xv {
			return false
		}
	}
	return true
}
