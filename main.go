package main

import (
	linkedlist "linkedlist/linkedlist"
)

func main() {

	list := linkedlist.Init()
	list.Append(1)
	list.PrintList()
	list.Append("hello")
	list.PrintList()
}
