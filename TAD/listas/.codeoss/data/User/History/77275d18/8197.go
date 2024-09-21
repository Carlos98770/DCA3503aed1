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

func (l *ArrayList) Add(elemento int) {
	l.v[l.insertSize] = elemento
	l.insertSize += 1

}

func (l *ArrayList) Add2(elemento int, index int) error {

}

func main() {
	l := &ArrayList{}
	l.Init(10)
	l.Add(10)
	fmt.Println(l.Size())
}
