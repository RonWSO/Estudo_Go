package main

import (
	"fmt"
	"time"
)

func main() {
	//1 - Declara o canal
	canal := make(chan string)
	//2 - Passa o canal e a mensagem para a função escrever
	go escrever("Olá mundo!", canal)

	fmt.Println("Depois da função escrever começar a ser executada")

	for mensagem := range canal {
		fmt.Println(mensagem)
	}
	fmt.Println("Fim do programa")
}

func escrever(texto string, canal chan string) {
	for i := 0; i < 5; i++ {
		//4 - atribui o valor ao canal
		canal <- texto
		time.Sleep(time.Second)
	}
	close(canal)
}
