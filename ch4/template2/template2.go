package main

import (
	"fmt"
	ht "html/template"
	"log"
	"os"
	tt "text/template"
	"time"

	"lemonguge.cn/learn_go/ch4/github"
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
var report = tt.Must(tt.New("issuelist").
	Funcs(tt.FuncMap{"daysAgo": daysAgo}).
	Parse(templ))

// html/template 模板包使用和 text/template 包相同的 API 和模板语言，但是增加了一个将字符串自动转义特性。
// 可以避免输入字符串和 HTML、JavaScript、CSS 或 URL 语法产生冲突的问题。以及一些长期存在的安全问题，比如 HTML 注入攻击。
var report2 = ht.Must(ht.New("issuelist").Parse(`
<h1>{{.TotalCount}} issues</h1>
<table>
<tr style='text-align: left'>
  <th>#</th>
  <th>State</th>
  <th>User</th>
  <th>Title</th>
</tr>
{{range .Items}}
<tr>
  <td><a href='{{.HTMLURL}}'>{{.Number}}</a></td>
  <td>{{.State}}</td>
  <td><a href='{{.User.HTMLURL}}'>{{.User.Login}}</a></td>
  <td><a href='{{.HTMLURL}}'>{{.Title}}</a></td>
</tr>
{{end}}
</table>
`))

func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}

func main() {
	result, err := github.SearchIssues([]string{"repo:golang/go", "is:open", "json", "decoder"})
	if err != nil {
		log.Fatal(err)
	}
	if err := report.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}
	fmt.Println("<===== =====>")
	if err := report2.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}
}
