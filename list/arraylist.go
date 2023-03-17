package list

import "fmt"

type ArrayList struct {
	values []int
	tam    int
}

func (arraylist *ArrayList) Init() {
	arraylist.values = make([]int, 10)
}

func (arraylist *ArrayList) double() {
	doubledValues := make([]int, len(arraylist.values)*2)
	for i := 0; i < len(arraylist.values); i++ {
		doubledValues[i] = arraylist.values[i]
	}
	arraylist.values = doubledValues
}

func (arraylist *ArrayList) Add(value int) {
	if arraylist.tam == len(arraylist.values) {
		arraylist.double()
	}
	arraylist.values[arraylist.tam] = value
	arraylist.tam++
}

func (arraylist *ArrayList) AddOnIndex(value int, index int) {
	if arraylist.tam == len(arraylist.values) {
		arraylist.double()
	}
	for i := arraylist.tam; i > index; i-- {
		arraylist.values[i] = arraylist.values[i-1]
	}
	arraylist.values[index] = value
	arraylist.tam++
}

func (arraylist *ArrayList) Remove() {
	if arraylist.tam > 0 {
		arraylist.tam--
	} else{
		fmt.Println("A lista está vazia, não há como remover um item")
	}
}

func (arraylist *ArrayList) RemoveOnIndex(index int) {
	if index >= 0 && index < arraylist.tam {
		for i := index; i < arraylist.tam; i++ {
			arraylist.values[i] = arraylist.values[i+1]
		}
		arraylist.tam--
	} else{
		fmt.Println("Índice informado é inválido")
	}
}

func (arraylist *ArrayList) Get(index int) int {
	if index >= 0 && index < arraylist.tam {
		return arraylist.values[index]
	}
	return -1 //TODO: add error handling
}

func (arraylist *ArrayList) Set(value, index int) {
	if index >= 0 && index < arraylist.tam {
		arraylist.values[index] = value
	}
	//else //TODO: add error handling
}

func (arraylist *ArrayList) Size() int {
	return arraylist.tam
}

