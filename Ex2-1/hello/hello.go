package main // mainじゃないとエントリポイントなくなってビルドできない。

import "fmt"

func main() {
	message := "Hello World!" // 関数内でのみ利用できる書き方。
	fmt.Println(message)
}
