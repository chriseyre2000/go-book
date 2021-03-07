package main

import "fmt"

// This is a comment

func average(xs []float64) float64 {
	total := 0.0
	for _, v := range(xs) {
		total += v
	}
	return total / float64(len(xs))
}

func f() (int, int) {
	return 5, 6
}

func add(args ...int) int {
	total := 0;
	for _, v := range(args) {
		total += v
	}
	return total
}

func makeEvenGenerator() func() uint {
	i := uint(0)
	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
}

func factorial(x uint) uint {
	if x == 0 {
		return 1
	}
	return x * factorial(x-1)
}

func main() {


	fmt.Println(factorial(100))

}