package main

import "fmt"

func main() {
	//Canal com buffer
	canal := make(chan string, 2)
	canal <- "Olá Mundo"

	mensagem := <-canal
	fmt.Println(mensagem)
}
