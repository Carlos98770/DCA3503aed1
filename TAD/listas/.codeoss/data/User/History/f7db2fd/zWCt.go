package implementacoes

import (
	"errors"
	"modPlat/interfaces"
	"fmt"
)

type ArrayList struct {
	v          []int
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
	fmt.Println(cap(l.v))
	return l.insertSize
}

func (l *ArrayList) Get(index int) (int, error) {
	if index > l.insertSize {
		return index, errors.New("Index fora dos limites")
	}

	return l.v[index], nil

}
func (l *ArrayList) duplicarVetor() {
	newSize := len(l.v) * 2
	newV := make([]int, newSize)

	for i := 0; i < l.insertSize; i++ {
		newV[i] = l.v[i]

	}

	l.v = newV

}

func (l *ArrayList) Add(elemento int) {
	if l.insertSize == len(l.v) {
		l.duplicarVetor()

	}

	l.v[l.insertSize] = elemento
	l.insertSize += 1

}

func (l *ArrayList) Remove(index int) error {

	//verificar se o index é valido
	//coloca o elemento na ultima posicao e deslocar oq estava a direita para esquerda e remover uma unidade no insertSize
	//caso o index seja == l.insertSize, apenas fazer a remoção de uma unidade no insertSize
	return nil
}

func (l *ArrayList) Add2(elemento int, index int) error {
	//verificar se o index é valido
	//faz os elemento a direito do index serem deslocados em um unidade, consequentemente aumento em uma unidade o insertSize
	////caso o index seja == l.insertSize, apenas fazer a adição de uma unidade no insertSize
	if index < 0 || index > l.insertSize-1 {
		return errors.New("Index invalido")
	} else if l.insertSize == len(l.v) {
		l.duplicarVetor()
	}
	l.Add(0)
	for i := l.insertSize; i > index; i-- {
		l.v[i] = l.v[i-1]
	}

	l.v[index] = elemento
	l.insertSize += 1

	return nil
}
