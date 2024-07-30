package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	canal := multiplexar(escrever("Olá Mundo!"), escrever("Programar em Go!"))

	for i := 0; i < 10; i++ {
		fmt.Println(<-canal)
	}
}

func multiplexar(canal1, canal2 <-chan string) <-chan string {
	canalDeSaide := make(chan string)

	go func() {
		for {
			select {
			case mensagem := <-canal1:
				canalDeSaide <- mensagem
			case mensagem := <-canal2:
				canalDeSaide <- mensagem
			}

		}
	}()

	return canalDeSaide
}

func escrever(texto string) <-chan string {
	canal := make(chan string)

	go func() {
		for {
			canal <- fmt.Sprintf("valor Recebido: %s", texto)
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(2000)))
		}
	}()

	return canal
}
