package main

import "fmt"

type pessoa struct {
	nome          string
	sobrenomenome string
	idade         uint8
	altura        uint8
}
type estudante struct {
	pessoa
	curso     string
	faculdade string
}

func main() {

	fmt.Println("")

	var p1 pessoa = pessoa{"joão", "Pedro", 20, 180}
	fmt.Println(p1)

	e1 := estudante{pessoa{"Carlos", "Alberto", 22, 178}, "ADM", "UECE"}

	fmt.Println(e1)

	fmt.Println(e1.nome)
}
