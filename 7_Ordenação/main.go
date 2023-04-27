package main

import(
	"fmt"
	"math/rand"
	"time"
	"sample-app/sort"
)

func main() {
	// Define o tamanho do slice de inteiros e gera valores aleatórios
	n := 10
	rand.Seed(time.Now().UnixNano())
	v := make([]int, n)
	for i := 0; i < n; i++ {
		v[i] = rand.Intn(100)
	}

	// Imprime o slice desordenado
	fmt.Println("Slice desordenado:", v)

	// Ordena o slice usando Bubble Sort
	sort.BubbleSort(v, n)
	fmt.Println("Slice ordenado pelo Bubble Sort:", v)

	// Gera novos valores aleatórios para o slice
	for i := 0; i < n; i++ {
		v[i] = rand.Intn(100)
	}

	// Imprime o slice desordenado
	fmt.Println("Slice desordenado:", v)

	// Ordena o slice usando Selection Sort
	sort.SelectionSort(v)
	fmt.Println("Slice ordenado pelo Selection Sort:", v)

	// Gera novos valores aleatórios para o slice
	for i := 0; i < n; i++ {
		v[i] = rand.Intn(100)
	}

	// Imprime o slice desordenado
	fmt.Println("Slice desordenado:", v)

	// Ordena o slice usando Insertion Sort
	sort.InsertionSort(v)
	fmt.Println("Slice ordenado pelo Insertion Sort:", v)

	// Gera novos valores aleatórios para o slice
	for i := 0; i < n; i++ {
		v[i] = rand.Intn(100)
	}
	// Imprime o slice desordenado
	fmt.Println("Slice desordenado:", v)

	// Ordena o slice usando Merge Sort
	sort.MergeSort(v)
	fmt.Println("Slice ordenado pelo Merge Sort:", v)
}
