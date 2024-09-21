package implentacoes

import(
	"fmt"
	""
)

func main(){
	l := &ArrayList{}
	l.Init(2)
	l.Add(10)
	l.Add(6)
	l.Add(4)
	fmt.Println(l.Size())
}