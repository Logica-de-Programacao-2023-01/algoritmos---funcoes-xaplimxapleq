package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "xaplimxapleq"
	quantidade := QuantasVogais(str)
	fmt.Printf("A quantidade de vogais na string é: %d\n", quantidade)
}

func QuantasVogais(x string) int {
	quantidadeVogais := 0

	for _, letra := range x {
		if isVogal(letra) {
			quantidadeVogais++
		}
	}

	return quantidadeVogais
}

func isVogal(letra rune) bool {
	vogais := "aeiouAEIOU"
	return strings.ContainsRune(vogais, letra)
}
