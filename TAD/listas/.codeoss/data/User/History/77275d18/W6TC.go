package main

type Lista interface {
	Get(index int) (int, error)
	Size() int
	Add(elemento int)
	AddAt(elemento int, index int) error
	Remove(index int) error
}

type ArrayList struct {
	v    []int
	size int
}

func (l *ArrayList) Size() int {
	return l.size
}

func (l *ArrayList) Get(index int) (int, error) {
	if index > l.size {
		return -1, nil
	}

	return l.v[index], nil

}
func (l *ArrayList) 
