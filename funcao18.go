package main

import (
	"fmt"
)

func aplicarFuncEmIntSlice(x []int, y func(int) (int, error)) (int, error) {
	if len(x) == 0 {
		return 0, fmt.Errorf("seu slice está vazio")
	}

	soma := 0

	for _, valor := range x {
		resultado, err := y(valor)
		if err != nil {
			return 0, err
		}
		soma += resultado
	}

	return soma, nil
}

func exemploFuncao(valor int) (int, error) {
	// Exemplo de implementação de uma função que pode ser passada para aplicarFuncEmIntSlice
	// Aqui, a função simplesmente retorna o valor multiplicado por 2.
	return valor * 2, nil
}

func main() {
	x := []int{1, 2, 3, 4, 5}

	fmt.Print("Soma dos resultados da função aplicada em cada elemento: ")
	resultado, err := aplicarFuncEmIntSlice(x, exemploFuncao)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(resultado)
	}
}
