package main

// func main() {
// 	root := &TreeNode{Data: "01"}
// 	BTreeInsertData(root, "07")
// 	BTreeInsertData(root, "05")
// 	BTreeInsertData(root, "12")
// 	BTreeInsertData(root, "02")
// 	BTreeInsertData(root, "10")
// 	BTreeInsertData(root, "03")
// 	node := BTreeSearchItem(root, "12")
// 	replacement := &TreeNode{Data: "55"}
// 	BTreeInsertData(replacement, "33")
// 	BTreeInsertData(replacement, "60")
// 	BTreeInsertData(replacement, "12")
// 	BTreeInsertData(replacement, "15")
// 	root = BTreeTransplant(root, node, replacement)
// 	BTreeApplyInorder(root, fmt.Println)
// }

// 	root := &TreeNode{Data: "33"}
// 	BTreeInsertData(root, "20")
// 	BTreeInsertData(root, "5")
// 	BTreeInsertData(root, "13")
// 	BTreeInsertData(root, "31")
// 	BTreeInsertData(root, "52")
// 	BTreeInsertData(root, "11")
// 	node := BTreeSearchItem(root, "20")
// 	replacement := &TreeNode{Data: "55"}
// 	BTreeInsertData(replacement, "60")
// 	BTreeInsertData(replacement, "33")
// 	BTreeInsertData(replacement, "12")
// 	BTreeInsertData(replacement, "15")
// 	root = BTreeTransplant(root, node, replacement)
// 	BTreeApplyInorder(root, fmt.Println)
// }

// 	root := &TreeNode{Data: "03"}
// 	BTreeInsertData(root, "39")
// 	BTreeInsertData(root, "11")
// 	BTreeInsertData(root, "99")
// 	BTreeInsertData(root, "14")
// 	BTreeInsertData(root, "44")
// 	BTreeInsertData(root, "11")
// 	node := BTreeSearchItem(root, "11")
// 	replacement := &TreeNode{Data: "55"}
// 	BTreeInsertData(replacement, "33")
// 	BTreeInsertData(replacement, "60")
// 	BTreeInsertData(replacement, "12")
// 	BTreeInsertData(replacement, "15")
// 	root = BTreeTransplant(root, node, replacement)
// 	BTreeApplyInorder(root, fmt.Println)
// }

type TreeNode struct {
	Left, Right, Parent *TreeNode
	Data                string
}

func BTreeTransplant(root, node, rplc *TreeNode) *TreeNode {
	if node == nil {
		return root
	}
	if node.Parent.Left.Data == node.Data {
		node.Parent.Left = rplc
	} else if node.Parent.Right.Data == node.Data {
		node.Parent.Right = rplc
	}
	return root
}

func BTreeInsertData(root *TreeNode, data string) {
	if data < root.Data {
		if root.Left == nil {
			root.Left = &TreeNode{Data: data, Parent: root}
		} else {
			BTreeInsertData(root.Left, data)
		}
	} else {
		if root.Right == nil {
			root.Right = &TreeNode{Data: data, Parent: root}
		} else {
			BTreeInsertData(root.Right, data)
		}
	}
}

func BTreeSearchItem(root *TreeNode, data string) *TreeNode {

	if root != nil {
		if data < root.Data {
			return BTreeSearchItem(root.Left, data)
		}

		if data > root.Data {
			return BTreeSearchItem(root.Right, data)
		}
	}

	return root
}

func BTreeApplyInorder(root *TreeNode, f func(...interface{}) (int, error)) {
	if root != nil {
		BTreeApplyInorder(root.Left, f)
		f(root.Data)
		BTreeApplyInorder(root.Right, f)
	}
}
