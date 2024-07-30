package main

import "fmt"

func main() {
	fmt.Println(" Maps ")

	usuario := map[string]string{
		"Nome":      "Pedro",
		"Sobrenome": "Oliveira",
	}

	fmt.Println(usuario)
	fmt.Println(usuario["Nome"])

	usuario2 := map[string]map[string]string{
		"nome": {
			"primeiro": "Afonso",
			"ultimo":   "Barros",
		},
	}
	fmt.Println(usuario2)
	fmt.Println(usuario2["nome"])
	fmt.Println(usuario2["nome"]["primeiro"])

	delete(usuario2, "nome")
	fmt.Println(usuario2)

	usuario2["curso"] = map[string]string{
		"nome":   "Engenharia",
		"campus": "Barra",
	}
	fmt.Println(usuario2)
}
