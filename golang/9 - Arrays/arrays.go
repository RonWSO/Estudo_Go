package main

import (
	"fmt"
)

func main() {

	//Arrays
	fmt.Println("Arrays e Slices")

	var array1 [5]string
	array1[0] = "Posição 1"
	fmt.Println(array1)

	array2 := [5]string{"Posição 1", "Posição 2", "Posição 3", "Posição 4", "Posição 5"}

	fmt.Println(array2)

	array3 := [...]int{1, 2, 3, 4, 5}

	fmt.Println(array3)

	//Slices
	//Não define tamanho
	slice := []int{2, 1, 3, 4, 5, 5, -2}
	fmt.Println(slice)

	//fmt.Println(reflect.TypeOf(array3))
	//fmt.Println(reflect.TypeOf(slice))

	slice = append(slice, -3)

	fmt.Println(slice)
	//Como pegar partes do array mas é como se fosse um ponteiro
	slice2 := array2[1:3]
	fmt.Println(slice2)

	array2[1] = "Posição alterada"
	fmt.Println(slice2)

	//Array interno
	slice3 := make([]float32, 10, 11)
	fmt.Println(len(slice3)) //length
	fmt.Println(cap(slice3)) //cap

	slice3 = append(slice3, 1.1)
	slice3 = append(slice3, 1.2)

	fmt.Println(slice3)

	fmt.Println(len(slice3)) //length
	fmt.Println(cap(slice3)) //cap

	slice4 := make([]float32, 5)

	fmt.Println(len(slice4)) //length
	fmt.Println(cap(slice4)) //cap

	slice4 = append(slice4, 1.2)

	fmt.Println(slice4)

	fmt.Println(len(slice4)) //length
	fmt.Println(cap(slice4)) //cap
}
