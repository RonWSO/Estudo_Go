package main

import "fmt"

//define os paramtros e o retorno no começo da função
func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

func calculos(n1, n2 int8) (int8, int8) {
	soma := n1 + n2
	subtracao := n1 - n2
	return soma, subtracao
}

func main() {
	soma := somar(10, 20)
	fmt.Println(soma)
	//funções são tipos, é possível colocar uma função em uma variavel
	var f = func(txt string) {
		fmt.Println(txt)
	}

	f("Função F")
	//go permite retorno de mais de um valor por função
	resultadoSoma, resultadoSubtracao := calculos(20, 10)
	fmt.Println(resultadoSoma, resultadoSubtracao)
	//se eu não quiser utilizar o valor do outro retorno coloca um _
	resultadoSoma2, _ := calculos(30, 10)
	fmt.Println(resultadoSoma2)
}
