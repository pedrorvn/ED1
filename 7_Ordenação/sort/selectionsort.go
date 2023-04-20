package sort

func SelectionSort(v []int){
	n := len(v)

    for varredura := 0; varredura < n-1; varredura++ {
        menorIndice := varredura
        for i := varredura+1; i < n; i++{
            if v[i]<v[menorIndice]{
                menorIndice = i
            }
        }
		v[varredura], v[menorIndice] = v[menorIndice], v[varredura]

    }
}