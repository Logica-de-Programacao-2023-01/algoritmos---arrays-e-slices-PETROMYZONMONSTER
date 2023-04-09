package main

import "fmt"

func main() {
	var array = [3]int{1, 2, 3}
	fmt.Println("Valores no array:", array)
	resultado := 0
	for _, numeros := range array {
		resultado += numeros
	}
	fmt.Println("A soma dos valores é:", resultado)
}
