package arraylist

import (
	"errors"
	"myproject/lista" // Substitua pelo caminho correto se necessário
)

type ArrayList struct {
	v          []int
	insertSize int
}

// Inicializa o ArrayList com um tamanho inicial
func (list *ArrayList) Init(size int) error {
	if size > 0 {
		list.v = make([]int, size)
		return nil
	}
	return errors.New("size <= 0")
}

// Retorna o tamanho do ArrayList
func (l *ArrayList) Size() int {
	return l.insertSize
}

// Retorna o valor no índice especificado
func (l *ArrayList) Get(index int) (int, error) {
	if index < 0 || index >= l.insertSize {
		return 0, errors.New("index fora dos limites")
	}
	return l.v[index], nil
}

// Duplica o tamanho do array quando atinge a capacidade máxima
func (l *ArrayList) duplicarVetor() {
	newSize := len(l.v) * 2
	newV := make([]int, newSize)
	copy(newV, l.v)
	l.v = newV
}

// Adiciona um elemento ao ArrayList
func (l *ArrayList) Add(elemento int) {
	if l.insertSize == len(l.v) {
		l.duplicarVetor()
	}
	l.v[l.insertSize] = elemento
	l.insertSize++
}

// Implemente Add2 e Remove conforme necessário
