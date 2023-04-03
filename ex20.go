package main

import "fmt"

// Crie um programa que leia um array
//de inteiros e verifique se
//ele está ordenado em ordem crescente.

func main() {
	lista := []int{1, 2, 3, 4, 5, 6}

	crescente := true
	for i := 1; i < len(lista); i++ {
		anterior := lista[i-1]
		atual := lista[i]

		if anterior > atual {
			crescente = false
			break
		}
	}

	if crescente {
		fmt.Println("A lista está em ordem crescente")
	} else {
		fmt.Println("A lista não está em ordem crescente")
	}
}
