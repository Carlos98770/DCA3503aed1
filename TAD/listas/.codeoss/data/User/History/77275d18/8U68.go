package main

import "errors"

type Lista interface {
	Get(index int) (int, error)
	Size() int
	Add(elemento int)
	Addt(elemento int, index int) error
	Remove(index int) error
}

type ArrayList struct {
	v    []int
	size int
}

func (l *ArrayList) Size() int {
	return l.size
}

func (l *ArrayList) Get(index int) (int, error) {
	if index > l.size {
		return index, errors.New("Index fora dos limites")
	}

	return l.v[index], nil

}

func (l *ArrayList) Add(elemento int) {
	l.v[l.size] = elemento
	l.size += 1

}