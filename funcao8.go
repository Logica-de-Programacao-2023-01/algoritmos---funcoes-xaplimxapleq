package main

import (
	"fmt"
)

func slicePares(x []int) ([]int, error) {
	pares := []int{}

	for _, num := range x {
		if num%2 == 0 {
			pares = append(pares, num)
		}
	}

	if len(pares) == 0 {
		return nil, fmt.Errorf("conjunto não tem par")
	}

	return pares, nil
}

func main() {
	conj := []int{2312, 32, 77, 10, 3456, 11, 1231, 3334}
	conjImpar := []int{3, 5, 7, 9, 11}

	fmt.Print("Os pares do seu conjunto são: ")
	pares, err := slicePares(conj)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(pares)
	}

	fmt.Print("Os pares do seu conjunto são: ")
	pares, err = slicePares(conjImpar)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(pares)
	}
}
