package sort

func InsertionSort(v []int){
	size := len(v)
  for i:=1; i<size; i++{
    for j:=i; j>0 && v[j] < v[j-1] ; j--{
      v[j] , v[j-1] = v[j-1], v[j]
    }
  }
}