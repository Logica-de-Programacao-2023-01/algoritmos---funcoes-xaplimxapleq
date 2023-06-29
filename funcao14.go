package main

import (
	"fmt"
)

func intersecao(x []int, y []int) ([]int, error) {
	if len(x) == 0 || len(y) == 0 {
		return nil, fmt.Errorf("pelo menos um dos slices é vazio")
	}

	frequenciaNumX := make(map[int]bool)
	intersecaoSlice := []int{}

	for _, num := range x {
		frequenciaNumX[num] = true
	}

	for _, num := range y {
		if frequenciaNumX[num] {
			intersecaoSlice = append(intersecaoSlice, num)
		}
	}

	return intersecaoSlice, nil
}

func main() {
	x := []int{3423, 7, 8, 4, 5, 6}
	y := []int{3, 7, 5, 23425423432, 6, 2, 1}

	fmt.Printf("A interseção dos conjuntos %v com %v é: ", x, y)
	resultado, err := intersecao(x, y)
	if err != nil {
		fmt.Println("Erro:", err)
		return
	}
	fmt.Println(resultado)
}
