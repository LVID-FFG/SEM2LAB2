package main

// ==================== КРАСНО-ЧЕРНОЕ ДЕРЕВО ====================

type Color int

const (
	RED Color = iota
	BLACK
)

type RBNode struct {
	data   string
	color  Color
	parent *RBNode
	left   *RBNode
	right  *RBNode
}

func NewRBNode(value string) *RBNode {
	return &RBNode{
		data:   value,
		color:  RED,
		parent: nil,
		left:   nil,
		right:  nil,
	}
}

type RBTree struct {
	root *RBNode
}

func NewRBTree() *RBTree {
	return &RBTree{root: nil}
}

// ==================== ФУНКЦИИ КРАСНО-ЧЕРНОГО ДЕРЕВА ====================

func leftRotate(tree *RBTree, x *RBNode) {
	y := x.right
	x.right = y.left
	if y.left != nil {
		y.left.parent = x
	}
	y.parent = x.parent
	if x.parent == nil {
		tree.root = y
	} else {
		if x == x.parent.left {
			x.parent.left = y
		} else {
			x.parent.right = y
		}
	}
	y.left = x
	x.parent = y
}

func rightRotate(tree *RBTree, x *RBNode) {
	y := x.left
	x.left = y.right
	if y.right != nil {
		y.right.parent = x
	}
	y.parent = x.parent
	if x.parent == nil {
		tree.root = y
	} else {
		if x == x.parent.left {
			x.parent.left = y
		} else {
			x.parent.right = y
		}
	}
	y.right = x
	x.parent = y
}

func fixAdd(tree *RBTree, z *RBNode) {
	for z.parent != nil && z.parent.color == RED {
		if z.parent.parent == nil {
			break
		}

		if z.parent == z.parent.parent.left {
			y := z.parent.parent.right

			if y != nil && y.color == RED {
				z.parent.color = BLACK
				y.color = BLACK
				z.parent.parent.color = RED
				z = z.parent.parent
			} else {
				if z == z.parent.right {
					z = z.parent
					leftRotate(tree, z)
				}
				z.parent.color = BLACK
				z.parent.parent.color = RED
				rightRotate(tree, z.parent.parent)
			}
		} else {
			y := z.parent.parent.left

			if y != nil && y.color == RED {
				z.parent.color = BLACK
				y.color = BLACK
				z.parent.parent.color = RED
				z = z.parent.parent
			} else {
				if z == z.parent.left {
					z = z.parent
					rightRotate(tree, z)
				}
				z.parent.color = BLACK
				z.parent.parent.color = RED
				leftRotate(tree, z.parent.parent)
			}
		}

		if z == tree.root {
			break
		}
	}
	tree.root.color = BLACK
}

func AddRBNode(tree *RBTree, value string) {
	newNode := NewRBNode(value)
	if tree.root == nil {
		newNode.color = BLACK
		tree.root = newNode
		return
	}
	address := tree.root
	for {
		if value < address.data {
			if address.left == nil {
				newNode.parent = address
				address.left = newNode
				fixAdd(tree, newNode)
				return
			} else {
				address = address.left
			}
		} else {
			if address.right == nil {
				newNode.parent = address
				address.right = newNode
				fixAdd(tree, newNode)
				return
			} else {
				address = address.right
			}
		}
	}
}

func GetRBNode(tree *RBTree, value string) *RBNode {
	if tree.root == nil {
		return nil
	}

	address := tree.root
	for address != nil {
		if value == address.data {
			return address
		} else if value < address.data {
			address = address.left
		} else {
			address = address.right
		}
	}

	return nil
}

func treeMinimum(node *RBNode) *RBNode {
	if node == nil {
		return nil
	}
	address := node
	for address.left != nil {
		address = address.left
	}
	return address
}

