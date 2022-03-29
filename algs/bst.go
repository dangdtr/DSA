package algs

import "fmt"

type BSTNode struct {
	key         Key
	val         interface{}
	n           int
	left, right *BSTNode
}

type BST struct {
	root *BSTNode
}

func NewBST() *BST {
	return &BST{}
}

func (bst *BST) Put(key Key, val interface{}) {
	// if val == nil {
	// 	bst.Delete(key)
	// 	return
	// }
	bst.root = bst.put(bst.root, key, val)
}

func (bst *BST) put(x *BSTNode, key Key, val interface{}) *BSTNode {
	if x == nil {
		return &BSTNode{key: key, val: val, n: 1}
	}

	cmp := key.compareTo(x.key)
	// fmt.Println("putting", key, "_")
	if cmp < 0 {
		x.left = bst.put(x.left, key, val)
	} else if cmp > 0 {
		x.right = bst.put(x.right, key, val)
	} else {
		x.val = val
	}
	x.n = 1 + bst.size(x.left) + bst.size(x.right)
	return x
}

func (bst *BST) Get(key Key) interface{} {
	return bst.get(bst.root, key)
}

func (bst *BST) get(x *BSTNode, key Key) interface{} {
	if x == nil {
		return nil
	}

	cmp := key.compareTo(x.key)

	if cmp < 0 {
		bst.get(x.left, key)
	} else if cmp > 0 {
		bst.get(x.right, key)
	} // else {
	return x.val
	// }

}

func (bst *BST) Contains(key Key) bool {
	return bst.Get(key) != nil
}

func (bst *BST) Delete(key Key) {
	bst.root = bst.delete(bst.root, key)
}

func (bst *BST) delete(x *BSTNode, key Key) *BSTNode {
	if x == nil {
		return nil
	}

	cmp := key.compareTo(x.key)

	if cmp < 0 {
		x.left = bst.delete(x.left, key)
	} else if cmp > 0 {
		x.right = bst.delete(x.right, key)
	} else {
		if x.right == nil {
			return x.left
		}
		if x.left == nil {
			return x.right
		}
		t := x
		x = bst.min(t.right)
		x.right = bst.deleteMin(t.right)
		x.left = t.left
	}
	x.n = bst.size(x.left) + bst.size(x.right) + 1
	return x
}

func (bst *BST) DeleteMin() {
	bst.root = bst.deleteMin(bst.root)
}
func (bst *BST) deleteMin(x *BSTNode) *BSTNode {
	if x.left == nil {
		return x.right
	}
	x.left = bst.deleteMin(x.left)
	x.n = bst.size(x.left) + bst.size(x.right) + 1
	return x
}

func (bst *BST) DeleteMax() {
	bst.root = bst.deleteMax(bst.root)
}

func (bst *BST) deleteMax(x *BSTNode) *BSTNode {
	if x.right == nil {
		return x.left
	}
	x.right = bst.deleteMax(x.right)
	x.n = bst.size(x.left) + bst.size(x.right) + 1
	return x
}

func (bst *BST) Min() Key {
	return bst.min(bst.root).key
}

func (bst *BST) min(x *BSTNode) *BSTNode {
	if x.left == nil {
		return x
	}
	return bst.min(x.left)
}

func (bst *BST) Max() Key {
	return bst.max(bst.root).key
}

func (bst *BST) max(x *BSTNode) *BSTNode {
	if x.right == nil {
		return x
	}
	return bst.max(x.right)
}

func (bst *BST) Size() int {
	return bst.size(bst.root)
}
func (bst *BST) size(x *BSTNode) int {
	if x == nil {
		return 0
	}
	return x.n
}

func (bst *BST) IsEmpty() bool {
	return bst.size(bst.root) == 0
}

func (bst *BST) Inorder() {
	bst.inorder(bst.root)
}

func (bst *BST) inorder(x *BSTNode) {
	if x == nil {
		return
	}
	bst.inorder(x.left)
	fmt.Print(" ", x.key, " ")
	bst.inorder(x.right)
}

func (bst *BST) IsBST() bool {
	return bst.isBST(bst.root, nil, nil)
}
func (bst *BST) isBST(x *BSTNode, min Key, max Key) bool {
	if x == nil {
		return true
	}
	if min != nil && x.key.compareTo(min) <= 0 {
		return false
	}
	if max != nil && x.key.compareTo(max) >= 0 {
		return false
	}
	return bst.isBST(x.left, min, x.key) && bst.isBST(x.right, x.key, max)
}
