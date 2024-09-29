package main

import (
	"fmt"
	"modPlat/implementacoes"
)

func main() {
	/*
		l := &implementacoes.ArrayList{}
		l.Init(10)
		l.Add(1)
		l.Add(2)
		l.Add(3)
		l.Add(4)
		l.Add(5)
		l.Add(6)
		l.Add(7)
		l.Add(8)
		l.Add(9)
		l.Add(10)

		l.Remove(9)
		l.Remove(8)
		l.PrintVetor()
	*/
	listaLigada := &implementacoes.LinkedList{}
	listaLigada.Add(200)
	listaLigada.Add(7)
	listaLigada.Add(10)
	listaLigada.Add2(0,2)
	//fmt.Println(listaLigada.Get(1))

	for i:=0; i < listaLigada.Size();i++{
		fmt.Println(listaLigada.Get(i))
	}

	//listaLigada.Remove(0)
	//fmt.Println(listaLigada.Size())
	//fmt.Println(listaLigada.Get(0))
	//fmt.Println(listaLigada.Get(1))
}
