package main

import (
	"fmt"
	"strconv"
)

type stack struct {
	i int
	data [10] int
}

func (s *stack) push(k int) {
	s.data[s.i] = k
	s.i ++
}

func (s *stack) pop() int {
	s.i --
	return s.data[s.i]
}

func (s stack) String() string {
	var str string
	for i := 0; i < s.i; i ++ {
		str = str + "[" + strconv.Itoa(i) + ":" + strconv.Itoa(s.data[i]) + "]"
	}
	return str
}

func main() {
	s := new(stack)
	s.push(25)
	s.push(2)
	s.push(384)
	fmt.Printf("%d\n", s.pop())
	fmt.Printf("%v\n", s)
}
