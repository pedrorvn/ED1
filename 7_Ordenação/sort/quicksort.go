package sort

import (
	
)

func QuickSort(v []int, start int, end int){
	if end > start {
		pivo := partition(v,start,end)
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