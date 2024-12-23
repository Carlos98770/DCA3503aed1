package main

import (
	"errors"
	"fmt"
)

type BSTNode struct {
	val   int
	left  *BSTNode
	right *BSTNode
}

func createNode(value int) *BSTNode {
	return &BSTNode{val: value}
}

func (node *BSTNode) Add(value int) *BSTNode {

	if node == nil {
		node = createNode(value)
	} else {

		if value > node.val {
			node.right = node.right.Add(value)
		} else {
			node.left = node.left.Add(value)
		}

	}

	return node

}

func (node *BSTNode) Search(value int) bool {

	if node == nil {
		return false
	} else {
		if node.val == value {
			return true
		} else if value > node.val {
			return node.right.Search(value)
		} else {
			return node.left.Search(value)
		}
	}

}

func (node *BSTNode) Min() (error, int) {
	if node == nil {
		return errors.New("Arvore está vazia"), -1
	}
	if node.left == nil {
		return nil, node.val
	}
	return node.left.Min()
}

func (node *BSTNode) Max() (error, int) {
	if node == nil {
		return errors.New("Arvore está vazia"), -1
	}
	if node.right == nil {
		return nil, node.val
	}

	return node.right.Max()
}
func (node *BSTNode) PrintPre() {

	if node == nil {
		return
	} else {
		fmt.Println(node.val)
		fmt.Print(" ")
		if node.left != nil {
			node.left.PrintPre()
		}
		if node.right != nil {
			node.right.PrintPre()
		}
	}

}

func (node *BSTNode) PrintIn() {

	if node == nil {
		return
	} else {
		if node.left != nil {
			node.left.PrintIn()
		}
		fmt.Println(node.val)
		fmt.Print(" ")

		if node.right != nil {
			node.right.PrintIn()
		}

	}

}
func (node *BSTNode) PrintPos() {

	if node == nil {
		return
	} else {
		if node.left != nil {
			node.left.PrintPos()
		}

		if node.right != nil {
			node.right.PrintPos()
		}
		fmt.Println(node.val)
		fmt.Print(" ")

	}

}

func (node *BSTNode) PrintLevels() {}

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func (node *BSTNode) Height() int {

	if node == nil {
		return -1
	} else {
		return 1 + max(node.left.Height(), node.right.Height())
	}

}
func (node *BSTNode) Remove(value int) *BSTNode {
	//Buscando o no com o valor
	if node != nil {
		if value > node.val {
			node.right = node.right.Remove(value)
		} else if value < node.val {
			node.left = node.left.Remove(value)
		} else {
			//Encontramos o nó a ser removido

			//Caso 1, nó folha
			if node.left == nil && node.right == nil {
				return nil
			}

			//caso 2, Nó com um filho
			if node.left != nil && node.right == nil {
				return node.left
			} else if node.left == nil && node.right != nil {
				return node.right
			} else { //Caso 3, nó com dois filhos
				_, maiorEsq := node.left.Max()
				node.val = maiorEsq

				node.left = node.left.Remove(maiorEsq)
			}

		}
	}
	return node
}

func (node *BSTNode) Size() int {
	size := 1

	if node.left != nil {
		size += node.left.Size()
	}
	if node.right != nil {
		size += node.right.Size()
	}
	return size
}

func (node *BSTNode) isBst() bool {
	// Caso base: se o nó for nulo, é considerado uma BST válida
	if node == nil {
		return true
	}

	// Verificar se o valor do nó atual está dentro da ordem de uma BST
	// Certificando-se de que o nó esquerdo é menor e o nó direito é maior
	if node.left != nil && node.val <= node.left.val {
		return false
	}
	if node.right != nil && node.val >= node.right.val {
		return false
	}

	// Verificar as subárvores à esquerda e à direita recursivamente
	return node.left.isBst() && node.right.isBst()
}

func main() {
	BST := &BSTNode{val: 2}
	BST.Add(1)
	BST.Add(3)
	BST.Add(5)
	BST.Add(4)
	BST.Add(6)

	fmt.Println(BST.Impar())
}
