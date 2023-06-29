package main

import (
	"fmt"
)

func somaAlgarismos(x int) (int, error) {
	if x < 0 {
		return 0, fmt.Errorf("números negativos não são permitidos")
	}

	soma := 0
	for x > 0 {
		digito := x % 10
		soma += digito
		x /= 10
	}

	return soma, nil
}

func main() {
	x := 15

	fmt.Printf("A soma dos algarismos de %d é: ", x)
	resultado, err := somaAlgarismos(x)
	if err != nil {
		fmt.Println("Erro:", err)
		return
	}
	fmt.Println(resultado)
}
