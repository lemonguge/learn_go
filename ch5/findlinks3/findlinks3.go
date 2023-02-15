// Findlinks3 crawls the web, starting with the URLs on the command line.
package main

import (
	"fmt"
	"log"

	"lemonguge.cn/learn_go/ch5/links"
)

// breadthFirst calls f for each item in the worklist.
// Any items returned by f are added to the worklist.
// f is called at most once for each item.
func breadthFirst(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			fmt.Println("crawl:", item)
			if !seen[item] {
				seen[item] = true
				// 将 f 返回的一组元素一个个添加到 worklist 中
				worklist = append(worklist, f(item)...)
			}
		}
	}
}

func crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}

func main() {
	// Crawl the web breadth-first,
	// starting from the command-line arguments.
	breadthFirst(crawl, []string{"https://golang.org"})
}
