package main

func main() {
	callback(3, printit)
}

func printit(x int) {
	println(x)
}

func callback(y int, f func(int)) {
	f(y)
}
