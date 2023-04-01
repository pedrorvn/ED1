package main

import (
	//"fmt"
	"sample-app/list"
)

func main() {
	lista := list.DoublyLinkedList{}
	for i := 0; i < 5; i++ {
		lista.InserirElementoNoFim(i)
	}

	lista.ImprimirLista()
	lista.InserirElementoNoInicio(9)
	lista.ImprimirLista()
	lista.InserirElementoEmPosicao(3, 2)
	lista.ImprimirLista()
	lista.Remover(0)
	lista.ImprimirLista()
	lista.Remover(6)

}
