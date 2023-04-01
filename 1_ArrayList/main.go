<<<<<<< HEAD
package main

import (
	"fmt"
	"sample-app/list"
)

func main() {
	arraylist := list.ArrayList{}
	arraylist.Init()

	limit := 5

	for i := 0; i < limit; i++ {
		arraylist.Add(i)
	}
	for i := 0; i < limit; i++ {
		fmt.Println(arraylist.Get(i))
	}

	fmt.Println("######################")
	fmt.Println("O tamanho da lista é de: ",arraylist.Size()," valores")
	fmt.Println("######################")
	arraylist.Add(-3)
	for i := 0; i < arraylist.Size(); i++ {
		fmt.Println(arraylist.Get(i))
	}
	fmt.Println("O tamanho da lista é de: ",arraylist.Size()," valores")

	arraylist.Remove()
	fmt.Println("######################")
	fmt.Println("O tamanho da lista é de: ",arraylist.Size()," valores")

	for i := 0; i < arraylist.Size(); i++ {
		fmt.Println(arraylist.Get(i))
	}
	arraylist.Remove()
	arraylist.Remove()
	arraylist.Remove()
	arraylist.Remove()
	arraylist.Remove()
	arraylist.Remove()
	
}
=======
package main

import (
	"fmt"
	"sample-app/list"
)

func main() {
	arraylist := list.ArrayList{}
	arraylist.Init()

	limit := 5

	for i := 0; i < limit; i++ {
		arraylist.Add(i)
	}
	for i := 0; i < limit; i++ {
		fmt.Println(arraylist.Get(i))
	}

	fmt.Println("######################")
	fmt.Println("O tamanho da lista é de: ",arraylist.Size()," valores")
	fmt.Println("######################")
	arraylist.Add(-3)
	for i := 0; i < arraylist.Size(); i++ {
		fmt.Println(arraylist.Get(i))
	}
	fmt.Println("O tamanho da lista é de: ",arraylist.Size()," valores")

	arraylist.Remove()
	fmt.Println("######################")
	fmt.Println("O tamanho da lista é de: ",arraylist.Size()," valores")

	for i := 0; i < arraylist.Size(); i++ {
		fmt.Println(arraylist.Get(i))
	}
	arraylist.Remove()
	arraylist.Remove()
	arraylist.Remove()
	arraylist.Remove()
	arraylist.Remove()
	arraylist.Remove()
	
}
>>>>>>> master
