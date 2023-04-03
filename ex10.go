package main

import "fmt"

//Crie um Slice de inteiros com tamanho 10
//e imprima o valor mínimo e o valor máximo
//armazenados no Slice.

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	min := 999999999999
	max := -99999999999

	for _, x := range slice {
		if x > max {
			max = x
		}
		if x < min {
			min = x
		}
	}

	fmt.Println("O maior valor é", max)
	fmt.Println("O menor valor é", min)
}
