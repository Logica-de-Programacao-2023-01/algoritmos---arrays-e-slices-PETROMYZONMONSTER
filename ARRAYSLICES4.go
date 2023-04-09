// Crie um Slice de inteiros e solicite ao usuário que informe o tamanho do Slice e os valores dos elementos. Em seguida, imprima o Slice e a soma dos valores armazenados.

package main

import "fmt"

func main() {
	slice := []int{}
	var n, x, soma int
	fmt.Println("Digite o tamanho da slice.")
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Println("Digite os elementos da slice.")
		fmt.Scan(&x)
		slice = append(slice, x)
		soma += slice[i]
	}
	fmt.Println(slice, soma)
}
