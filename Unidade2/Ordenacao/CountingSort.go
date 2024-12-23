package main

import "fmt"

// v = [5,9,2,1,3,8,6]
func CountingSort(v []int) {
	menor := v[0]
	maior := v[0]
	//Passo 1, Armazenar o menor e maior valor do vetor v
	//maior = 9
	//menor = 1

	for i := 1; i < len(v); i++ {

		if v[i] < menor {
			menor = v[i]
		}
		if v[i] > maior {
			maior = v[i]
		}

	}

	//Passo 2, Criar o Vetor de contagem
	tamC := (maior - menor) + 1
	c := make([]int, tamC) //vetor Contagem
	//indices Virtuais	 1     2     3     4     5     6     7     8	9
	//indices	//   0     1     2     3     4     5     6     7  , 8
	//C = [  0  ,  0  ,  0  ,  0  ,  0  ,  0  ,  0  ,  0  , 0 ]

	//Passo 3, Contar quantas vezes cada elemento saiu no vetor original e armazenar no array de contagem

	for i := 0; i < len(v); i++ { //indices Virtuais	 1     2     3     4     5     6     7     8	9
		//indices	//   0     1     2     3     4     5     6     7  , 8
		//C = [  1  ,  1  ,  1  ,  0  ,  1  ,  1  ,  0  ,  1  , 1 ]
		c[v[i]-menor]++
	}

	// Passo 4, Realizar as somas acumulativas
	//indices Virtuais	 1     2     3     4     5     6     7     8	9
	//indices	//   0     1     2     3     4     5     6     7  , 8
	//C = [  1  ,  2  ,  3  ,  3  ,  4  ,  5  ,  5  ,  6  , 7 ]
	for i := 1; i < len(c); i++ {
		c[i] = c[i] + c[i-1]
	}

	//Passo 5, Criar um novo array para colocar os elementos ordenados

	vOrd := make([]int, len(v))

	//Passo 5.1, Criar um vetor que mapea se uma determinada posição no vetor ordenado já está ocupada
	ocupado := make([]int, len(v))

	//Passo 6, Mapear os elementos do vetor desordenado para o vetor ordenado utilizando o vetor contagem
	for i := 0; i < len(vOrd); i++ {
		indOrd := c[v[i]-menor]
		for ocupado[indOrd-1] == 1 {
			indOrd -= 1
		}

		vOrd[indOrd-1] = v[i]
		ocupado[indOrd-1] = 1
	}

	copy(v, vOrd)

}

func main() {
	v := []int{3, 3, 6, 9, 1, 2, 5, 7, 8, 20}

	CountingSort(v)

	for i := 0; i < len(v); i++ {
		fmt.Println(v[i])
	}

}
