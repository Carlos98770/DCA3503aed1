package implementacoes_listas


import (
	"errors"
	
)


type Node struct {
	elemento int
	next *Node
}

type LinkedList struct {
	head *Node
	insertSize int
}

func (l *LinkedList) Add(value int) {
	newNode := &Node{elemento: value}
	if l.head == nil{
		l.head = newNode
	} else {
		aux := l.head
		for i:=0; i < l.insertSize-1; i++{
			aux = aux.next
		}

		aux.next = newNode
	}
	l.insertSize +=1
}

func (l *LinkedList) AddOnIndex(value int, index int) error{
	newNode := &Node{elemento: value}
	if index >=0 && index < l.insertSize{
		if index == 0{
			l.head = newNode
		} else {
			aux := l.head
			for i:=0; i < index - 1; i++{
				aux = aux.next
			}
			newNode.next = aux.next
			aux.next = newNode
		}
		l.insertSize+=1
		return nil
	}
	return errors.New("Index invalido")
}

func (l *LinkedList) RemoveOnIndex(index int) error{
	if index >=0 && index < l.insertSize{
		if l.insertSize == 0{
			return errors.New("Lista vazia")
		} else if index == 0{
			l.head = l.head.next
		} else {
			aux := l.head
			for i:=0; i < index-1; i++{
				aux = aux.next
			}

			aux.next = aux.next.next
		}
		l.insertSize -=1
		return nil
	}

	return errors.New("Index invalido")
}

func (l *LinkedList) Set(value int, index int) error{
	if index >=0 && index < l.insertSize{
		if index == 0{
			l.head.elemento = value
		} else {
			aux := l.head
			for i:=0; i < index - 1; i++{
				aux = aux.next
			}

			aux.next.elemento = value
		}
		return nil
	}
	return errors.New("Index invalido")
}

func (l *LinkedList) Size() int {
	return l.insertSize
}

func (l *LinkedList) Get(index int) (int,error) {
	if index >=0 && index < l.insertSize{
		if index == 0{
			return l.head.elemento, nil
		} else {
			aux := l.head
			for i:=0; i < index - 1; i++{
				aux = aux.next
			}

			return aux.next.elemento, nil
		}
	}//prev = nil
	return -1,errors.New("Index invalido")
}

func (l *LinkedList) Reverse(){
	var prev *Node
	atual := l.head
	//next := atual.next

	for atual != nil  {
		next := atual.next
		atual.next = prev
		prev = atual
		atual = next
		
	}
	
	l.head = prev

}

func (l *LinkedList) OnlyPar(){
	var prev *Node
	cont := 0

	if l.head.elemento %2==0{
		l.head = l.head.next
		cont++
	}else {
		prev =l.head
		current := l.head.next

		for current != nil{
			next:= current.next

			if current.elemento%2 == 0{
				prev.next = next
				current.next = nil
				cont++
			} else {
				prev = current
			}

			current = next
	
			
		}
	
	}
	
	l.insertSize = l.insertSize - cont


}









