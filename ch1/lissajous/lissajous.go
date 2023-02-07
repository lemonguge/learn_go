package lissajous

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
)

// 当 import 的包路径包含有多个单词的 package，通常只需要用最后那个单词表示这个包
var palette = []color.Color{color.White, color.Black}

// 包级别的常量在整个包中都是可以共享的，函数内的常量只能在函数体内用
const (
	whiteIndex = 0
	blackIndex = 1
)

// 利萨如图形
func Draw(out io.Writer) {
	// 常量声明的值必须是一个数字值、字符串或者一个固定的 boolean 值
	const (
		cycles  = 5
		res     = 0.001
		size    = 100
		nframes = 64
		// 帧之间的延迟以 10ms 为单位
		delay = 8
	)

	freq := rand.Float64() * 3.0
	// anim 是一个 gif.GIF 类型的 struct 变量，其内部变量 LoopCount 字段被设置为 nframes
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		// 图片像素为 201*201
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		// 每个像素点被默认设置为调色板的第 0 个值（白色）
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			// x 轴偏振使用 sin 函数
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			// 为 (x,y) 点来染黑色
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), blackIndex)
		}
		phase += 0.1
		// 设置一个默认的 80ms 的延迟值
		anim.Delay = append(anim.Delay, delay)
		// 将图片添加到 anim 中的帧列表末尾
		anim.Image = append(anim.Image, img)
	}
	// 指针是一种直接存储了变量的内存地址的数据类型；& 操作符可以返回一个变量的内存地址，* 操作符可以获取指针指向的变量内容
	gif.EncodeAll(out, &anim)
}
