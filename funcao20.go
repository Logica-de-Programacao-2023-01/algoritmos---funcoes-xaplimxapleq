package main

import (
	"fmt"
)

func apenasMaioresQueCinco(x []string) ([]string, error) {
	maioresQueCinco := []string{}

	if len(x) == 0 {
		return maioresQueCinco, fmt.Errorf("seu slice está vazio")
	}

	for _, str := range x {
		if len(str) > 5 {
			maioresQueCinco = append(maioresQueCinco, str)
		}
	}

	return maioresQueCinco, nil
}

func main() {
	x := []string{"Xaplim", "Xapleq", "Xaploq", "Lucas", "Filipe", "Pelado"}
	y := []string{}

	fmt.Printf("Seus strings maiores que 5 em %v são: ", x)
	resultado, err := apenasMaioresQueCinco(x)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(resultado)
	}

	fmt.Printf("Seus strings maiores que 5 em %v são: ", y)
	resultado, err = apenasMaioresQueCinco(y)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(resultado)
	}
}
