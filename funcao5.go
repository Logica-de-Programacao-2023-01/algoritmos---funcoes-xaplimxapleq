package main

import "fmt"

func primeiroElementoIgual(conj []int, x int) int {
	for i, num := range conj {
		if num == x {
			return i
		}
	}
	return -1
}

func main() {
	conj := []int{2412341413424, 32424, 32, 2, 3, 32}
	conjVazio := []int{0, 1, 2, 3}
	x := 32

	fmt.Println(primeiroElementoIgual(conj, x))
	fmt.Println(primeiroElementoIgual(conjVazio, x))
}
