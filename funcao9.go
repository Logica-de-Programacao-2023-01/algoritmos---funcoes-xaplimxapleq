package main

import (
	"fmt"
	"strings"
)

func palavrasContidasNoString(x string) ([]string, error) {
	conjPalavras := strings.Fields(x)

	if len(conjPalavras) == 0 {
		return nil, fmt.Errorf("slice vazio")
	}

	return conjPalavras, nil
}

func main() {
	var input string
	fmt.Println("Informe uma string:")
	_, err := fmt.Scanln(&input)
	if err != nil {
		fmt.Println("Erro ao ler a entrada:", err)
		return
	}

	palavras, err := palavrasContidasNoString(input)
	if err != nil {
		fmt.Println("Erro:", err)
		return
	}

	fmt.Println("Suas palavras são:", palavras)
}
