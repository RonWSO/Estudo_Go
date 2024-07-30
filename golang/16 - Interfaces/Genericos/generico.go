package main

import "fmt"

func generica(interf interface{}) {
	fmt.Println(interf)
}

func main() {
	generica("Bla")
	generica(1)
	generica(1.4)
	generica(-2)
	generica(true)
}
