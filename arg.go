package main

func main() {
	myFunc(1, 2, 3, 5)
}

func myFunc(arg ...int) {
	for _, v := range arg {
		println(v)
	}
}
