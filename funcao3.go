package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "xaplimxapleq"
	concatenacao := concSlices(str)
	fmt.Println(concatenacao)
}

func concSlices(x string) string {
	lista := strings.Split(x, " ")
	concatenacao := strings.Join(lista, "")
	return concatenacao
}
