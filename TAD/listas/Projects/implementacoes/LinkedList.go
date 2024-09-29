package implementacoes

import (
	"errors"
)

type Node struct {
	elemento int
	nextNode *Node
}

type LinkedList struct {
	head       *Node
	insertSize int
}

func (l *LinkedList) Get(index int) (int, error) {
	if index >= 0 && index < l.insertSize {

		aux := l.head
		for i := 0; i < index; i++ {
			aux = aux.nextNode
		}

		return aux.elemento, nil
	}

	return -1, errors.New("Index inválido")
}

func (l *LinkedList) Size() int {
	return l.insertSize
}
func (l *LinkedList) Add(elemento int) {
	//Função para adicionar elemento no final

	//Instacianciamos o Nó que será incluido
	newNode := &Node{elemento: elemento}
	if l.head == nil {
		//verificar se a cabeça é nula
		l.head = newNode
	} else {
		//interamos e vereficamos se o endereço é null, caso seja será o ultimo elemento
		//ou interamos até a ultima posição e fazemos a "Ultima posicao" apontar para o novo nó

		aux := l.head
		for i := 0; i < l.insertSize-1; i++ {
			aux = aux.nextNode
		}

		aux.nextNode = newNode
	}

	l.insertSize++
}

func (l *LinkedList) Remove(index int) error {
	if index >= 0 && index < l.insertSize {

		if index == 0 {
			// Remover o primeiro nó (caso especial)
			l.head = l.head.nextNode
		} else {

			aux := l.head
			//Para no nó anterior ao que queiramos remover
			for i := 0; i < index-1; i++ {
				aux = aux.nextNode

			}

			 // "Pular" o nó que será removido
			aux.nextNode = aux.nextNode.nextNode
			return nil

		}
		l.insertSize--
	}
	return errors.New("Index inválido")
}

func (l *LinkedList) Add2(elemento int, index int) error{
	if index >=0 && index < l.insertSize{
		newNode := &Node{elemento:elemento}
		if index == 0 {
			//Caso a adição seja na cabeça
			newNode.nextNode = l.head.nextNode
			l.head = newNode
		} else {
			
			aux := l.head
			for i:= 0; i < index - 1;i++{
				aux = aux.nextNode
			}

			//Pecorremos antes o index - 1, fazendo o newNode apontar para o index e fazendo o index -1 apontar para o newNodes
			newNode.nextNode = aux.nextNode
			aux.nextNode = newNode

		}

		l.insertSize ++
		return nil
	
	}

	return errors.New("Index inválido")
}


