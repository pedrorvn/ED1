package main

import (
	//"fmt"
	"sample-app/tree"
)

func main (){
	//                9
	//           7        12
	//         5   8    10   13
	//       1                   14
	//
	//
	//
	//
	rootNode := tree.NewNode(9)
	rootNode.Add(7)
	rootNode.Add(12)
	rootNode.Add(5)
  	rootNode.Add(8)
  	rootNode.Add(10)
	rootNode.Add(13)
	rootNode.Add(1)
	rootNode.Add(14)

  	//rootNode = rootNode.Remove(12)
  	//rootNode = rootNode.Remove(9)

	//rootNode.PrintIn()
  	//fmt.Println("Size: ", rootNode.Size())
	rootNode.PrintLevels()

  	//fmt.Println("IsBst: ", nonBstNode.IsBst())

}