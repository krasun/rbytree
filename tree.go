package rbytree

import (
	"bytes"
)

// Tree holds red-black tree.
// It is not goroutine-safe, make sure that
// the access to the instance of the tree is always synchronized.
type Tree struct {
	root *node
	size int
}

type color byte

const (
	red color = iota
	black
)

// node represents the node in the tree.
type node struct {
	key    []byte
	value  []byte
	parent *node
	left   *node
	right  *node
	color  color
}

// New creates new empty instance of Red-black tree.
func New() *Tree {
	return &Tree{}
}

// Put inserts the key with the associated value into the tree.
func (t *Tree) Put(key []byte, value []byte) {
	// too guarantee that the invariants are not violated
	key = copyBytes(key)

	newNode := &node{key, value, nil, nil, nil, red}
	if t.root == nil {
		newNode.color = black
		t.root = newNode
		t.size = 1

		return
	}

	current := t.root
	var parent *node
	var cmp int
	for current != nil {
		parent = current

		cmp = bytes.Compare(key, current.key)
		if cmp == 0 {
			current.value = value
			return
		}

		if cmp < 0 {
			current = current.left
		} else {
			current = current.right
		}
	}

	if cmp < 0 {
		parent.left = newNode
	} else {
		parent.right = newNode
	}
	newNode.parent = parent

	t.fixAfterInsertion(newNode)

	t.size++
}

// Get searches the key and returns the associated value and true if found,
// otherwise nil and false.
func (t *Tree) Get(key []byte) ([]byte, bool) {
	if t.root == nil {
		return nil, false
	}

	current := t.root
	for current != nil {
		cmp := bytes.Compare(key, current.key)
		if cmp < 0 {
			current = current.left
		} else if cmp > 0 {
			current = current.right
		} else {
			return current.value, true
		}
	}

	return nil, false
}

// ForEach traverses tree in ascending order.
// The function uses Morris in-order traversal
// (https://en.wikipedia.org/wiki/Tree_traversal#Morris_in-order_traversal_using_threading).
func (t *Tree) ForEach(action func(key []byte, value []byte)) {
	current := t.root
	for current != nil {
		if current.left == nil {
			action(current.key, current.value)
			current = current.right
		} else {
			// find theinorder predecessor of the current
			predecessor := current.left
			for predecessor.right != nil && predecessor.right != current {
				predecessor = predecessor.right
			}

			if predecessor.right == nil {
				// make the current as the right
				// child of its inorder predecessor
				predecessor.right = current
				current = current.left
			} else {
				// revert the changes
				predecessor.right = nil
				action(current.key, current.value)
				current = current.right
			}
		}
	}
}

// fixAfterInsertion fixes the tree to satisfy the red-black tree
// properties of the tree.
func (t *Tree) fixAfterInsertion(newNode *node) {
	current := newNode

	for current != t.root && current.parent.color == red {
		if current.parent.parent.left == current.parent {
			uncle := current.parent.parent.right
			if uncle != nil && uncle.color == red {
				current.parent.color = black
				uncle.color = black
				current.parent.parent.color = red

				current = current.parent.parent
			} else {
				if current == current.parent.right {
					current = current.parent

					t.rotateLeft(current)
				}

				current.parent.color = black
				current.parent.parent.color = red

				t.rotateRight(current.parent.parent)
			}
		} else if current.parent.parent.right == current.parent {
			uncle := current.parent.parent.left
			if uncle != nil && uncle.color == red {
				current.parent.color = black
				uncle.color = black
				current.parent.parent.color = red
				current = current.parent.parent
			} else {
				if current == current.parent.left {
					current = current.parent

					t.rotateRight(current)
				}

				current.parent.color = black
				current.parent.parent.color = red

				t.rotateLeft(current.parent.parent)
			}
		}
	}

	t.root.color = black
}

func (t *Tree) rotateLeft(node *node) {
	nodeRight := node.right
	node.right = nodeRight.left

	if nodeRight.left != nil {
		nodeRight.left.parent = node
	}
	nodeRight.parent = node.parent

	if node.parent == nil {
		t.root = nodeRight
	} else if node == node.parent.left {
		node.parent.left = nodeRight
	} else if node == node.parent.right {
		node.parent.right = nodeRight
	}

	nodeRight.left = node
	node.parent = nodeRight
}

func (t *Tree) rotateRight(node *node) {
	nodeLeft := node.left
	node.left = nodeLeft.right

	if nodeLeft.right != nil {
		nodeLeft.right.parent = node
	}

	nodeLeft.parent = node.parent
	if node.parent == nil {
		t.root = nodeLeft
	} else if node == node.parent.left {
		node.parent.left = nodeLeft
	} else if node == node.parent.right {
		node.parent.right = nodeLeft
	}

	nodeLeft.right = node
	node.parent = nodeLeft
}

// Size returns tree size.
func (t *Tree) Size() int {
	return t.size
}

func copyBytes(s []byte) []byte {
	c := make([]byte, len(s))
	copy(c, s)

	return c
}
