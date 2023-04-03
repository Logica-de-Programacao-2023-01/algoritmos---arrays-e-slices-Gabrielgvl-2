package main

import "fmt"

//Crie um Array de inteiros com 10 elementos.
//Crie um novo Slice que contenha apenas
//os elementos pares do Array original.
//Imprima o novo Slice.

func main() {
	array := [10]int{4, 5, 6, 1, 12, 33, 45, 89, 75, 12}
	var resultado []int

	for _, x := range array {
		if x%2 == 0 {
			resultado = append(resultado, x)
		}
	}

	fmt.Println(resultado)
}
