package implementacoes_listas

import (
	"errors"
)

type ArrayList struct {
	v          []int
	insertSize int
}

func (l *ArrayList) Init(size int) error {
	if size > 0 {
		l.v = make([]int, size)
		l.insertSize = 0
		return nil
	} else {
		return errors.New("size <=0")
	}

}

func DuplicarVetor(l *ArrayList) {
	aux := make([]int, len(l.v)*2)
	for i := 0; i < l.insertSize; i++ {
		aux[i] = l.v[i]
	}
	l.v = aux
}

func (l *ArrayList) Add(value int) {
	if l.insertSize == len(l.v) {
		DuplicarVetor(l)
	}

	l.v[l.insertSize] = value
	l.insertSize++
}

func (l *ArrayList) AddOnIndex(value int, index int) error {
	if index >= 0 && index < l.insertSize {
		if l.insertSize == len(l.v) {
			DuplicarVetor(l)
		} else {
			for i := l.insertSize; i > index; i-- {
				l.v[i] = l.v[i-1]
			}

			l.v[index] = value
			l.insertSize++
			return nil
		}
	}
	return errors.New("INDEX INVALIDO, IMPOSSIVEL ADICIONAR AO VETOR")
}

func (l *ArrayList) RemoveOnIndex(index int) error {
	if index >= 0 && index < l.insertSize {

		if l.insertSize == 0 {
			return errors.New("Lista vazia")
		} else {
			for i := index; i < l.insertSize-1; i++ {
				l.v[i] = l.v[i+1]
			}

			l.insertSize--
			return nil
		}

	}
	return errors.New("INDEX INVALIDO")
}

func (l *ArrayList) Set(value int, index int) error {
	if index >= 0 && index < l.insertSize {
		l.v[index] = value
		return nil
	}
	return errors.New("INDEX INVALIDO")
}

func (l *ArrayList) Size() int {
	return l.insertSize
}

func (l *ArrayList) Get(index int) (error, int) {
	if index >= 0 && index < l.insertSize {
		return nil, l.v[index]
	}
	return errors.New("INDEX INVALIDO"), -1
}

func (l *ArrayList) Reverse() {
	final := l.insertSize - 1
	for i := 0; final > i; i++ {
		aux := l.v[i]
		l.v[i] = l.v[final]
		l.v[final] = aux
		final -= 1
	}
	
}
