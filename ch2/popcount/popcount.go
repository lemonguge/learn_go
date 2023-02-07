package popcount

// pc[i] is the population count of i.
var pc [256]byte

// init 函数初始化 pc 变量
func init() {
	// range 循环只使用了索引，省略了没有用到的值部分
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// 也可以通过将初始化逻辑包装为一个匿名函数处理
// var pc [256]byte = func() (pc [256]byte) {
//     for i, _ := range pc {
//         pc[i] = pc[i/2] + byte(i&1)
//     }
//     return
// }()

// PopCount returns the population count (number of set bits) of x.
func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}
