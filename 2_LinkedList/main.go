package main

import (
	"fmt"
	"sample-app/list"
)

func main() {
	lista := list.LinkedList{}

	for i := 0; i < 8; i++{
		lista.InserirElementoNoFim(i)
	}

	lista.ImprimirLista()
	fmt.Println(lista.Tamanho())
	lista.Remove(5)
	lista.ImprimirLista()
	lista.InserirEmPosicao(24,3)
	lista.ImprimirLista()
}