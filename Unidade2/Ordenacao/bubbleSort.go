package main

import "fmt"

func bubbleSort(v []int){ 	

	
	for i:= 0; i < len(v) - 1; i++{   //laço indo até n-1, assim eliminando a ultima comparação
		troca := false
		for j:= 0; j < len(v)-1 - i; j++{  // A cada varredura, os elementos finais do vetor ja está ordenadas

			if v[j] > v[j+1]{
				v[j],v[j+1] = v[j+1],v[j]		// Realiza a troca
				troca = true
			}

			

		}
		if troca == false{
			return
		}
	

	}

}
