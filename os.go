package main

import (
	"os"
	"bufio"
)

func main() {
	f, _ := os.Open("/etc/passwd")
	defer f.Close()
	r := bufio.NewReader(f)
	s, _ := r.ReadString('\n')
	println(s)
}
