package avlTree

import "fmt"

type Node struct {
	height      int
	left, right *Node

	data int
	key  int
}

type AVLTree struct {
	root *Node
}

func (t *AVLTree) Add(key, data int) {
	t.root = t.root.add(key, data)
}

func (t *AVLTree) Remove(key int) {
	t.root = t.root.remove(key)
}

func (t *AVLTree) Update(oldKey, newKey int, newData int) {
	t.root = t.root.remove(oldKey)
	t.root = t.root.add(newKey, newData)
}

func (t *AVLTree) Search(key int) (node *Node) {
	return t.root.search(key)
}

func (t *AVLTree) DisplayInorder() {
	t.root.display()
}

func (node *Node) display() {
	if node == nil {
		return
	}
	node.left.display()
	fmt.Printf("[%v, %v] ", node.key, node.data)
	node.right.display()
}

func (node *Node) recalculateHeight() {
	node.height = 1 + max(node.left.getHeight(), node.right.getHeight())
}

func (node *Node) getHeight() int {
	if node == nil {
		return 0
	}
	return node.height
}

func (node *Node) search(key int) *Node {
	if node == nil {
		return nil
	}
	if key < node.key {
		return node.left.search(key)
	}
	if key > node.key {
		return node.right.search(key)
	}
	return node
}

func (node *Node) add(key int, data int) *Node {
	if node == nil {
		return &Node{
			key:    key,
			data:   data,
			height: 1,
			left:   nil,
			right:  nil,
		}
	}

	if key < node.key {
		node.left = node.left.add(key, data)
	} else if key > node.key {
		node.right = node.right.add(key, data)
	} else {
		node.data = data
	}
	return node.rebalance()
}

func (node *Node) remove(key int) *Node {
	if node == nil {
		return nil
	}
	if key < node.key {
		node.left = node.left.remove(key)
	} else if key > node.key {
		node.right = node.right.remove(key)
	} else {
		if node.left != nil && node.right != nil {
			rightMin := node.right.getSmallest()
			node.key = rightMin.key
			node.data = rightMin.data
			node.right = node.right.remove(rightMin.key)
		} else if node.left != nil {
			node = node.left
		} else if node.right != nil {
			node = node.right
		} else {
			node = nil
			return node
		}
	}
	return node.rebalance()
}

func (node *Node) rebalance() *Node {
	if node == nil {
		return node
	}
	node.recalculateHeight()
	bf := node.left.getHeight() - node.right.getHeight()
	if bf == -2 {
		if node.right.left.getHeight() > node.right.right.getHeight() {
			node.right = node.right.rightRotate()
		}
		return node.leftRotate()
	} else if bf == 2 {
		if node.left.right.getHeight() > node.left.left.getHeight() {
			node.left = node.left.leftRotate()
		}
		return node.rightRotate()
	}
	return node
}

func (node *Node) leftRotate() *Node {
	r := node.right
	node.right = r.left
	r.left = node
	node.recalculateHeight()
	r.recalculateHeight()
	return r
}

func (node *Node) rightRotate() *Node {
	l := node.left
	node.left = l.right
	l.right = node
	node.recalculateHeight()
	l.recalculateHeight()
	return l
}

func (node *Node) getSmallest() *Node {
	if node.left != nil {
		return node.left.getSmallest()
	} else {
		return node
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
