package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Printf("%d\n", sum)
	println(sum)

	list := []string{"a", "c", "e"}
	for k, v := range list {
		println(k, v)
	}
}