func deleteFix(tree *RBTree, x *RBNode) {
	if x == nil {
		return
	}

	for x != tree.root && x.color == BLACK {
		if x == x.parent.left {
			w := x.parent.right
			if w == nil {
				break
			}

			if w.color == RED {
				w.color = BLACK
				x.parent.color = RED
				leftRotate(tree, x.parent)
				w = x.parent.right
				if w == nil {
					break
				}
			}

			leftBlack := (w.left == nil) || (w.left.color == BLACK)
			rightBlack := (w.right == nil) || (w.right.color == BLACK)

			if leftBlack && rightBlack {
				w.color = RED
				x = x.parent
			} else {
				if w.right == nil || w.right.color == BLACK {
					if w.left != nil {
						w.left.color = BLACK
					}
					w.color = RED
					rightRotate(tree, w)
					w = x.parent.right
					if w == nil {
						break
					}
				}

				w.color = x.parent.color
				x.parent.color = BLACK
				if w.right != nil {
					w.right.color = BLACK
				}
				leftRotate(tree, x.parent)
				x = tree.root
			}
		} else {
			w := x.parent.left
			if w == nil {
				break
			}

			if w.color == RED {
				w.color = BLACK
				x.parent.color = RED
				rightRotate(tree, x.parent)
				w = x.parent.left
				if w == nil {
					break
				}
			}

			leftBlack := (w.left == nil) || (w.left.color == BLACK)
			rightBlack := (w.right == nil) || (w.right.color == BLACK)

			if leftBlack && rightBlack {
				w.color = RED
				x = x.parent
			} else {
				if w.left == nil || w.left.color == BLACK {
					if w.right != nil {
						w.right.color = BLACK
					}
					w.color = RED
					leftRotate(tree, w)
					w = x.parent.left
					if w == nil {
						break
					}
				}

				w.color = x.parent.color
				x.parent.color = BLACK
				if w.left != nil {
					w.left.color = BLACK
				}
				rightRotate(tree, x.parent)
				x = tree.root
			}
		}
	}

	if x != nil {
		x.color = BLACK
	}
}

func transplant(tree *RBTree, u *RBNode, v *RBNode) {
	if u.parent == nil {
		tree.root = v
	} else if u == u.parent.left {
		u.parent.left = v
	} else {
		u.parent.right = v
	}

	if v != nil {
		v.parent = u.parent
	}
}

func DelRBNode(tree *RBTree, value string) {
	z := GetRBNode(tree, value)
	if z == nil {
		return
	}

	y := z
	var x *RBNode
	yOriginalColor := y.color

	if z.left == nil {
		x = z.right
		transplant(tree, z, x)
	} else if z.right == nil {
		x = z.left
		transplant(tree, z, x)
	} else {
		y = treeMinimum(z.right)
		yOriginalColor = y.color
		x = y.right

		if y.parent != z {
			transplant(tree, y, x)
			y.right = z.right
			if y.right != nil {
				y.right.parent = y
			}
		} else {
			if x != nil {
				x.parent = y
			}
		}

		transplant(tree, z, y)
		y.left = z.left
		if y.left != nil {
			y.left.parent = y
		}
		y.color = z.color
	}

	if yOriginalColor == BLACK {
		if x != nil {
			deleteFix(tree, x)
		}
	}
}

// ==================== МНОЖЕСТВО НА ОСНОВЕ КРАСНО-ЧЕРНОГО ДЕРЕВА ====================

type StringSet struct {
	tree *RBTree
}

func NewStringSet() *StringSet {
	return &StringSet{tree: NewRBTree()}
}

func NewStringSetFromSlice(elements []string) *StringSet {
	set := NewStringSet()
	for _, elem := range elements {
		set.Add(elem)
	}
	return set
}

func (s *StringSet) Add(value string) {
	AddRBNode(s.tree, value)
}

func (s *StringSet) Remove(value string) {
	DelRBNode(s.tree, value)
}

func (s *StringSet) Contains(value string) bool {
	return GetRBNode(s.tree, value) != nil
}

func (s *StringSet) Size() int {
	count := 0
	s.ForEach(func(value string) {
		count++
	})
	return count
}

func (s *StringSet) Empty() bool {
	return s.tree.root == nil
}

func (s *StringSet) Clear() {
	s.tree.root = nil
}

func (s *StringSet) UnionWith(other *StringSet) *StringSet {
	result := NewStringSet()
	s.ForEach(func(value string) {
		result.Add(value)
	})
	other.ForEach(func(value string) {
		result.Add(value)
	})
	return result
}

func (s *StringSet) IntersectionWith(other *StringSet) *StringSet {
	result := NewStringSet()
	s.ForEach(func(value string) {
		if other.Contains(value) {
			result.Add(value)
		}
	})
	return result
}

func (s *StringSet) DifferenceWith(other *StringSet) *StringSet {
	result := NewStringSet()
	s.ForEach(func(value string) {
		if !other.Contains(value) {
			result.Add(value)
		}
	})
	return result
}

func (s *StringSet) Equals(other *StringSet) bool {
	if s.Size() != other.Size() {
		return false
	}
	equal := true
	s.ForEach(func(value string) {
		if !other.Contains(value) {
			equal = false
		}
	})
	return equal
}

func (s *StringSet) ForEach(f func(string)) {
	inOrderTraversal(s.tree.root, f)
}

func inOrderTraversal(node *RBNode, f func(string)) {
	if node == nil {
		return
	}
	inOrderTraversal(node.left, f)
	f(node.data)
	inOrderTraversal(node.right, f)
}

func (s *StringSet) Print() {
	s.ForEach(func(value string) {
		print(value + " ")
	})
	println()
}
