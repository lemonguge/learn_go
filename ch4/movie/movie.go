package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Movie struct {
	Title string
	// 一个结构体成员 Tag 是和在编译阶段关联到该成员的元信息字符串
	// 通常是一系列用空格分隔的 key:"value" 键值对序列，因为值中含有双引号字符，成员 Tag 一般用原生字符串面值的形式书写。
	// json 开头键名对应的值用于控制 encoding/json 包的编码和解码的行为，第一部分用于指定 JSON 对象的名字
	Year int `json:"released"`
	// omitempty 选项，表示当 Go 语言结构体成员为空或零值时不生成该 JSON 对象
	Color  bool `json:"color,omitempty"`
	Actors []string
}

func main() {
	var movies = []Movie{
		{Title: "Casablanca", Year: 1942, Color: false,
			Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}},
		{Title: "Cool Hand Luke", Year: 1967, Color: true,
			Actors: []string{"Paul Newman"}},
		{Title: "Bullitt", Year: 1968, Color: true,
			Actors: []string{"Steve McQueen", "Jacqueline Bisset"}},
		// ...
	}
	// 将一个类似 movies 的结构体 slice 转为 JSON 的过程叫编组（marshaling）
	data, err := json.Marshal(movies)
	// json.MarshalIndent 函数将产生整齐缩进的输出。有两个额外的字符串参数用于表示每一行输出的前缀和每一个层级的缩进。
	// data, err := json.MarshalIndent(movies, "", "    ")
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}
	fmt.Printf("%s\n", data)

	// 将 JSON 数据解码为 Go 语言的数据结构的过程叫解组（unmarshaling）
	var titles []struct{ Title string }
	if err := json.Unmarshal(data, &titles); err != nil {
		log.Fatalf("JSON unmarshaling failed: %s", err)
	}
	fmt.Println(titles) // [{Casablanca} {Cool Hand Luke} {Bullitt}]
}
