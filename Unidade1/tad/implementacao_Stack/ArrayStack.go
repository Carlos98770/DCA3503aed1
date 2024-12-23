package main

import ( "errors"
	"fmt"
)

type ArrayStack struct {
	v []int
	topo int
	insertSize int

}

func (ad *ArrayStack) Init(size int) error {
	if size > 0 {
		ad.v = make([]int, size)
		ad.insertSize = 0
		return nil
	} else {
		return errors.New("size <=0")
	}

}

func (ad *ArrayStack) duplicarVetor() {
	newSize := len(ad.v) * 2
	newV := make([]int, newSize)
	

	for i := 0; i < ad.insertSize; i++ {
		newV[i] = ad.v[i]

	}

	ad.v = newV

}

func (ad *ArrayStack) Push (value int) {

	if ad.insertSize == len(ad.v){
		ad.duplicarVetor()
	}
	ad.topo = ad.insertSize
	ad.v[ad.insertSize] = value
	ad.insertSize++

}

func (ad *ArrayStack) Pop () (int,error) {

        if ad.insertSize == 0{
		return -1,errors.New("Pilha Vazia")
	}else {

		aux := ad.v[ad.topo]
		ad.topo -=1
		ad.insertSize-=1
		
		return aux,nil
	}

	
}

func (ad *ArrayStack) Peek () (int, error){
	if ad.insertSize == 0{
                return -1,errors.New("Pilha Vazia")
        }else{
                return ad.v[ad.topo],nil
        }

	

}

func (ad *ArrayStack) IsEmpty () bool{
        if ad.insertSize == 0{
                return true
        }else{
                return false
        }



}

func (ad *ArrayStack) Size () int{
        return ad.insertSize

}
