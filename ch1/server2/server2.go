package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var mu sync.Mutex
var count int

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)
	// 服务器每一次接收请求处理时都会另起一个 goroutine，这样服务器就可以同一时间处理多个请求
	log.Fatal(http.ListenAndServe("localhost:9009", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	// 竞态条件
	count++
	mu.Unlock()
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock()
}
