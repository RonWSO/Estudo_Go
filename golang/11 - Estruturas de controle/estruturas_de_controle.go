package main

import "fmt"

func main() {
	fmt.Println("Estruturas de controle")

	numero := 10

	if numero > 15 {
		fmt.Println("Numero maior que 15")
	} else {
		fmt.Println("Numero menor ou igual a 15")
	}

	//If init ela só existe dentro do if mesmo não sendo declarado antes
	if outroNumero := numero; outroNumero > 0 {
		fmt.Println("Numero é maior que zero")
	} else {
		fmt.Println("Numero é menor ou igual a zero")
	}

}
