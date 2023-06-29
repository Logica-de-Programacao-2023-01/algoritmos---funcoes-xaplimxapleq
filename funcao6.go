package main

import (
	"fmt"
	"strings"
)

func concatenacaoComVirgula(x []string) (string, error) {
	if len(x) == 0 {
		return "", fmt.Errorf("slice vazio")
	}
	return strings.Join(x, ","), nil
}

func main() {
	frases := []string{"fafsafa", "ebaaaaaaaa", "kkkkkk", "putz"}
	conjVazio := []string{}

	fmt.Print("Sua string concatenada com vírgulas é: ")
	resultado, err := concatenacaoComVirgula(frases)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(resultado)
	}

	fmt.Print("Sua string concatenada com vírgulas é: ")
	resultado, err = concatenacaoComVirgula(conjVazio)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(resultado)
	}
}
