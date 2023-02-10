package main

import "fmt"

func nonempty(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}
	// 输入的 slice 和输出的 slice 共享一个底层数组，导致原来的数据可能会被覆盖。
	return strings[:i]
}

func main() {
	data := []string{"one", "", "three"}
	fmt.Printf("%q\n", nonempty(data)) // ["one" "three"]
	fmt.Printf("%q\n", data)           // ["one" "three" "three"]

	s := []int{5, 6, 7, 8, 9}
	fmt.Println(remove(s, 2)) // [5 6 8 9]
	s2 := []int{5, 6, 7, 8, 9}
	fmt.Println(remove2(s2, 2)) // [5 6 9 8]
}

// 删除 slice 中间的某个元素并保存原有的元素顺序
func remove(slice []int, i int) []int {
	// 通过内置的 copy 函数将后面的子 slice 向前依次移动一位
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}

// 删除 slice 中间的某个元素，不保存原有的元素顺序
func remove2(slice []int, i int) []int {
	// 用最后一个元素覆盖被删除的元素
	slice[i] = slice[len(slice)-1]
	return slice[:len(slice)-1]
}
