package main

import (
	"fmt"
)

const pi float32 = 3.14159265358979323846

func main() {

	var nome string = "Carlos" //Define a variável de um certo tipo

	nome2 := "Eduardo" //Aqui o tipo é inferido //Can only be used inside functions

	var numero int //Can be used inside and outside of functions
	numero = 5

	fmt.Println(nome + " " + nome2)
	fmt.Println(numero)
	fmt.Println(pi)
}
