package template2

import (
	"text/template"
	"time"
)

// 一个模板是一个字符串或一个文件，里面包含了一个或多个由双花括号包含的 {{action}} 对象。
// 大部分的字符串只是按字面值打印，但是对于 actions 部分将触发其它的行为。
// 每个 actions 都包含了一个用模板语言书写的表达式，一个 action 虽然简短但是可以输出复杂的打印值，
// 模板语言包含通过选择结构体的成员、调用函数或方法、表达式控制流 if-else 语句和 range 循环语句，还有其它实例化模板等诸多特性。

// 对于每一个 action，都有一个当前值的概念，对应点操作符，写作“.”。
// 模板中 {{range .Items}} 和 {{end}} 对应一个循环 action，每次迭代的当前值对应当前的 Items 元素的值。
// 在一个 action 中，“|”操作符表示将前一个表达式的结果作为后一个函数的输入，类似于 UNIX 中管道的概念。
const templ = `{{.TotalCount}} issues:
{{range .Items}}----------------------------------------
Number: {{.Number}}
User:   {{.User.Login}}
Title:  {{.Title | printf "%.64s"}}
Age:    {{.CreatedAt | daysAgo}} days
{{end}}`

// template.New 函数先创建并返回一个模板；Funcs 方法将自定义函数注册到模板中，并返回模板；最后调用 Parse 函数分析模板。
// template.Must 函数可以检测 error 是否为 nil（如果不是 nil 则发出 panic 异常），然后返回传入的模板。
var report = template.Must(template.New("issuelist").
	Funcs(template.FuncMap{"daysAgo": daysAgo}).
	Parse(templ))

func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}
