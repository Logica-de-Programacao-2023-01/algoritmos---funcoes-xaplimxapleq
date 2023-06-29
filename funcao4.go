package main

import "fmt"

func segundoMaior(x []int) int {
	idxMaior := -1
	maior := x[0]
	segMaior := x[0]

	for _, ranX := range x {
		if ranX > maior {
			maior = ranX
		}
	}

	for i := 0; i < len(x); i++ {
		if x[i] == maior {
			idxMaior = i
			x = append(x[:idxMaior], x[idxMaior+1:]...)
			i = 0
		}
	}

	for _, ranX := range x {
		if ranX > segMaior {
			segMaior = ranX
		}
	}

	return segMaior

}

func main() {
	conj := []int{234234, 43434, 434, 0, 3, 234234, 999999999, 999999999}

	fmt.Println("O segundo maior do seu conjunto é: ", segundoMaior(conj))
}
