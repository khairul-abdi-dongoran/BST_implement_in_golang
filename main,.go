package main

import "fmt"

const (
	COUNT int = 5
)

type node struct {
	left  *node
	data  int
	right *node
}
type bst struct {
	root *node
}

type BSTService interface {
	Insert(data int) *bst
	GetRoot() *node
	Search(key int) bool
}

func (tree *bst) Insert(data int) *bst {
	if tree.root == nil {
		tree.root = &node{
			data: data,
		}
	} else {
		tree.root.insert(data)
	}
	return tree
}
func (nd *node) insert(data int) {
	if nd == nil {
		return
	} else if data < nd.data {
		if nd.left == nil {
			nd.left = &node{
				data: data,
			}
		} else {
			nd.left.insert(data)
		}
	} else {
		if nd.right == nil {
			nd.right = &node{
				data: data,
			}
		} else {
			nd.right.insert(data)
		}
	}
}
func (tree *bst) Search(key int) bool {
	var keyFound bool
	if tree.root == nil {
		keyFound = false
	} else if tree.root.data != key {
		keyFound = tree.root.search(key)
	} else {
		keyFound = true
	}
	return keyFound
}
func (nd *node) search(key int) bool {
	var keyFound bool
	if nd == nil {
		keyFound = false
	} else if key == nd.data {
		keyFound = true
	} else {
		if nd.data > key {
			keyFound = nd.left.search(key)
		} else {
			keyFound = nd.right.search(key)
		}
	}
	return keyFound
}
func (tree *bst) GetRoot() *node {
	return tree.root
}

func preOrderTraversal(root *node) {
	if root == nil {
		return
	}
	fmt.Print(root.data, " ")
	preOrderTraversal(root.left)
	preOrderTraversal(root.right)
}
func inOrderTraversal(root *node) {
	if root == nil {
		return
	}
	inOrderTraversal(root.left)
	fmt.Print(root.data, " ")
	inOrderTraversal(root.right)
}
func postOrderTraversal(root *node) {
	if root == nil {
		return
	}
	postOrderTraversal(root.left)
	postOrderTraversal(root.right)
	fmt.Print(root.data, " ")
}
func printTree(root *node, space int) {
	if root == nil {
		return
	}
	space += COUNT
	printTree(root.right, space)
	for i := COUNT; i < space; i++ {
		fmt.Printf(":")
	}
	fmt.Printf("[")
	fmt.Printf("%v", root.data)
	fmt.Println("]")
	printTree(root.left, space)
}

func NewBST() BSTService {
	return &bst{}
}

func main() {
	tree := NewBST()
	tree.Insert(10).
		Insert(5).
		Insert(15).
		Insert(4).
		Insert(16).
		Insert(6).
		Insert(13).
		Insert(1).
		Insert(12).
		Insert(2).
		Insert(20) // Input Data : 10, 5, 15, 4, 16, 6, 13, 1, 12, 2, 20

	printTree(tree.GetRoot(), 0)
	fmt.Printf("Pre-Order: ")
	preOrderTraversal(tree.GetRoot())
	fmt.Printf("\nIn-Order: ")
	inOrderTraversal(tree.GetRoot())
	fmt.Printf("\nPost-Order: ")
	postOrderTraversal(tree.GetRoot())
	fmt.Println()
	fmt.Println("20 Exists:", tree.Search(20))
	fmt.Println("21 Exists:", tree.Search(21))
	tree.Insert(50)
}

/* Output
:::::::::::::::[20]
::::::::::[16]
:::::[15]
::::::::::[13]
:::::::::::::::[12]
[10]
::::::::::[6]
:::::[5]
::::::::::[4]
::::::::::::::::::::[2]
:::::::::::::::[1]
Pre-Order: 10 5 4 1 2 6 15 13 12 16 20
In-Order: 1 2 4 5 6 10 12 13 15 16 20
Post-Order: 2 1 4 6 5 12 13 20 16 15 10
20 Exists: true
21 Exists: false
*/
