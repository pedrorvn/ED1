package queue

import (
	"fmt"
	"errors"
)

type Fila struct {
	valores []int
	cap int
	tam    int
}

func (fila *Fila) NovaFila(size int) {
	fila.valores = make([]int, size)
	fila.tam = 0
	fila.cap = size
}

func (fila *Fila) Push(value int) error {
	if (fila.tam) < fila.cap{
		fila.valores[fila.tam] = value
		fila.tam++
		return nil
	} else{
		return errors.New("Full queue")
	}

}

func (fila *Fila) Pop() (error){
	aux := Fila{}
	aux.NovaFila(fila.cap)
	if fila.tam > 0 {
		for i:=0;i<(fila.cap-1);i++{
			aux.valores[i] = fila.valores[i+1]
		}
		fila.valores = aux.valores
		fila.tam--
		return nil
	} else{
		return errors.New("A fila está vazia, não há como desenfileirar um item")
	}
}

func (fila *Fila) Peek() (int,error) {
	if fila.tam == 0 {
		return 0,errors.New("Empty queue")
	}
	return fila.valores[0],nil
}

func (fila *Fila) Size() int{
	return fila.tam
}

func (fila *Fila) Empty() bool{
	if fila.tam == 0 {
		return true
	} else{
		return false
	}
}

func (fila *Fila) Full() bool{
	if fila.cap == fila.tam {
		return true
	} else{
		return false
	}
}

func (fila *Fila) Imprimir() {
	fmt.Print("[")
	for i:=0; i<fila.tam; i++ {
		fmt.Print(fila.valores[i])
		if (i+1) != fila.tam{
			fmt.Print(", ")
		}
 	}
	fmt.Println("]")


}