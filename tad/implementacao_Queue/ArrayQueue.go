package implementacoes_Queue

import (
	"errors"

)

type ArrayQueue struct{
	v	[]int
	front int
	rear int
	insertSize int
}

func (aq *ArrayQueue) Init(size int) error{
	if size > 0{
		aq.v = make([]int,size)
		aq.insertSize = 0
	} else {

		return errors.New("Impossivel inicializar o array")
	}
	return nil
}

func (aq *ArrayQueue) Enqueue(value int) error  {
	if aq.insertSize == len(aq.v) {
		return errors.New("Fila está cheia")
	} else {
		if  aq.insertSize == 0{
			aq.front++
			aq.rear++
		}else {
			aq.rear = (aq.rear + 1)% len(aq.v)
		}
	}

	aq.insertSize++
	aq.v[aq.rear] = value
	return nil
	
}

func (aq *ArrayQueue) Dequeue() (int,error)  {
	if aq.insertSize == 0 {
		return -1,errors.New("Fila está vazia")

	} else if aq.insertSize == 1{

		aq.insertSize--
		aux := aq.v[aq.front]
		aq.front = -1
		aq.rear = -1
		return aux,nil

	} else {
		aux := aq.v[aq.front]
		aq.front = (aq.front + 1)% len(aq.v)
		aq.insertSize--
		return aux,nil
	}

}

func (aq *ArrayQueue) Front() (int,error)  {
	if aq.insertSize == 0 {
		return -1,errors.New("Fila está vazia")
	} else {package ed

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
		
		
		
		aux := aq.v[aq.front]
		return aux,nil
	}

}

func (aq *ArrayQueue) IsEmpty() bool  {
	if aq.insertSize == 0 {
		return true
	} else {
		return false
	}

}

func (aq *ArrayQueue) Size() (int,error)  {
	if aq.front != -1 && aq.rear != -1{
		if aq.rear >= aq.front{
			return (aq.rear - aq.front) + 1,nil
		} else {
			return len(aq.v) - aq.front + aq.rear + 1,nil
		}
	}
	return -1, errors.New("Lista vazia")

}



