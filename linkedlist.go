package linkedlist

import (
	"fmt"
)

type Node struct {
	Data interface{}
	Next *Node
}

func Init() *Node {
	return &Node{}
}

func (list *Node) Append(data interface{}) {
	if list.Data == nil {
		list.Data = data
		return
	}
	current := list
	for current.Next != nil {
		current = current.Next
	}
	current.Next = &Node{Data: data}
}

func (list *Node) PrintList() {
	current := list
	for current != nil {
		fmt.Println(current.Data)
		current = current.Next
	}
}
