package sort

import (
	//"fmt"
	"math/rand"
	"time"
)

func QuickSort(v []int, start int, end int){
	if end > start {
		pivo := partition2(v,start,end)
		QuickSort(v,start,pivo-1)
		QuickSort(v,pivo+1,end)
	}
}

func partition(v []int, ini int, fim int) int{
	pIndex := ini
	pivo := v[fim]
	for i:=ini; i<fim; i++ {
		if v[i]<=pivo {
			v[i], v[pIndex] = v[pIndex], v[i]
			pIndex++
		}
	}
	v[pIndex], v[fim] = v[fim], v[pIndex]
	return pIndex
}
// Função que escolhe o pivo de forma aleatória
func partition2(v []int, ini int, fim int) int{
	//Escolhendo o pivo
	rand.Seed(time.Now().UnixNano())
	aleatorio := rand.Intn(fim - ini + 1) + ini
	//Mandando o pivo para o fim p/ aproveitar o resto do código
	v[aleatorio],v[fim] = v[fim], v[aleatorio]

	//Continuação do quicksort tradicional
	pIndex := ini
	pivo := v[fim]
	for i:=ini; i<fim; i++ {
		if v[i]<=pivo {
			v[i], v[pIndex] = v[pIndex], v[i]
			pIndex++
		}
	}
	v[pIndex], v[fim] = v[fim], v[pIndex]
	return pIndex
}

