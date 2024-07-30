package main

import (
	"fmt"
	"math"
)

type retangulo struct {
	altura  float64
	largura float64
}

func (r retangulo) area() float64 {
	return r.altura * r.largura
}

type circulo struct {
	raio float64
}

func (c circulo) area() float64 {
	return math.Pi * (c.raio * c.raio)
}

// A interface serve para fazer uma função generica entre duas estruturas diferentes, uma struct só pode usar uma interface se tiver a função generica
type forma interface {
	area() float64
}

func escreverArea(f forma) {
	fmt.Printf("A Area da forma é %0.2f\n", f.area())
}

func main() {
	r := retangulo{10, 15}
	escreverArea(r)

	c := circulo{10}
	escreverArea(c)
}
