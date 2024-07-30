package main

import "fmt"

func main() {
	//Declara channels com 45 de buffer
	tarefas := make(chan int, 45)
	resultados := make(chan int, 45)

	//Quantos mais workers mais o processo passa a ser trabalhado, agilizando o processo
	go worker(tarefas, resultados)
	go worker(tarefas, resultados)
	go worker(tarefas, resultados)
	go worker(tarefas, resultados)

	for i := 0; i < 45; i++ {
		tarefas <- i
	}
	close(tarefas)
	for i := 0; i < 45; i++ {
		resultado := <-resultados
		fmt.Println(resultado)
	}
	close(resultados)
}

// Declara que o canal tarefas apenas recebe, e o canal resultado apenas escreve, dentro da função
func worker(tarefas <-chan int, resultados chan<- int) {

	for numero := range tarefas {
		resultados <- fibonacci(numero)
	}
}

func fibonacci(posicao int) int {
	if posicao <= 1 {
		return posicao
	}
	return fibonacci(posicao-2) + fibonacci(posicao-1)
}
