package main

import "fmt"

func recuperarExecucao() {
	if r := recover(); r != nil {
		fmt.Println("Execução foi recuperada!")
	}
}

func alunoEstaAprovado(nota1, nota2 float32) bool {

	defer recuperarExecucao()
	fmt.Println("entrando na função de media")

	media := (nota1 + nota2) / 2

	if media > 6 {

		return true
	} else if media < 6 {

		return false
	}

	panic("Média é exatamente 6")
}

func main() {
	fmt.Println(alunoEstaAprovado(6, 6))
	fmt.Println("Programa pós execução")

}
