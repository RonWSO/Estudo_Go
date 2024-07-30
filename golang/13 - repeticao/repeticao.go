package main

import (
	"fmt"
	//"time"
)

func main() {
	//fmt.Println("Repetição")
	//i := 0
	//for i < 10 {
	//	i++
	//	fmt.Println("Incrementando 1")
	//	time.Sleep(time.Second)
	//}
	//fmt.Println(i)
	//for j := 0; j < 10; j++ {
	//	fmt.Println("Incrementando J", j)
	//	time.Sleep(time.Second)
	//}

	nomes := [3]string{"João", "Davi", "Ronald"}

	for indice, nome := range nomes {
		fmt.Println(indice, nome)
	}
	for _, nome := range nomes {
		fmt.Println(nome)
	}

	for _, letra := range "Palavra" {
		fmt.Println(string(letra))
	}

	usuario := map[string]string{
		"Nome":      "Ronald",
		"Sobrenome": "Oliveira",
	}
	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}
}
