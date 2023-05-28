package tree

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

