package main

import "fmt"

func main() {
	var slice = make([]float64, 10)
	slice[0] = 222.0
	fmt.Printf("%f\n", avg(slice))
}

func avg(slice []float64) (a float64) {
	sum := 0.0
	for _, v := range slice {
		sum += v
	}
	a = sum / float64(len(slice))
	return
}
