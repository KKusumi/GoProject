package main

import "fmt"

// 大文字だと他モジュールから参照できるようになってしまう。
// 命名規則的にはキャメルにするのが正しい？
const message string = "Hello World!"

func main() {
	fmt.Println(message)
	omit()
}

func omit() {
	const (
		a = 3
		b // 右辺を省略して全て3にできる。あんま使わん予感する。
		c
	)
	fmt.Println(a, b, c)
}
