package main

import "fmt"

func soma(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}

	return total
}

//Só permite um paramtro variatico por função e tem que ser o ultimo parametro da função
func escrever(texto string, numeros ...int) {
	for _, numero := range numeros {
		fmt.Println(texto, numero)
	}
}
func main() {
	printTotal := soma(1, 2, 3, 4, 5, 6, 7, 8)

	fmt.Println(printTotal)

	escrever("Olá mundo", 1, 2, 3, 4, 5, 6)
}
