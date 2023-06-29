package main

import (
	"fmt"
	"strings"
	"unicode"
)

func comecoMaiusculo(x []string) (string, error) {
	var palavras []string

	if len(x) == 0 {
		return "", fmt.Errorf("slice vazio")
	}

	for _, palavra := range x {
		if len(palavra) > 0 && unicode.IsUpper(rune(palavra[0])) {
			palavras = append(palavras, palavra)
		}
	}

	return strings.Join(palavras, " "), nil
}

func main() {
	palavras := []string{"Xaplim", "feijao", "Doido", "diabo", "Boing", "Felipeneto"}

	fmt.Print("As palavras que começam com letra maiúscula são: ")
	resultado, err := comecoMaiusculo(palavras)
	if err != nil {
		fmt.Println("Erro:", err)
		return
	}
	fmt.Println(resultado)
}
