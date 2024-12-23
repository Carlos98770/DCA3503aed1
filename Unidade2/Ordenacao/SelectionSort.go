
package main

import (
	"fmt"
	"math"
)

func SelectionSort_inPlace(v []int) {    

	for i := 0; i < len(v)-1; i++ {
		menorID := i

		for j := i + 1; j < len(v); j++ {

			if v[menorID] > v[j] {
				menorID = j
			}
		}

		v[menorID], v[i] = v[i], v[menorID]

	}
}

func SelectionSort_OutPlace(v []int) []int {
	v_ord := make([]int, len(v))
	Ind_ord := 0
	max := math.MaxInt64

	for i := 0; i < len(v); i++ {
		menorID := i
		for j := 0; j < len(v); j++ {
			if v[menorID] > v[j] {
				menorID = j
			}
		}

		v_ord[Ind_ord] = v[menorID]
		v[menorID] = max
		Ind_ord++
	}

	return v_ord

}
