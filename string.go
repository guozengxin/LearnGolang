package main

import "fmt"

func main() {
	s := "hello"
	c := []rune(s)
	c[0] = 'c'
	s2 := string(c)
	fmt.Printf("%s\n", s2)

	ss := "First" +
		"Second"
	fmt.Printf("%s\n", ss)

	ss1 := `First
		Second`
	fmt.Printf("%s\n", ss1)
}
