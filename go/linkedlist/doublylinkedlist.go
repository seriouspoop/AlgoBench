package linkedlist

type dllNode struct {
	value int
	next  *dllNode
}

// TODO hellow mein hu
type dLinkedList struct {
	head *dllNode
	size int
}

func NewDlinkedlist() *dLinkedList {
	return &dLinkedList{nil, 0}
}

// TODO:  implement doubly linked list functionalities
