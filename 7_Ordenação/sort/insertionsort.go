package sort

func InsertionSort(v []int){
	n := len(v)
	for i:=1;i<n;i++{
		valor:= v[i]
		var j int
		for j = i; j>0 && v[j-1] > valor;j--{
			v[j] = v[j-1]
		}
		v[j] = valor
	}
}