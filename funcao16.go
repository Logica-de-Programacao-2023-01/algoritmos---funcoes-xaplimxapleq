package main

import (
	"fmt"
	"strings"
)

func ocultarVogais(x string) (string, error) {
	if len(x) == 0 {
		return x, fmt.Errorf("seu string está vazio")
	}

	vogais := "AEIOU"
	x = strings.ToUpper(x)
	resultado := ""

	for _, char := range x {
		if strings.ContainsRune(vogais, char) {
			resultado += "*"
		} else {
			resultado += string(char)
		}
	}

	return resultado, nil
}

func main() {
	x := "xaplim xapleq"
	y := ""

	fmt.Print("Sua frase com as vogais ocultas é: ")
	resultado, err := ocultarVogais(x)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(resultado)
	}

	fmt.Print("Sua frase com as vogais ocultas é: ")
	resultado, err = ocultarVogais(y)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(resultado)
	}
}
