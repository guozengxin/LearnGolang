package main

import "fmt"

func main() {
	var a = [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	var s = make([]int, 6)
	n1 := copy(s, a[0:])
	n2 := copy(s, s[2:])
	fmt.Printf("%d %d\n", n1, n2)
}
