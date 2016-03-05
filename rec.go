package main

func rec(i int) {
	if i == 10 {
		return
	}
	rec(i + 1)
	print(i, " ")
}

func main() {
	rec(1)
	println()
}
