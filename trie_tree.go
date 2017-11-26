package trie

import (
	"sync"
	"fmt"
)

type TrieNode struct {
	value    rune
	is_leaf  bool
	children []*TrieNode
}

type TrieTree struct {
	sync.Mutex
	root *TrieNode
}

func NewTrieTree() *TrieTree {
	return &TrieTree{
		root: nil,
	}
}

func newNode() *TrieNode {
	return &TrieNode{
		is_leaf: false,
		children: make([]*TrieNode, 0),
	}
}

func (n *TrieNode) isLeaf() bool {
	return n.is_leaf
}

func (n *TrieNode) find(v rune) *TrieNode {
	for _, c := range n.children {
		if c.value == v {
			return c
		}
	}
	return nil
}

func (n *TrieNode) insert(child *TrieNode, v rune) {
	if n.find(v) != nil {
		return
	}

	n.children = append(n.children, child)
}

func (n *TrieNode) remove(child *TrieNode) {
	for i, c := range n.children {
		if c == child {
			n.children = append(n.children[:i], n.children[i+1:]...)
		}
	}
}

func (n *TrieNode) dump() {
	for _, c := range n.children {
		if c.isLeaf() {
			fmt.Printf("%c[1;40;32m%c%c[0m\n", 0x1B, c.value, 0x1B)
		} else {
			fmt.Printf("%c\n", c.value)
		}
		c.dump()
	}
}

func (tt *TrieTree) Add(entry string) {
	tt.Lock()
	defer tt.Unlock()

	if tt.root == nil {
		tt.root = newNode()
	}

	node := tt.root

	for i, c := range entry {
		n := node.find(c)
		if n == nil {
			n = newNode()
			n.value = c
			node.insert(n, c)
		}
		if i == len(entry) - 1 {
			n.is_leaf = true
		}
		node = n
	}
}

func (tt *TrieTree) Has(entry string) bool {
	tt.Lock()
	defer tt.Unlock()

	node := tt.root

	for i, c := range entry {
		n := node.find(c)
		if n == nil {
			return false
		}
		if i == len(entry) - 1 {
			return n.isLeaf()
		}
		node = n
	}
	return false
}

func (tt *TrieTree) Delete(entry string) bool {
	tt.Lock()
	defer tt.Unlock()

	path := make([]*TrieNode, 0)
	node := tt.root

	for _, c := range entry {
		n := node.find(c)
		if n == nil {
			return false
		}
		path = append(path, n)
		node = n
	}

	for i := range path {
		level := len(path) - i - 1
		n := path[level]
		// root node
		if level == 0 {
			if len(n.children) < 1 {
				tt.root = nil
			}
		} else {
			parent := path[level - 1]
			if len(n.children) < 1 {
				parent.remove(n)
			}
		}
	}
	return true
}

func (tt *TrieTree) Dump() {
	tt.Lock()
	defer tt.Unlock()

	if tt.root != nil {
		tt.root.dump()
	}
}
