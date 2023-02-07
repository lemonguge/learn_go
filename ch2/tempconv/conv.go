package tempconv

func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

// Celsius(t) 和 Fahrenheit(t) 是类型转换操作，将 float64 转为对应的类型，并不是函数调用

func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}
