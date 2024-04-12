package tree

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
	
} 
