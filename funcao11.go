package main

import (
	"fmt"
	"math"
)

func primo(x int) (bool, error) {
	if x < 2 {
		return false, fmt.Errorf("seu número não é primo")
	}

	limite := int(math.Sqrt(float64(x)))
	for i := 2; i <= limite; i++ {
		if x%i == 0 {
			return false, nil
		}
	}

	return true, nil
}

func main() {
	var x int

	fmt.Println("Informe um número:")
	fmt.Scanln(&x)

	primoBool, err := primo(x)
	if err != nil {
		fmt.Println("Erro:", err)
		return
	}

	fmt.Print("Seu número é primo: ", primoBool)
}
