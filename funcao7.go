package main

import "fmt"

func someDez(x int) int {
	return x + 10
}

func aplicarFunc(conj []int, funcao func(int) int) ([]int, error) {
	if len(conj) == 0 {
		return nil, fmt.Errorf("conjunto vazio")
	}

	resultado := make([]int, len(conj))
	for i, num := range conj {
		resultado[i] = funcao(num)
	}

	return resultado, nil
}

func main() {
	conj := []int{2312, 32, 77, 10}
	conjVazio := []int{}

	fmt.Print("Sua função aplicada no conjunto é: ")
	resultado, err := aplicarFunc(conj, someDez)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(resultado)
	}

	fmt.Print("Sua função aplicada no conjunto é: ")
	resultado, err = aplicarFunc(conjVazio, someDez)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(resultado)
	}
}
