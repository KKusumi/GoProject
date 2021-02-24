package main

import (
	"fmt"
	"time"
)

func main() {
	//defer fmt.Println("main done")
	//go func() {
	//	defer fmt.Println("goroutine1 done")
	//	time.Sleep(3 * time.Second)
	//}()
	//go func() {
	//	defer fmt.Println("goroutine2 done")
	//	time.Sleep(1 * time.Second)
	//}()
	//time.Sleep(5 * time.Second)

	useCommonValue()
}

func useCommonValue() {
	var v int // 共有の変数
	go func() {
		time.Sleep(1 * time.Second)
		v = 100
	}()
	go func() {
		time.Sleep(1 * time.Second)
		fmt.Println(v) // 処理順序が保証されない。0になることも、100になることもある。
	}()
	time.Sleep(2 * time.Second)
}

/**
 * データ競合の解決方法
 * 1. 一つの変数には1つのゴールーチンからアクセスする構成にする。
 * 2. チャネルを使ってゴールーチン間で通信をする。
 * 3. ロックをとる(syncパッケージ)。
 */

func channelSample() {
	ch1 := make(chan int)
	var ch2 chan string
	go func() { ch1 <- 100 }()
	go func() { ch2 <- "hi"}()

	select {
	case v1 := <-ch1:
	 	fmt.Println(v1)
	case v2 := <-ch2:
		fmt.Println(v2)
	}
}