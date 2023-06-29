package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	media := MediaAritmetica(nums)
	fmt.Printf("A média aritmética é: %.2f\n", media)
}

func MediaAritmetica(nums []int) float64 {
	if len(nums) == 0 {
		return 0
	}

	soma := 0
	for _, num := range nums {
		soma += num
	}
	return float64(soma) / float64(len(nums))
}
