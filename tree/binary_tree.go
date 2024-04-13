package tree

import "fmt"

type TreeNode struct {
	Data  int
	Left  *TreeNode
	Right *TreeNode
}

func NewTreeNode(data int) *TreeNode {
	return &TreeNode{
		Data:  data,
		Left:  nil,
		Right: nil,
	}
}

func (t *TreeNode) Insert(data int) bool {
	if t.Data == data {
		return false
	} else if t.Data > data {
		if t.Left == nil {
			t.Left = NewTreeNode(data)
			return true
		} else {
			return t.Left.Insert(data)
		}
	} else {
		if t.Right == nil {
			t.Right = NewTreeNode(data)
			return true
		} else {
			return t.Right.Insert(data)
		}
	}
}

func (t *TreeNode) Min() *TreeNode {
	current := t
	for current.Left != nil {
		current = current.Left
	}
	return current
}

func (t *TreeNode) Delete(data int) *TreeNode {
	if t == nil {
		return nil
	}
	if t.Data > data {
		t.Left = t.Left.Delete(data)
	} else if t.Data < data {
		t.Right = t.Right.Delete(data)
	} else {
		if t.Left == nil {
			temp := t.Right
			t = nil 
			return temp
		} else if t.Right == nil {
			temp := t.Left
			t = nil
			return temp
		}

		temp := t.Right.Min()
		t.Data = temp.Data
		t.Right = t.Right.Delete(temp.Data)
	}
	
	return t
} 

func (t *TreeNode) Find(data int) bool {
	if t == nil {
		return false
	}

	if t.Data == data {
		return true
	} else if t.Data > data {
		return t.Left.Find(data)
	} else {
		return t.Right.Find(data)
	}
}

func (t *TreeNode) PreOrder() {
	if t != nil {
		fmt.Printf("%d ", t.Data)
		t.Left.PreOrder()
		t.Right.PreOrder()
	}
}

func (t *TreeNode) InOrder() {
	if t != nil {
		t.Left.InOrder()
		fmt.Printf("%d ", t.Data)
		t.Right.InOrder()
	}
}

func (t *TreeNode) PostOrder() {
	if t != nil {
		t.Left.PostOrder()
		t.Right.PostOrder()
		fmt.Printf("%d ", t.Data)
	}
}