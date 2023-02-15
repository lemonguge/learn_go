// Title1 prints the title of an HTML document specified by a URL.
package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}
	if post != nil {
		post(n)
	}
}

func title(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	// 在调用普通函数或方法前加上关键字 defer，就完成了 defer 所需要的语法。
	// 当执行到 defer 语句时，函数和参数表达式得到计算，直到包含 defer 语句的函数执行完毕时，defer 后的函数才会被执行，
	// 不论函数是通过 return 正常结束，还是由于 panic 导致的异常结束。
	// 可以在一个函数中执行多条 defer 语句，它们的执行顺序与声明顺序相反。
	defer resp.Body.Close()

	ct := resp.Header.Get("Content-Type")
	if ct != "text/html" && !strings.HasPrefix(ct, "text/html;") {
		return fmt.Errorf("%s has type %s, not text/html", url, ct)
	}
	doc, err := html.Parse(resp.Body)
	if err != nil {
		return fmt.Errorf("parsing %s as HTML: %v", url, err)
	}

	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "title" &&
			n.FirstChild != nil {
			fmt.Println(n.FirstChild.Data)
		}
	}
	forEachNode(doc, visitNode, nil)
	return nil
}

func main() {
	if err := title("http://gopl.io"); err != nil {
		fmt.Fprintf(os.Stderr, "title: %v\n", err)
	}
}
