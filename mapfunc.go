package main

func main() {
	var xs = map[int]func() int{
		1: func() int {return 10},
		2: func() int { return 20 },
		3: func() int { return 30 },
	}
	print(xs[2]())
}
