package main

import "fmt"

func estaPresente(x int, y []int) (bool, error) {

	estaPresenteSlice := false

	if len(y) == 0 {
		return estaPresenteSlice, fmt.Errorf("seu slice está vazio")
	}
	for _, ranY := range y {
		if ranY == x {
			estaPresenteSlice = true
			break
		} else {
			estaPresenteSlice = false
		}
	}

	return estaPresenteSlice, nil
}

func main() {
	x := 5
	y := []int{5, 56, 4, 5, 11, 17}
	z := []int{4, 2, 12, 74, 84}
	vazio := []int{}

	fmt.Printf("Seu valor %d está presente em %d? ", x, y)
	fmt.Println(estaPresente(x, y))
	fmt.Printf("Seu valor %d está presente em %d? ", x, z)
	fmt.Println(estaPresente(x, z))
	fmt.Printf("Seu valor %d está presente em %d? ", x, vazio)
	fmt.Println(estaPresente(x, vazio))
}
