package main

import "fmt"

type usuario struct {
	nome  string
	idade uint8
}

func (u usuario) salvar() {
	fmt.Printf("Salvando os dados do usuario %s no banco de dados.\n", u.nome)
}

func (u usuario) maiorDeIdade() bool {
	return u.idade >= 18
}

func (u *usuario) fazerAniversario() {
	u.idade++
}

func main() {
	usuario1 := usuario{"Usuário 1", 20}
	usuario1.salvar()

	usuario2 := usuario{"Davi", 20}
	usuario2.salvar()

	maiorDeidade := usuario2.maiorDeIdade()
	fmt.Println(maiorDeidade)

	usuario2.fazerAniversario()

	fmt.Println(usuario2.idade)
}
