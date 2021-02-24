package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main () {
	ch := input(os.Stdin)
	for {
		fmt.Print(">")
		fmt.Println(<-ch)
	}
}

func input(r io.Reader) <- chan string {
	// TODO: チャネルを作る
	c := make(chan string)
	go func() {
		s := bufio.NewScanner(r)
		for s.Scan() {
			// TODO: チャネルに読み込んだ文字列を送る
			c <- s.Text()
		}
		// TODO: チャネルを閉じる
		close(c)
	}()
	// TODO: チャネルを返す
	return c
}