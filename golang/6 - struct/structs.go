package main

import "fmt"

type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}

type endereco struct {
	logradouro string
	numero     uint8
}

func main() {
	fmt.Println("Struct")

	var u usuario
	u.nome = "Ronald"
	u.idade = 21

	u2 := usuario{"Ron", 22, endereco{"Rua dos bobos", 0}}

	fmt.Println(u, u2)

	//Para instanciar uma struct sem ter todas as informações
	end := endereco{"Rua dos bobos", 2}
	u3 := usuario{idade: 24, endereco: end}

	fmt.Println(u3)
}
