package main

func main() {
	println(checkPanic(f1))
	println(checkPanic(f2))
}

func checkPanic(f func()) (b bool) {
	defer func() {
		if x := recover(); x != nil {
			b = true
		}
	}()
	f()
	return
}

func f1() {
	panic(1)
}

func f2() {
}
