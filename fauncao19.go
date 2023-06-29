package main

import (
	"fmt"
	"math"
)

func isPrimo(num int) bool {
	if num < 2 {
		return false
	}

	limite := int(math.Sqrt(float64(num)))

	for i := 2; i <= limite; i++ {
		if num%i == 0 {
			return false
		}
	}

	return true
}

func primosAteInt(x int) ([]int, error) {
	if x < 0 {
		return nil, fmt.Errorf("seu número é negativo")
	}

	slicePrimos := []int{}

	for i := 2; i <= x; i++ {
		if isPrimo(i) {
			slicePrimos = append(slicePrimos, i)
		}
	}

	return slicePrimos, nil
}

func main() {
	x := 156
	y := -45

	fmt.Printf("Todos os números primos até %d são: ", x)
	resultado, err := primosAteInt(x)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(resultado)
	}

	fmt.Printf("Todos os números primos até %d são: ", y)
	resultado, err = primosAteInt(y)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(resultado)
	}
}
