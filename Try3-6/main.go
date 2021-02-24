package main

func main() {
	n, m := swap(10, 20)
	println(n, m)
}

func swap(n int, m int) (swappedN, swappedM int) {
	swappedN, swappedM = m, n
	return
}
