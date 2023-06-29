package main

import (
	"fmt"
	"sort"
)

func ordemAlfabetica(x []string) (string, error) {
	if len(x) == 0 {
		return "", fmt.Errorf("slice vazio")
	}

	sort.Strings(x)
	newString := ""

	for _, str := range x {
		newString += str
	}

	return newString, nil
}

func main() {
	x := []string{"gragas", "homem", "Barril", "Adam", "sandler"}
	vazio := []string{}

	fmt.Print("Sua string em ordem alfabética é: ")
	resultado, err := ordemAlfabetica(x)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(resultado)
	}

	fmt.Print("Sua string em ordem alfabética é: ")
	resultado, err = ordemAlfabetica(vazio)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(resultado)
	}
}
