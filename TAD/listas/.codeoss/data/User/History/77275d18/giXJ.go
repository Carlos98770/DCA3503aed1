package main

type Lista interface {
	Get(index int) (int, error)
	Size() int
	Add(elemento int)
	AddAt(elemento int, index int) error
	Remove(index int) error
}

type ArrayList struct{
    v []int
    size int
}

func (l *ArrayList) Size() int {
    return l.size   
}

func(l *ArrayList) Get(index int) int{
    if index > l.size {
        return -1
    }
    
    return l.v[index], nil
    
}
