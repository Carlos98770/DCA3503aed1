package main

import (
	"fmt"
	"modPlat/implementacoes"
)

func main() {
	l := &implementacoes.ArrayList{}
	l.Init(2)
	l.Add(10)
	l.Add(6)
	l.Add(4)
	fmt.Println(l.Add2(59, 6))
	l.Remove(0)
	l.PrintVetor()

}
