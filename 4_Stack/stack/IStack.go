package stack

type IStack interface {
	Init(size int)
	Push(value int)
	Pop() (int, error)
	Peek() (int, error)
	Empty() (bool)
	Size() (int)
}