package algs

import (
	"cmp"
	"fmt"
)

type BSTNode[K cmp.Ordered, V any] struct {
	key         K
	val         V
	n           int
	left, right *BSTNode[K, V]
}

type BST[K cmp.Ordered, V any] struct {
	root *BSTNode[K, V]
}

func NewBST[K cmp.Ordered, V any]() *BST[K, V] {
	return &BST[K, V]{}
}

func (bst *BST[K, V]) Put(key K, val V) {
	bst.root = bst.put(bst.root, key, val)
}

func (bst *BST[K, V]) put(x *BSTNode[K, V], key K, val V) *BSTNode[K, V] {
	if x == nil {
		return &BSTNode[K, V]{key: key, val: val, n: 1}
	}

	c := cmp.Compare(key, x.key)
	if c < 0 {
		x.left = bst.put(x.left, key, val)
	} else if c > 0 {
		x.right = bst.put(x.right, key, val)
	} else {
		x.val = val
	}
	x.n = 1 + bst.size(x.left) + bst.size(x.right)
	return x
}

func (bst *BST[K, V]) Get(key K) (V, bool) {
	return bst.get(bst.root, key)
}

func (bst *BST[K, V]) get(x *BSTNode[K, V], key K) (V, bool) {
	if x == nil {
		var zero V
		return zero, false
	}

	c := cmp.Compare(key, x.key)
	if c < 0 {
		return bst.get(x.left, key)
	} else if c > 0 {
		return bst.get(x.right, key)
	}
	return x.val, true
}

func (bst *BST[K, V]) Contains(key K) bool {
	_, ok := bst.Get(key)
	return ok
}

func (bst *BST[K, V]) Delete(key K) {
	bst.root = bst.delete(bst.root, key)
}

func (bst *BST[K, V]) delete(x *BSTNode[K, V], key K) *BSTNode[K, V] {
	if x == nil {
		return nil
	}

	c := cmp.Compare(key, x.key)
	if c < 0 {
		x.left = bst.delete(x.left, key)
	} else if c > 0 {
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

func (bst *BST[K, V]) DeleteMin() {
	bst.root = bst.deleteMin(bst.root)
}

func (bst *BST[K, V]) deleteMin(x *BSTNode[K, V]) *BSTNode[K, V] {
	if x.left == nil {
		return x.right
	}
	x.left = bst.deleteMin(x.left)
	x.n = bst.size(x.left) + bst.size(x.right) + 1
	return x
}

func (bst *BST[K, V]) DeleteMax() {
	bst.root = bst.deleteMax(bst.root)
}

func (bst *BST[K, V]) deleteMax(x *BSTNode[K, V]) *BSTNode[K, V] {
	if x.right == nil {
		return x.left
	}
	x.right = bst.deleteMax(x.right)
	x.n = bst.size(x.left) + bst.size(x.right) + 1
	return x
}

func (bst *BST[K, V]) Min() K {
	return bst.min(bst.root).key
}

func (bst *BST[K, V]) min(x *BSTNode[K, V]) *BSTNode[K, V] {
	if x.left == nil {
		return x
	}
	return bst.min(x.left)
}

func (bst *BST[K, V]) Max() K {
	return bst.max(bst.root).key
}

func (bst *BST[K, V]) max(x *BSTNode[K, V]) *BSTNode[K, V] {
	if x.right == nil {
		return x
	}
	return bst.max(x.right)
}

func (bst *BST[K, V]) Size() int {
	return bst.size(bst.root)
}

func (bst *BST[K, V]) size(x *BSTNode[K, V]) int {
	if x == nil {
		return 0
	}
	return x.n
}

func (bst *BST[K, V]) IsEmpty() bool {
	return bst.size(bst.root) == 0
}

func (bst *BST[K, V]) Inorder() {
	bst.inorder(bst.root)
}

func (bst *BST[K, V]) inorder(x *BSTNode[K, V]) {
	if x == nil {
		return
	}
	bst.inorder(x.left)
	fmt.Print(" ", x.key, " ")
	bst.inorder(x.right)
}

func (bst *BST[K, V]) IsBST() bool {
    return bst.isBST(bst.root, nil, nil)
}

func (bst *BST[K, V]) isBST(x *BSTNode[K, V], minKey *K, maxKey *K) bool {
	if x == nil {
		return true
	}
    if minKey != nil && cmp.Compare(x.key, *minKey) <= 0 {
		return false
	}
    if maxKey != nil && cmp.Compare(x.key, *maxKey) >= 0 {
		return false
	}
    return bst.isBST(x.left, minKey, &x.key) && bst.isBST(x.right, &x.key, maxKey)
}
