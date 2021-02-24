package main

func main() {
	n, m := 10, 20
	swap(&n, &m)
	println(n, m)
}

func swap(nP *int, mP *int) {
	var p = nP
	*nP = *mP
	*mP = *p
}
