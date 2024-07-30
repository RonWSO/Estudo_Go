package main

import "fmt"

func diaDaSemana(numero int8) string {
	switch numero {
	case 1:
		return "Domingo"
	case 2:
		return "Segunda-Feira"
	case 3:
		return "Terça-Feira"
	case 4:
		return "Quarta-Feira"
	case 5:
		return "Quinta-Feira"
	case 6:
		return "Sexta-Feira"
	case 7:
		return "Sábado"
	default:
		return "Número inválido"
	}

}

func diaDaSemana2(numero int8) string {

	var Dia string
	switch {
	case numero == 1:
		Dia = "Domingo"
		fallthrough
	case numero == 2:
		Dia = "Segunda-Feira"
	case numero == 3:
		Dia = "Terça-Feira"
	case numero == 4:
		Dia = "Quarta-Feira"
	case numero == 5:
		Dia = "Quinta-Feira"
	case numero == 6:
		Dia = "Sexta-Feira"
	case numero == 7:
		Dia = "Sábado"
	default:
		Dia = "Número inválido"
	}
	return Dia
}
func main() {
	fmt.Println("Switch")

	var dia string = diaDaSemana(3)

	fmt.Println(dia)

	var dia2 string = diaDaSemana2(1)

	fmt.Println(dia2)
}
