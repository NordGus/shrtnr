package trie

type node struct {
	value    string
	children map[string]*node
	end      bool
}

func newNode(val string) *node {
	return &node{
		value:    val,
		children: make(map[string]*node),
		end:      false,
	}
}

func (n *node) hasChild(val string) bool {
	_, ok := n.children[val]

	return ok
}

func (n *node) hasChildren() bool {
	return len(n.children) > 0
}

func (n *node) getChild(val string) *node {
	return n.children[val]
}

func (n *node) addChild(val string) {
	n.children[val] = newNode(val)
}

func (n *node) removeChild(val string) {
	delete(n.children, val)
}
