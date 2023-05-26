package sort

func CountingSort (v []int) {
	maior := v[0]
	menor := v[0]
	//Encontrando o menor e o maior elemento
	for i:=1;i<len(v);i++{
		if v[i]<menor {menor = v[i]}
		if v[i]>maior {maior = v[i]}
	}

	//Criando o vetor de contagem
	tamC := maior-menor+1
	c := make([]int,tamC)

	//contar ocorrencias de cada elemento em v

	for i:=0; i<len(v);i++{
		c[v[i]-menor]++
	}

	//somar cumulativamente
	for i:=1;i<tamC;i++ {c[i]+=c[i-1]}

	//criar o vetor para elementos ordenados
	o := make([]int,len(v))

	//vetor para detectar se uma dada posição está ocupada (util no caso com repetição de valores)
	ocupado := make([]bool,len(v))


	// mapear elementos de v para o ordenadamente
	for i := 0; i < len(v); i++ {
		indiceO := c[v[i]-menor] - 1
		for ocupado[indiceO] {
			indiceO--
		}
		o[indiceO] = v[i]
		ocupado[indiceO] = true
	}

	copy(v,o)
}