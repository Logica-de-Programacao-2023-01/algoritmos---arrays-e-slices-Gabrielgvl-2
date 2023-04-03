package main

import "fmt"

// Crie um Slice de strings com
//tamanho 8 e solicite ao
//usuário que informe um valor.
//Remova todas as ocorrências desse valor
//no Slice e imprima o resultado.

func main() {
	slice := []string{"a", "a", "b", "c", "d", "a", "ab", "c"}
	var valor string
	fmt.Print("Digite um valor: ")
	fmt.Scan(&valor)

	var resultado []string

	for _, x := range slice {
		if x != valor {
			resultado = append(resultado, x)
		}
	}

	fmt.Println(resultado)
}
