package main

import (
	"fmt"
	"sample-app/search"
)

func main() {
	lista := make([]int,10)
	for i:=0;i<10;i++{
		lista[i] = (i*5)
	}
	
	for j:=0;j<10; j++{
		fmt.Print(" ",lista[j]," ")
	}
	
	fmt.Println()
	fmt.Println(search.BuscaSeq(lista,45))
	fmt.Println(search.BuscaBinariaRec(lista,120,0,len(lista)-1))	

}