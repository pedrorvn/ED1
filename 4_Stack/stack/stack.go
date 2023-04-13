package stack

import (
	//"fmt"
	"errors"
)


type Stack struct {
	values []int
	tam    int
}


func (pilha *Stack) Init(size int) {
	pilha.values = make([]int,0)
	pilha.tam = 0
}

/*func (pilha *Stack) double() {
	pilhaMaior := make([]int, len(pilha.values)*2)
	for i := 0; i < len(pilha.values); i++ {
		pilhaMaior[i] = pilha.values[i]
	}
	pilha.values = pilhaMaior
}*/

func (pilha *Stack) Push(value int) {
	/*if pilha.tam == len(pilha.values) {
		pilha.double()
	}*/
	pilha.values = append(pilha.values, value)
	pilha.tam++
}

func (pilha *Stack) Pop() (int,error){
    if pilha.tam == 0 {
        return 0, errors.New("stack is empty")
    }

    res := pilha.values[pilha.tam-1]
    pilha.values = pilha.values[:pilha.tam-1]
	pilha.tam--
    return res, nil
}

func (pilha *Stack) Peek() (int,error) {
	if pilha.tam > 0 {
		return pilha.values[pilha.tam-1],nil
	} else{
		return 0,errors.New("You need to insert values on your stack first.")
	}
}

func (pilha *Stack) Size() int {
	return pilha.tam
}

func (pilha *Stack) Empty() bool {
	if pilha.tam != 0 {
		return false
	} else {
		return true
	}

}

