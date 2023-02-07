package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

// goroutine 是一种函数的并发执行方式，而 channel 是用来在 goroutine 之间进行参数传递；main 函数本身也运行在一个 goroutine 中
func main() {
	start := time.Now()
	// 创建了一个传递 string 类型参数的 channel
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		// go function 则表示创建一个新的 goroutine，并在这个新的 goroutine 中执行这个函数
		go fetch(url, ch)
	}
	for range os.Args[1:] {
		// 主函数负责接收这些值（<-ch）
		fmt.Println(<-ch)
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		// 往 channel 里发送一个值（ch <- expression）
		ch <- fmt.Sprint(err)
		return
	}
	// 把响应的 Body 内容拷贝到 io.Discard 输出流中
	nbytes, err := io.Copy(io.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs  %7d  %s", secs, nbytes, url)
}
