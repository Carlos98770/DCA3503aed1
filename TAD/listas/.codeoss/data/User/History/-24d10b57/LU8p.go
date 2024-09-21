package lista

import (
	"errors"
)

type Lista interface {
	Get(index int) (int, error)
	Size() int
	Add(elemento int)
	Add2(elemento int, index int) error
	Remove(index int) error
}



/*
func main() {
	l := &ArrayList{}
	l.Init(2)
	l.Add(10)
	l.Add(6)
	l.Add(4)
	fmt.Println(l.Size())
}
*/
