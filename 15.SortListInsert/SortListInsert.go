package main

import (
	"fmt"
)

func main() {

	var link1 *NodeI
	var link2 *NodeI
	var link3 *NodeI
	var link4 *NodeI
	var link5 *NodeI

	link1 = listPushBack(link1, 0)
	link1 = SortListInsert(link1, 39)
	PrintList(link1)

	link2 = listPushBack(link2, 0)
	link2 = listPushBack(link2, 1)
	link2 = listPushBack(link2, 2)
	link2 = listPushBack(link2, 3)
	link2 = listPushBack(link2, 4)
	link2 = listPushBack(link2, 5)
	link2 = listPushBack(link2, 24)
	link2 = listPushBack(link2, 25)
	link2 = listPushBack(link2, 54)
	link2 = SortListInsert(link2, 33)
	PrintList(link2)

	link3 = listPushBack(link3, 0)
	link3 = listPushBack(link3, 2)
	link3 = listPushBack(link3, 18)
	link3 = listPushBack(link3, 33)
	link3 = listPushBack(link3, 37)
	link3 = listPushBack(link3, 37)
	link3 = listPushBack(link3, 39)
	link3 = listPushBack(link3, 52)
	link3 = listPushBack(link3, 53)
	link3 = listPushBack(link3, 57)
	link3 = SortListInsert(link3, 53)
	PrintList(link3)

	link4 = listPushBack(link4, 0)
	link4 = listPushBack(link4, 5)
	link4 = listPushBack(link4, 18)
	link4 = listPushBack(link4, 24)
	link4 = listPushBack(link4, 28)
	link4 = listPushBack(link4, 35)
	link4 = listPushBack(link4, 42)
	link4 = listPushBack(link4, 45)
	link4 = listPushBack(link4, 52)
	link4 = SortListInsert(link4, 52)
	PrintList(link4)

	link5 = listPushBack(link5, 0)
	link5 = listPushBack(link5, 12)
	link5 = listPushBack(link5, 20)
	link5 = listPushBack(link5, 23)
	link5 = listPushBack(link5, 23)
	link5 = listPushBack(link5, 24)
	link5 = listPushBack(link5, 30)
	link5 = listPushBack(link5, 41)
	link5 = listPushBack(link5, 53)
	link5 = listPushBack(link5, 57)
	link5 = listPushBack(link5, 59)
	link5 = SortListInsert(link5, 38)
	PrintList(link5)

}

type NodeI struct {
	Data int
	Next *NodeI
}

func SortListInsert(l *NodeI, data_ref int) *NodeI {
	node := &NodeI{}
	node.Data = data_ref
	node.Next = nil

	if l == nil || l.Data > node.Data {
		node.Next = l
		return node
	} else {
		a := l
		for a.Next != nil && a.Next.Data <= node.Data {
			a = a.Next
		}

		node.Next = a.Next
		a.Next = node
	}
	return l
}

func PrintList(l *NodeI) {
	it := l
	for it != nil {
		fmt.Print(it.Data, " -> ")
		it = it.Next
	}
	fmt.Print(nil, "\n")
}

func listPushBack(l *NodeI, data int) *NodeI {
	n := &NodeI{Data: data}

	if l == nil {
		return n
	}
	iterator := l
	for iterator.Next != nil {
		iterator = iterator.Next
	}
	iterator.Next = n
	return l
}
