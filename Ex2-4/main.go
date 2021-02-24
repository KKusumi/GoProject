package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	var number = rand.Intn(6) + 1
	switch {
	case number == 1:
		fmt.Println("大吉")
	case number == 5 || number == 4:
		fmt.Println("中吉")
	case number == 3 || number == 2:
		fmt.Println("吉")
	case number == 1:
		fmt.Println("凶")
	}
}
