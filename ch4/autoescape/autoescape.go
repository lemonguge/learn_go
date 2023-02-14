package main

import (
	"html/template"
	"log"
	"os"
)

// go run autoescape.go > autoescape.html
func main() {
	const templ = `<p>A: {{.A}}</p><p>B: {{.B}}</p>`
	t := template.Must(template.New("escape").Parse(templ))
	var data struct {
		A string        // untrusted plain text
		B template.HTML // trusted HTML
	}
	// 普通字符串，被自动转义
	data.A = "<b>Hello!</b>"
	// 信任的 template.HTML 字符串，不会被转义
	data.B = "<b>Hello!</b>"

	if err := t.Execute(os.Stdout, data); err != nil {
		log.Fatal(err)
	}
}
