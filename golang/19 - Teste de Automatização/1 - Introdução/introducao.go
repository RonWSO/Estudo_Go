package main

import (
	"fmt"
	"introducao-testes/enderecos"
)

func main() {
	tipoEndereco := enderecos.TipoDeEnderecos("Rua Félix Gomes")
	fmt.Println(tipoEndereco)
}
