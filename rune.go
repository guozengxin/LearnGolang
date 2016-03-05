package main

import "fmt"
import "unicode/utf8"

func main() {
	str := "adsfasdf......js"
	fmt.Printf("String %s\nLength: %d, Runes: %d\n", str,
			len([]byte(str)), utf8.RuneCount([]byte(str)))

	s := "                "
	r := []rune(s)
	copy(r[4:4+3], []rune("abc")) 
	fmt.Printf("Before: %s\n", s); 
	fmt.Printf("After : %s\n", string(r))
}
