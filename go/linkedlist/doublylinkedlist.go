package linkedlist

type dllNode struct {
	value int
	next  *dllNode
	prev  *dllNode
}

type dLinkedList struct {
	head *dllNode
	size int
}

func NewDlinkedlist() *dLinkedList {
	return &dLinkedList{nil, 0}
}

// TODO: Display
// TODO: Reverse
// TODO: Append
// TODO: Insert
// TODO: Add at position
// TODO: Delete Root
// TODO: Delete at position
