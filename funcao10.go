package main

import (
	"fmt"
)

func sliceCrescente(x []int) ([]int, error) {
	if len(x) == 0 {
		return x, fmt.Errorf("slice vazio")
	}

	for i := 0; i < len(x)-1; i++ {
		for j := 0; j < len(x)-i-1; j++ {
			if x[j] > x[j+1] {
				x[j], x[j+1] = x[j+1], x[j]
			}
		}
	}

	return x, nil
}

func main() {
	conj := []int{34, 3, 434, 23, 3, 22, 2, 0}
	conjVazio := []int{}

	fmt.Print("Seus números em ordem crescente são: ")
	resultado, err := sliceCrescente(conj)
	if err != nil {
		fmt.Println("Erro:", err)
		return
	}
	fmt.Println(resultado)

	fmt.Print("Seus números em ordem crescente são: ")
	resultado, err = sliceCrescente(conjVazio)
	if err != nil {
		fmt.Println("Erro:", err)
		return
	}
	fmt.Println(resultado)
}
