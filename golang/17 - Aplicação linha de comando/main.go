package main

import (
	"fmt"
	"linha-de-comando/app"
	"log"
	"os"
)

func main() {
	fmt.Println("Ponto de partida")
	//instancia a aplicação
	aplicacao := app.Gerar()
	//Checa se a aplicação esta em erro e termina a aplicação mostrando a razão do erro.
	if erro := aplicacao.Run(os.Args); erro != nil {
		log.Fatal(erro)
	}

}
