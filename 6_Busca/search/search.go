package search

import "errors"

func BuscaSeq (l []int, val int) (int,error){
	tam := len(l)
	for i:=0; i<tam; i++{
		if l[i]==val{
			return i,nil
		}
	}
	return 0,errors.New("Value not found")
}

func BuscaBinariaRec (l []int, val int, inicio int, fim int) (int,error) {
    if fim < inicio {
        return 0,errors.New("not found or start value is smaller than end value")
    }
    meio := (inicio + fim) / 2
    if val == l[meio] {
        return meio,nil
    } else if val < l[meio] {
        return BuscaBinariaRec(l, val, inicio, meio-1)
    } else{
        return BuscaBinariaRec(l, val, meio+1, fim)
    }
	
}