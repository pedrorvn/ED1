package main

import (
	"fmt"
	"sample-app/queue"
)

func main() {
	tamanho_da_fila := 10
	fila := queue.Fila{}
	fila.NovaFila(tamanho_da_fila)

	fila.Push(5)
	fila.Push(9)
	fila.Push(25)
	fmt.Println(fila.Size())
	fmt.Println(fila.Peek())
	fila.Imprimir()
	fmt.Println(fila.Pop())
	fila.Pop()
	fila.Imprimir()
	fmt.Println(fila.Pop())
	fmt.Println(fila.Pop())
	fmt.Println(fila.Size())
	fmt.Println(fila.Empty())

	for i:=0; i<tamanho_da_fila; i++ {
		fmt.Println(i)
		fila.Push(i)
		fila.Imprimir()
	}
	fila.Imprimir()
	fmt.Println(fila.Push(8))
	fila.Pop()
	fila.Imprimir()


}