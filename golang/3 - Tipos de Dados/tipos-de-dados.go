package main

import (
	"errors"
	"fmt"
)

func main() {
	var numero int16 = 100
	var numero1 int64 = 1000000000000000000
	var numero2 int = 00000000000000
	fmt.Println(numero)
	fmt.Println(numero1)
	fmt.Println(numero2)

	var numero3 uint = 1000
	fmt.Println(numero3)
	//igual a int32
	var numero4 rune = 12345
	fmt.Println(numero4)

	//igual a uint8
	var valor1 byte = 123
	fmt.Println(valor1)

	//floats devem ser definidos como 32 ou 64
	var numeroReal1 float32 = 1234.45
	fmt.Println(numeroReal1)

	var numeroReal2 float64 = 1234.45
	fmt.Println(numeroReal2)

	///Começo string
	var str string = "Texto"
	fmt.Println(str)

	//não existe char como e o classico

	char := 'B'
	fmt.Println(char)

	//Fim Strings

	var booleano bool
	fmt.Println(booleano)
	//Golang se usa o tipo error
	var erro error = errors.New("Erro interno")
	fmt.Println(erro)
}
