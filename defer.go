package main

func main() {
	defer func(i int){
		println(i)
	}(2)
	println(3, 4, 5)
}
