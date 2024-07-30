package main

import "fmt"

func main() {
	fmt.Println("Função Recursiva.")
	posicao := uint(12)
	fmt.Println(fibonacci(posicao))

	for i := uint(1); i < posicao; i++ {
		fmt.Println(fibonacci(i))
	}
}

func fibonacci(posicao uint) uint {
	if posicao <= 1 {
		return posicao
	}
	return fibonacci(posicao-2) + fibonacci(posicao-1)
}
