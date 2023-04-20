package sort

func BubbleSort(v []int, n int){    
    for varredura := 0; varredura < n-1; varredura++ {
        trocou := false
        for i := 0; i < n-varredura-1; i++{
            if v[i]>v[i+1]{
                temp := v[i];
                v[i] = v[i+1];
                v[i+1] = temp;
                trocou = true;
            }
        }
		if trocou == false {
			break
		}
    }
}