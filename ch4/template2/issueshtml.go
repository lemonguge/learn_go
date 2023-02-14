package template2

import "html/template"

// html/template 模板包使用和 text/template 包相同的 API 和模板语言，但是增加了一个将字符串自动转义特性。
// 可以避免输入字符串和 HTML、JavaScript、CSS 或 URL 语法产生冲突的问题。以及一些长期存在的安全问题，比如 HTML 注入攻击。
var report2 = template.Must(template.New("issuelist").Parse(`
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
