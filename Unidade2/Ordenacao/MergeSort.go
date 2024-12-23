package main

import "fmt"

/*

    -> Ordenação por mesclagem
    -> É classificado como um algoritmo de Divisão e Conquista.
        Divisão: dividimos recursivamente o array pela metade, até que ele seja indivisível
        Conquista: mesclamos (juntamos/combinamos) os arrays de forma ordenada
*/



func Merge(v []int, vLeft, vRight []int) {
    i, j, k := 0, 0, 0

    // Compara os elementos dos dois subvetores e coloca no vetor original
    for i < len(vLeft) && j < len(vRight) {
        if vLeft[i] < vRight[j] {
            v[k] = vLeft[i]
            i++
        } else {
            v[k] = vRight[j]
            j++
        }
        k++
    }

    // Copia os elementos restantes de vLeft, se houver
    for i < len(vLeft) {
        v[k] = vLeft[i]
        i++
        k++
    }

    // Copia os elementos restantes de vRight, se houver
    for j < len(vRight) {
        v[k] = vRight[j]
        j++
        k++
    }
}


func MergeSort(v [] int){			// v = [3,3,4,6,9,1,2,5,7,8]
    if len(v) > 1 {
        tamLeft:= int(len(v)/2)
        tamRight := len(v) - tamLeft
        vLeft := make([]int, tamLeft)
        vRight := make([]int , tamRight)

        
        // copia dos valores
        for i:= 0; i < len(v); i++{
            if i < tamLeft{
                vLeft[i] = v[i]
            } else {
                vRight[i - tamLeft] = v[i]
            }

        }

        MergeSort(vLeft)
        MergeSort(vRight)
        Merge(v,vLeft,vRight)
            
    }
	

}



func main(){
	v := []int{3,3,4,6,9,1,2,5,7,8}
	

	MergeSort(v)

	
	

	for i:=0; i < len(v); i++{
		fmt.Println(v[i])
	}
	


}
