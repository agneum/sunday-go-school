package main

type Tree struct {
	root *Node
}

type Node struct {
	left  *Node
	value int
	right *Node
}

func (t *Tree) insert(value int) error {
	if t.root == nil {
		t.root = &Node{value: value}
		return nil
	}
	return t.root.insert(value)
}

func (t *Tree) exists(value int) bool {
	if t.root == nil {
		return false
	}
	return t.root.exists(value)
}

func (n *Node) insert(value int) error {
	if n.value == value {
		return nil
	}

	if n.value > value {
		if n.left == nil {
			n.left = &Node{value: value}
			return nil
		}
		return n.left.insert(value)
	}

	if n.value < value {
		if n.right == nil {
			n.right = &Node{nil, value, nil}
			return nil
		}
		return n.right.insert(value)
	}

	return nil
}

func (n *Node) exists(value int) bool {
	if n == nil {
		return false
	}
	if n.value == value {
		return true
	}
	if n.value > value {
		return n.left.exists(value)
	}
	return n.right.exists(value)
}
