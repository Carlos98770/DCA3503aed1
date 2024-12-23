package main

import ( "errors"
	"fmt"	
)

type Node struct {
	elemento int
	next *Node
}

type LinkedStack struct {
	top *Node
	insertSize int
}

func (ls *LinkedStack) Push(value int) {
	newNode := &Node{elemento : value}
	if ls.top == nil{
		ls.top = newNode
	} else {
		newNode.next = ls.top
		ls.top = newNode

	}
}

func (ls *LinkedStack) Pop () (int,error){

        if ls.top == nil{
        	return -1, errors.New("pilha vazia")
        } else {
		aux := ls.top.elemento
        	ls.top = ls.top.next
		ls.insertSize -=1
		return aux,nil
        }
}

func (ls *LinkedStack) Peek () (int,error){

        if ls.top == nil{
                return -1, errors.New("pilha vazia")
        } else {
                aux := ls.top.elemento
                return aux,nil
        }
}

func (ls *LinkedStack) IsEmpty () bool{

        if ls.top == nil{
                return true
        } else{
                return false
        }
}

func (ls *LinkedStack) Size () int{

        
        return ls.insertSize
        
}


