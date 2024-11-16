package implementacoes_Queue

import (
	"errors"
)

type Node2 struct{
	elemento int
	next *Node2
}

type LinkedQueue struct{
	insertSize int
	front *Node2
	rear *Node2
}

func (lq *LinkedQueue) Enqueue(value int) error  {
	newNode := &Node2{elemento:value}
	if lq.front == nil{
		lq.front = newNode
		lq.rear = newNode
	} else {
		lq.rear.next = newNode
		lq.rear = newNode
	}

	lq.insertSize++

	return nil

}

func (lq *LinkedQueue) Dequeue() (int,error)  {
	if lq.front == nil{
		return -1,errors.New("Fila vazia")
	} else {
		aux := lq.front.elemento
		lq.front = lq.front.next
		lq.insertSize--
		return aux,nil
	}

	
}

func (lq *LinkedQueue) Front() (int,error)  {
	if lq.insertSize == 0 {
		return -1,errors.New("Fila está vazia")
	} else {
		aux := lq.front.elemento
		return aux,nil
	}

}

func (lq *LinkedQueue) IsEmpty() bool  {
	if lq.insertSize == 0 {
		return true
	} else {
		return false
	}

}

func (lq *LinkedQueue) Size() int  {
	return lq.insertSize
}


