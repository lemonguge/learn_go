package tempconv

import "fmt"

// Celsius 和 Fahrenheit 虽然有着相同的底层类型 float64，但是它们是不同的数据类型
// Celsius 和 Fahrenheit 类型的算术运算行为和底层的 float64 类型是一样的
type Celsius float64    // 摄氏温度
type Fahrenheit float64 // 华氏温度

const (
	AbsoluteZeroC Celsius = -273.15    // 绝对零度
	BoilingC      Celsius = 100        // 沸水温度
	FreezingC     Celsius = Celsius(0) // 结冰点温度
)

// 命名类型还可以为该类型的值定义新的行为。这些行为表示为一组关联到该类型的函数集合，我们称为类型的方法集
// Celsius 类型的参数 c 出现在了函数名的前面，表示声明的是 Celsius 类型的一个名叫 String 的方法
func (c Celsius) String() string {
	return fmt.Sprintf("%g°C", c)
}

func (f Fahrenheit) String() string {
	return fmt.Sprintf("%g°F", f)
}
