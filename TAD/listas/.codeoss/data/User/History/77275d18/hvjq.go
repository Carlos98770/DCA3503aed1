package main

import (
	"errors"
	"fmt"
)

type Lista interface {
	Get(index int) (int, error)
	Size() int
	Add(elemento int)
	Add2(elemento int, index int) error
	Remove(index int) error
}

type ArrayList struct {
	v    []int
	insertSize int
}

func (list *ArrayList) Init(size int) error {
	if size > 0 {
		list.v = make([]int, size)
		return nil
	} else {
		return errors.New("size <=0")
	}

}

func (l *ArrayList) Size() int {
	return l.insertSize
}

func (l *ArrayList) Get(index int) (int, error) {
	if index > l.insertSize {
		return index, errors.New("Index fora dos limites")
	}

	return l.v[index], nil

}
func (l *ArrayList) duplicarVetor(){
    newSize := len(l.v)*2
    newV := make([]int, newSize)
    
    for i :=0; i < newSize; i++{
        newV[i] = l.v[i]
    }
    
}

func (l *ArrayList) Add(elemento int) {
    if l.insertSize > len(l.v){
        duplicarVetor()
    }

	l.v[l.insertSize] = elemento
	l.insertSize += 1

}

func main() {
	l := &ArrayList{}
	l.Init(10)
	l.Add(10)
	fmt.Println(l.Size())
}
