package main

import ("fmt"
		"math/rand"		
)

func partition(v []int, ini int, fim int) int {

	pIndex := ini
	indPivo := fim
	for i := ini; i < fim; i++ {

		if v[i] <= v[indPivo] {
			v[pIndex], v[i] = v[i], v[pIndex]
			pIndex++																		
																								
		}																

	}

	v[pIndex], v[indPivo] = v[indPivo], v[pIndex]

	return pIndex

}

func partition_randPivo(v []int, ini int, fim int) int {

	indPivoRand := rand.Intn((fim - ini)) + ini

	v[indPivoRand], v[fim] = v[fim] , v[indPivoRand]


	pIndex := ini
	indPivo := fim
	for i := ini; i < fim; i++ {

		if v[i] <= v[indPivo] {
			v[pIndex], v[i] = v[i], v[pIndex]														
			pIndex++																			
		}																						

	}

	v[pIndex], v[indPivo] = v[indPivo], v[pIndex]

	return pIndex

}

func quickSort(v []int, ini int, fim int) {

	if fim > ini {
		idPivo := partition(v, ini, fim)

		quickSort(v, ini, idPivo-1)
		quickSort(v, idPivo+1, fim)
	}

}

func main() {

	v := []int{38, 27, 43, 3, 9, 82, 10, 14, 6, 23, 5, 16, 31, 12, 29}

	for i := 0; i < len(v); i++ {
		fmt.Print(v[i])
		fmt.Print(", ")
	}

	fmt.Println(" ")

	quickSort(v, 0, len(v)-1)

	for i := 0; i < len(v); i++ {
		fmt.Print(v[i])
		fmt.Print(", ")
	}

}
