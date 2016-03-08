package main

import "fmt"

func main() {
	var p *int
	var i int
	p = &i
	fmt.Println("%v\n", p)
}
