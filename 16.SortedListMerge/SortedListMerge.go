package main

import (
	"fmt"
)

func main() {
	var link1 *NodeI
	var link11 *NodeI
	PrintList(SortedListMerge(link1, link11))

	var link2 *NodeI
	var link22 *NodeI
	link2 = listPushBack(link2, 2)
	link2 = listPushBack(link2, 2)
	link2 = listPushBack(link2, 4)
	link2 = listPushBack(link2, 9)
	link2 = listPushBack(link2, 12)
	link2 = listPushBack(link2, 12)
	link2 = listPushBack(link2, 19)
	link2 = listPushBack(link2, 20)
	PrintList(SortedListMerge(link2, link22))

	var link3 *NodeI
	var link33 *NodeI
	link33 = listPushBack(link33, 4)
	link33 = listPushBack(link33, 4)
	link33 = listPushBack(link33, 6)
	link33 = listPushBack(link33, 9)
	link33 = listPushBack(link33, 13)
	link33 = listPushBack(link33, 18)
	link33 = listPushBack(link33, 20)
	link33 = listPushBack(link33, 20)
	PrintList(SortedListMerge(link33, link3))

	var link4 *NodeI
	var link44 *NodeI
	link4 = listPushBack(link4, 0)
	link4 = listPushBack(link4, 7)
	link4 = listPushBack(link4, 39)
	link4 = listPushBack(link4, 92)
	link4 = listPushBack(link4, 97)
	link4 = listPushBack(link4, 93)
	link4 = listPushBack(link4, 91)
	link4 = listPushBack(link4, 28)
	link4 = listPushBack(link4, 64)
	link44 = listPushBack(link44, 80)
	link44 = listPushBack(link44, 23)
	link44 = listPushBack(link44, 27)
	link44 = listPushBack(link44, 30)
	link44 = listPushBack(link44, 85)
	link44 = listPushBack(link44, 81)
	link44 = listPushBack(link44, 75)
	link44 = listPushBack(link44, 70)
	PrintList(SortedListMerge(link4, link44))

	var link5 *NodeI
	var link55 *NodeI
	link5 = listPushBack(link5, 0)
	link5 = listPushBack(link5, 2)
	link5 = listPushBack(link5, 11)
	link5 = listPushBack(link5, 30)
	link5 = listPushBack(link5, 54)
	link5 = listPushBack(link5, 56)
	link5 = listPushBack(link5, 70)
	link5 = listPushBack(link5, 79)
	link5 = listPushBack(link5, 99)
	link55 = listPushBack(link55, 23)
	link55 = listPushBack(link55, 28)
	link55 = listPushBack(link55, 38)
	link55 = listPushBack(link55, 67)
	link55 = listPushBack(link55, 67)
	link55 = listPushBack(link55, 79)
	link55 = listPushBack(link55, 95)
	link55 = listPushBack(link55, 97)
	PrintList(SortedListMerge(link5, link55))

	var link6 *NodeI
	var link66 *NodeI
	link6 = listPushBack(link6, 0)
	link6 = listPushBack(link6, 3)
	link6 = listPushBack(link6, 8)
	link6 = listPushBack(link6, 8)
	link6 = listPushBack(link6, 13)
	link6 = listPushBack(link6, 19)
	link6 = listPushBack(link6, 34)
	link6 = listPushBack(link6, 38)
	link6 = listPushBack(link6, 46)
	link66 = listPushBack(link66, 7)
	link66 = listPushBack(link66, 39)
	link66 = listPushBack(link66, 45)
	link66 = listPushBack(link66, 53)
	link66 = listPushBack(link66, 59)
	link66 = listPushBack(link66, 70)
	link66 = listPushBack(link66, 76)
	link66 = listPushBack(link66, 79)
	PrintList(SortedListMerge(link6, link66))
}

type NodeI struct {
	Data int
	Next *NodeI
}

func SortedListMerge(n1 *NodeI, n2 *NodeI) *NodeI {
	n1 = ListSort(n1)
	n2 = ListSort(n2)

	if n1 == nil {
		return n2
	}
	if n2 == nil {
		return n1
	}
	if n1.Data <= n2.Data {
		n1.Next = SortedListMerge(n1.Next, n2)
		return n1
	} else {
		n2.Next = SortedListMerge(n1, n2.Next)
		return n2
	}
}

func ListSort(l *NodeI) *NodeI {

	if l != nil {
		NewH := l
		for NewH.Next != nil {
			NewH2 := l
			for NewH2.Next != nil {
				if NewH2.Next.Data < NewH2.Data {
					NewH2.Data, NewH2.Next.Data = NewH2.Next.Data, NewH2.Data
				}
				NewH2 = NewH2.Next
			}
			NewH = NewH.Next
		}
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
