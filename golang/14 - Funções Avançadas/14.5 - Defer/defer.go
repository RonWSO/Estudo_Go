package main

import "fmt"

func funcao1() {
	fmt.Println("Executando a função 1.")
}

func funcao2() {
	fmt.Println("Executando a função 2.")
}

func alunoEstaAprovado(nota1, nota2 float32) bool {

	defer fmt.Println("Média calculada resultado será retornado!")
	fmt.Println("entrando na função de media")

	media := (nota1 + nota2) / 2

	if media > 6 {

		return true
	}

	return false
}

func main() {
	//defer funcao1()
	//funcao2()

	fmt.Println(alunoEstaAprovado(6, 7))
}
