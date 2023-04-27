package sort

func MergeSort(v []int){
	if len(v) <= 1 {
		return
	}
	metade := len(v)/2
	v1 := make([]int,metade)
	v2 := make([]int,len(v)-metade)

	for i:=0;i<metade;i++{
		v1[i]=v[i]
	}
	for i:=metade; i<len(v);i++{
		v2[i-metade] = v[i]
	}
	
	MergeSort(v1)
	MergeSort(v2)
	copy(v,Merge(v1, v2))
}

func Merge(v1 []int, v2 []int) []int{
	i,k,j := 0,0,0
	mesclado := make([]int, len(v1)+len(v2))

	for j<len(v1) && k < len(v2) {
		if v1[j] < v2[k] {
		mesclado[i] = v1[j]
		j++
		} else{
		mesclado[i]=v2[k]
		k++
		}
		i++
	}
	for j<len(v1){
		mesclado[i]=v1[j]
		i++
		j++
	}
	for k<len(v2){
		mesclado[i]=v2[k]
		i++
		k++
	}

	return mesclado	
}