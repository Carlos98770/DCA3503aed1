package implementacoes

import (
	"errors"
)

type DllNode struct {
	prev     *DllNode
	elemento int
	next     *DllNode
}

type DoublyLinkedList struct {
	head       *DllNode
	tail       *DllNode
	insertSize int
}

func (l *DoublyLinkedList) Add(elemento int) {
	//Criar novo nó
	newNode := &DllNode{elemento: elemento}

	if l.head == nil {
		l.head = newNode
		l.tail = newNode
	} else {
		//a cabeça não é nula, a calda aponta para o novo nó e o prev do novo nó aponta para o tail
		l.tail.next = newNode
		newNode.prev = l.tail
	}

	l.tail = newNode
	l.insertSize++

}
func (l *DoublyLinkedList) Size() int {
	return l.insertSize
}

func (l *DoublyLinkedList) Get(index int) (int, error) {
	if index >= 0 && index < l.insertSize {

		aux := l.head
		for i := 0; i < index; i++ {
			aux = aux.next
		}

		return aux.elemento, nil
	}

	return -1, errors.New("Index inválido")
}

func (l *DoublyLinkedList) Remove(index int) error {
	if index >= 0 && index < l.insertSize {
		if index == 0 {
			l.head = l.head.next
		} else {

			aux := l.head
			for i := 0; i < index; i++ {
				aux = aux.next
			}

			if aux.next != nil {
				//elemento não é a cauda
				aux.prev.next = aux.next
				aux.next.prev = aux.prev

			} else {
				//elemento é a cauda, ou seja é o index == l.insertSize -1
				aux.next = nil
				l.tail = aux
			}
		}

		l.insertSize -= 1

		return nil
	}
	return errors.New("Index inválido")
}
func (l *DoublyLinkedList) Add2(elemento int, index int) error {
	newNode := &DllNode{elemento: elemento}
	if index >= 0 && index <= l.insertSize {
		if l.head == nil {
			l.head = newNode
			l.tail = newNode
		} else {

			if index == 0 {
				//Adiciona na primeira posição, implica mudar a cabeça, ou seja
				// o newNode vira a cabeça
				newNode.next = l.head
				l.head.prev = newNode
				l.head = newNode

			} else {
				aux := l.head
				for i := 0; i < index; i++ {
					aux = aux.next
				}
				//Caso seja a ultima posição da lista, newNode virará a cauda
				if aux.next == nil {
					l.tail.next = newNode
					newNode.prev = l.tail

				} else {
					// inserindo elemento entre dois
					aux.prev.next = newNode
					newNode.prev = aux.prev
					newNode.next = aux
					aux.prev = newNode
				}
			}
		}
		//Aumentando o tamanho da lista
		l.insertSize += 1

		return nil
	}

	return errors.New("Index inválido")

}
