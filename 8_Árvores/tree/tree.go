package tree

import (
	"fmt"
	"sort"
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

  func (bstNode *BstNode) PrintLevels() {
		fila := []*BstNode{}
		fila = append(fila, bstNode)

		for len(fila) > 0 {
    		atual := fila[0] //separa o primeiro nó numa var
    		fila = fila[1:]
			fmt.Print(atual.value, ", ")
   
   			// Insere os filhos do último nó na fila
   			if atual.left != nil {
   				fila = append(fila, atual.left)
   			}
   			if atual.right != nil {
   			 	fila = append(fila, atual.right)
   			}
			
		}
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

  func (bstNode *BstNode) isBST() bool {
    if bstNode.left != nil{
        if bstNode.left.value < bstNode.value {
            return bstNode.left.isBST()
        } else{
            return false
        }
    }
    if bstNode.right != nil{
        if bstNode.right.value > bstNode.value {
            return bstNode.right.isBST()
        } else{
            return false
        }
    }
    return true
 }
 
func (bstNode *BstNode) Size() int {
    /*if bstNode == nil{
        return 0
    }*/
    contador := 1 
    if bstNode.left != nil{
        contador = contador + bstNode.left.Size()
    }
    if bstNode.right != nil{
        contador = contador + bstNode.right.Size()
    }
    return contador
}

func createBst(v []int, start int, end int) *BstNode {
    if start > end {
        return nil
    }
    sort.Ints(v)
    
    mid := (start + end)/2
    raiz := NewNode(v[mid])

    raiz.left = createBst(v,start,mid - 1)
    raiz.right = createBst(v,mid+1,end)

    return raiz
}

