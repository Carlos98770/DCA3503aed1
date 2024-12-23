
package main

import "fmt"

/*
Agora uma ilustração com mais detalhes, que se aproxima mais da implementação.

tempo	Vetor	comentário
t0	8,2,4,3,7	[8] já está ordenado, e [2, 4, 3, 7] desordenado
t1	2,8,4,3,7	2 é menor que 8, por isso trocamos 2 e 8; [2,8] já está ordenado, e [4, 3, 7] desordenado
t2	2,4,8,3,7	4 é menor que 8, por isso trocamos 4 e 8
t3	2,4,8,3,7	4 não é menor que 2, por isso não trocamos 4 e 2; [2,4,8] já está ordenado, e [3, 7] desordenado
t4	2,4,3,8,7	3 é menor que 8, por isso trocamos 8 e 3;
t5	2,3,4,8,7	3 é menor que 4, por isso trocamos 4 e 3;
t6	2,3,4,8,7	3 não é menor que 2, por isso não trocamos 2 e 3; [2,3,4,8] já está ordenado, e [7] desordenado
t7	2,3,4,7,8	7 é menor que 8, por isso trocamos 8 e 7;
t7	2,3,4,7,8	7 é não menor que 4, por isso não trocamos 4 e 7; todo o array está ordenado: [2,3,4,7,8]

*/
							//Implementação out-of-place
func insertionSort_out_of_place(v []int){
	aux := make([]int,len(v))			//Cria o vetor auxiliar

	aux[0] = v[0]						//Insere o na primera posicao do vetor odernado o primeiro elemento do vetor
	for i := 1; i < len(v); i++{		// i = 1 por que o primeiro elemento do vetor ja foi inserido no vetor ordenado
		insert := false					// flag para insert
		for j:= 0; j < i; j++{			// for para percorrer o vetor ordenado
			if v[i] < aux[j]{			// verifica se o elemento da vez é menor que algum elemento do vetor ordenado
				
				for k:=i; k > j; k--{		// caso seja, abri espaco para insercao do mesmo
					aux[k] = aux[k-1]
				}
				
				aux[j] = v[i]				// insere o elemento da vez no vetor ordenado na posição certa
				insert = true				// Garante que insert receba true, isso significa que o elemento da vez não é maior que todos os elementos do vetor ordenado
				break						// para o laço
			}
			

		}	
		if !insert {						// se insert for false, o elemento da vez é o maior, entao a inserção dele é no final do vetor ordenado
			aux[i] = v[i]					//o tamanho do vetor ordenado é controlado pelo indice i, ou seja a ultima posicao do vetor é i
		}
	}

	copy(v,aux)							// copia o vetor auxiliar para o endereço do vetor v
}

										//Implentação in-place
func insertionSort_in_place(v []int){		//v= [8,2,4,3,7]			8 	2,4,3,7  			2,8		4,3,7		2,4,8     3,7	

	for i := 1; i < len(v) ; i++{			//pecorre o vetor para a seleção do elemento da vez
		for j := 0 ; j < i;j++{				// for para pecorrer a parte ordenada
			if v[j] > v[i]{					// se o elemento do vetor ordenado for maior que o elemento da vez
				
				v[j],v[i] = v[i],v[j]		//Troca

			}
			
		}

	}
}

func insertionSort_in_place2(v []int){		//v= [8,2,4,3,7]			8 	2,4,3,7  			2,8		4,3,7		2,4,8     3,7	

	for i := 1; i < len(v) ; i++{
		for j := i ; j > 0 && v[j-1] > v[j];j--{			//for com condição de parada se a ultima posição do vetor ordenado for maior que a proxima posicao e j > 0
			if v[j - 1] > v[j]{
				
				v[j - 1],v[j] = v[j] , v[j-1]
			}
			
		}

	}
}

func main(){

	v := []int{8,2,4,3,7,5,4,2,8,6,63,14,58,72,14,2,4,5,89,6,7,11,21,23,52,65,8,78,46,0,64,32,38,39,9,58,15,56,12,31,56,466,89,789956,456489,459459,49465}
	

	insertionSort_out_of_place(v)				//In-place, estável, O(n²), Ômega(n).

	for i:=0; i < len(v); i++{
		fmt.Println(v[i])
	}


}
