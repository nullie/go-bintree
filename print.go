package bintree

import (
	"fmt"
)

type TreeNode interface {
	Value() string
	Left() TreeNode
	Right() TreeNode
	Leaf() bool
}

type Node struct {
	value int
	left  *Node
	right *Node
}

func (n *Node) Value() string {
	return fmt.Sprint(n.value)
}

func (n *Node) Left() TreeNode {
	return n.left
}

func (n *Node) Right() TreeNode {
	return n.right
}

func (n *Node) Leaf() bool {
	return n == nil
}

func (n *Node) String() string {
	return fmt.Sprintf("{%s %d %s}", n.left, n.value, n.right)
}

type PathNode struct {
	left bool
	node *Node
}

func (pn PathNode) String() string {
	return fmt.Sprintf("%d %v", pn.node.value, pn.left)
}

func Path(tree *Node, value int) (path []PathNode) {
	node := tree
	for node != nil {
		pn := PathNode{false, node}
		if value <= node.value {
			pn.left = true
			node = node.left
		} else {
			node = node.right
		}
		path = append(path, pn)
	}
	return
}

func Split(path []PathNode) (left, right *Node) {
	for i := len(path) - 1; i >= 0; i-- {
		pn := path[i]
		p := pn.node
		if pn.left {
			p.left = right
			right = p
		} else {
			p.right = left
			left = p
		}
	}
	return
}

func BuildSorted(values []int) (res *Node) {
	if len(values) == 0 {
		return nil
	} else if len(values) == 1 {
		return &Node{values[0], nil, nil}
	}

	mid := len(values) / 2
	left := BuildSorted(values[:mid])
	right := BuildSorted(values[mid+1:])
	return &Node{values[mid], left, right}
}

func InOrder(node *Node) (res []int) {
	if node == nil {
		return
	}
	res = append(res, InOrder(node.left)...)
	res = append(res, node.value)
	res = append(res, InOrder(node.right)...)
	return
}

func Print(root TreeNode) {
	printBranch(root.Left(), nil, false)
	fmt.Print("── ")
	if root != nil {
		fmt.Println(root.Value())
	} else {
		fmt.Println(nil)
	}
	printBranch(root.Right(), nil, true)
}

func printBranch(node TreeNode, prefix []bool, below bool) {
	if node.Leaf() {
		return
	}
	pl := len(prefix)
	np := make([]bool, pl+1)
	copy(np, prefix)
	np[pl] = below
	printBranch(node.Left(), np, false)
	printValue(node, prefix, below)
	np[pl] = !below
	printBranch(node.Right(), np, true)
}

func printValue(node TreeNode, prefix []bool, below bool) {
	for _, b := range prefix {
		if b {
			fmt.Print("   │")
		} else {
			fmt.Print("    ")
		}
	}
	if below {
		fmt.Print("   ╰── ")
	} else {
		fmt.Print("   ╭── ")
	}
	fmt.Println(node.Value())
}

type widthNode struct {
	node  TreeNode
	width int
	left  *widthNode
	right *widthNode
}

func (n *widthNode) Value() string {
	return fmt.Sprint(n.node.Value(), n.width)
}

func (n *widthNode) Left() TreeNode {
	return n.left
}

func (n *widthNode) Right() TreeNode {
	return n.right
}

func (n *widthNode) Leaf() bool {
	return n == nil
}

func calcWidth(root *Node) *widthNode {
	_, node := widthHelper(root, 0, 0)
	return node
}

func widthHelper(node TreeNode, minLeft, minRight int) (int, *widthNode) {
	if node.Leaf() {
		return 0, nil
	}
	vl := len(node.Value())
	vw := vl + vl%2
	hw := vw / 2
	lw, ln := widthHelper(node.Left(), 0, hw+1)
	rw, rn := widthHelper(node.Right(), hw+1, 0)
	w := max(minLeft, lw) + 1 + max(minRight, rw)
	return w, &widthNode{node, w, ln, rn}
}

func max(a, b int) int {
	if a < b {
		return b
	} else {
		return a
	}
}

func PrintWide(root *widthNode) {
	lines := printWideHelper(root, 0, 0, nil)
	for _, l := range lines {
		fmt.Println(string(l))
	}
}

func printWideHelper(node *widthNode, leftMargin, level int, lines [][]byte) [][]byte {
	if node.Leaf() {
		return lines
	}
	var line []byte
	if level >= len(lines) {
		lines = append(lines, line)
	} else {
		line = lines[level]
	}

	for i := len(line); i < leftMargin; i++ {
		line = append(line, ' ')
	}

	// hw := node.w / 2

	line = append(line, fmt.Sprint(node.Value())...)
	lines[level] = line

	lines = printWideHelper(node.left, leftMargin, level+1, lines)

	var rightMargin int
	if node.left != nil {
		rightMargin = leftMargin + node.left.width + 1
	} else {
		rightMargin = leftMargin + node.width/2 + 1
	}

	lines = printWideHelper(node.right, rightMargin, level+1, lines)

	return lines
}

//func main() {
//	tree := BuildSorted([]int{-10, -3, -1, 0, 2, 4, 5, 6, 10, 20})
//	fmt.Println(tree)
//	fmt.Println(InOrder(tree))
//
//	Print(tree)
//
//	path := Path(tree, 6)
//	fmt.Println(path)
//
//	left, right := Split(path)
//	fmt.Println(left)
//	fmt.Println(right)
//	fmt.Println(InOrder(left), InOrder(right))
//
//	Print(left)
//	Print(right)
//
//	wt := calcWidth(left)
//	Print(wt)
//	PrintWide(wt)
//}
