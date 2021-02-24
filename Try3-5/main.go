package main

func main() {
	for i := 1; i <= 100; i++ {
		print(i)
		if isEven(i) {
			println("-偶数")
		} else {
			println("-奇数")
		}
	}
}

func isEven(n int) bool {
	return n%2 == 0
}
