package main

import (
	"fmt"
)

func main() {
	link1 := &List{}
	ListRemoveIf(link1, 1)
	PrintList(link1)

	link2 := &List{}
	ListRemoveIf(link2, 96)
	PrintList(link2)

	link3 := &List{}
	ListPushBack(link3, 98)
	ListPushBack(link3, 98)
	ListPushBack(link3, 33)
	ListPushBack(link3, 34)
	ListPushBack(link3, 33)
	ListPushBack(link3, 34)
	ListPushBack(link3, 33)
	ListPushBack(link3, 89)
	ListPushBack(link3, 33)
	ListRemoveIf(link3, 34)
	PrintList(link3)

	link4 := &List{}
	ListPushBack(link4, 79)
	ListPushBack(link4, 74)
	ListPushBack(link4, 99)
	ListPushBack(link4, 79)
	ListPushBack(link4, 7)
	ListRemoveIf(link4, 99)
	PrintList(link4)

	link5 := &List{}
	ListPushBack(link5, 56)
	ListPushBack(link5, 93)
	ListPushBack(link5, 68)
	ListPushBack(link5, 56)
	ListPushBack(link5, 87)
	ListPushBack(link5, 68)
	ListPushBack(link5, 56)
	ListPushBack(link5, 68)
	ListRemoveIf(link5, 68)
	PrintList(link5)

	link6 := &List{}
	ListPushBack(link6, "mvkUxbqhQve4l")
	ListPushBack(link6, "4Zc4t hnf SQ")
	ListPushBack(link6, "q2If E8BPuX")
	ListRemoveIf(link6, "4Zc4t hnf SQ")
	PrintList(link6)
}

func ListRemoveIf(l *List, data_ref interface{}) {
	if l == nil {
		return
	}
	l2 := &List{}
	for l.Head != nil {
		if l.Head.Data != data_ref {
			ListPushBack(l2, l.Head.Data)
		}
		l.Head = l.Head.Next
	}
	*l = *l2
}

func PrintList(l *List) {
	it := l.Head
	for it != nil {
		fmt.Print(it.Data, " -> ")
		it = it.Next
	}

	fmt.Print(nil, "\n")
}

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListPushBack(l *List, data interface{}) {
	node := &NodeL{Data: data, Next: nil}
	if l.Head == nil {
		l.Head = node
		l.Tail = node

		return
	}
	l.Tail.Next = node
	l.Tail = node

}
