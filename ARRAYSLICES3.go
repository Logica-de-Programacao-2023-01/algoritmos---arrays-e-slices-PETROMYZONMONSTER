package main

import "fmt"

func main() {
	var array = [4]float64{1, 2, 3, 4}
	var resultado float64
	resultado = 1
	for _, numeros := range array {
		resultado *= numeros
	}
	fmt.Println(resultado)
}
