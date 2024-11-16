package tad

type Ilist interface {
	Add(value int)
	AddOnIndex(value int, index int) error
	RemoveOnIndex(index int) error
	Set(value int, index int) error
	Size() int
	Get(index int) (error, int)
	
}
