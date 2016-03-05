package main

func main() {
	var a [10] int
	slice := a[0:5]
	println(len(a))
	println(cap(a))
	println(len(slice))
	println(cap(slice))
}
