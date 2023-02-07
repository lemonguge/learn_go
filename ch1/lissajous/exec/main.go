package main

import (
	"math/rand"
	"os"
	"time"

	"lemonguge.cn/learn_go/ch1/lissajous"
)

// go run main.go > lissajous.gif
func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	lissajous.Draw(os.Stdout)
}
