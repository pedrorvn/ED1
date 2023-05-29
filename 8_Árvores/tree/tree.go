package tree

import (
	"fmt"
)

type BstNode struct {
	left  *BstNode
	value int
	right *BstNode
}

func NewNode(value int) *BstNode {
	node := BstNode{}
	node.value = value
	return &node
}

func (bstNode *BstNode) Add(value int) {
	if value <= bstNode.value {    //insercao à esq
		if bstNode.left == nil {
			bstNode.left = NewNode(value)
		} else {
			bstNode.left.Add(value)
		}
	} else {
		if bstNode.right == nil {
			bstNode.right = NewNode(value)
		} else {
			bstNode.right.Add(value)
		}
	}
}
//Função de busca de forma recursiva
func (bstNode *BstNode) Search(value int) bool{
	if value == bstNode.value{
	  return true
	} else if value < bstNode.value && bstNode.left != nil {
	  return bstNode.left.Search(value)
	} else if value > bstNode.value && bstNode.right != nil {
	  return bstNode.right.Search(value)
	} else {
	  return false
	}
}	
//Função de buscar o valor minimo recursiva
func (bstNode *BstNode) Min() int{
	if bstNode.left == nil {
	  return bstNode.value
	} else {
	  return bstNode.left.Min()
	}
  }
//Função de buscar o valor máximo recursiva
  func (bstNode *BstNode) Max() int{
	if bstNode.right == nil {
	  return bstNode.value
	} else {
	  return bstNode.right.Max()
	}
}

//Verificar altura (recursiva)
func (bstNode *BstNode) Height() int {
	hEsq := 0
	hDir := 0
	if bstNode.left != nil {
	  hEsq = 1 + bstNode.left.Height()
	}
	if bstNode.right != nil {
	  hDir = 1 + bstNode.right.Height()
	}
	if hEsq >= hDir{
	  return hEsq
	} else{
	  return hDir
	}
}
//Funções de navegação/impressão
//pre ordem raiz,esq,dir
func (bstNode *BstNode) PrintPre() {
	fmt.Print(bstNode.value,", ")
	if bstNode.left != nil {
	  bstNode.left.PrintPre()
	}
	if bstNode.right != nil {
	  bstNode.right.PrintPre()
	}
  }

//em ordem esq, raiz, dir
func (bstNode *BstNode) PrintIn() {
	if bstNode.left != nil {
	  bstNode.left.PrintIn()
	}
	fmt.Print(bstNode.value,", ")
	if bstNode.right != nil {
	  bstNode.right.PrintIn()
	}
  }
  
  //pos ordem esq, dir, raiz
  func (bstNode *BstNode) PrintPos() {
	if bstNode.left != nil {
	  bstNode.left.PrintPos()
	}
	if bstNode.right != nil {
	  bstNode.right.PrintPos()
	}
	fmt.Print(bstNode.value,", ")
  }

  func (bstNode *BstNode) Remove(value int) *BstNode {
	if value < bstNode.value{
	  bstNode.left = bstNode.left.Remove(value)
	} else if value > bstNode.value{
	  bstNode.right = bstNode.right.Remove(value)
	} else {   //encontramos o nó a ser removido    
	  if bstNode.left == nil && bstNode.right == nil{          //caso 1: nó folha
		return nil      
	  } else if bstNode.left != nil && bstNode.right == nil {  //caso 2: nó com 1 filho (à esquerda)
		return bstNode.left
	  } else if bstNode.left == nil && bstNode.right != nil {  //caso 2: nó com 1 filho (à direita)
		return bstNode.right
	  } else {                                                  //caso 3: nó com 2 filhos
		maxEsq := bstNode.left.Max()
		bstNode.value = maxEsq
		bstNode.left = bstNode.left.Remove(maxEsq)
		return bstNode
	  }
	}
	return bstNode
  }
