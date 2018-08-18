package bintree

import (
	"fmt"
)

type gtCountNode struct {
	value   int
	eqCount uint64
	gtCount uint64
	left    *gtCountNode
	right   *gtCountNode
	balance int
}

func insertBalanced(rootPtrPtr **gtCountNode, value int) {
	path := []*gtCountNode{nil}
	nodePtrPtr := rootPtrPtr
	for {
		node := *nodePtrPtr
		if node == nil {
			node = &gtCountNode{value, 1, 0, nil, nil, 0}
			*nodePtrPtr = node
			path = append(path, node)
			break
		}
		path = append(path, node)
		if value < node.value {
			nodePtrPtr = &node.left
		} else if value > node.value {
			node.gtCount++
			nodePtrPtr = &node.right
		} else {
			node.eqCount++
			return
		}
	}

	// return

	for i := len(path) - 1; i >= 2; i-- {
		z := path[i]
		x := path[i-1]
		g := path[i-2]
		var n *gtCountNode

		if z == x.right {
			// right subtree increases
			if x.balance > 0 {
				// x is already right heavy, rebalancing required
				if z.balance < 0 {
					// right-left case
					n = rotateRL(x, z)
				} else {
					n = rotateL(x, z)
				}
			} else {
				x.balance++
				if x.balance == 0 {
					// z height increase absorbed
					break
				} else {
					// z height increases
					continue
				}
			}
		} else {
			// left subtree increases
			if x.balance < 0 {
				// x is already left heavy, rebalancing required
				if z.balance > 0 {
					// left-right case
					n = rotateLR(x, z)
				} else {
					n = rotateR(x, z)
				}
			} else {
				x.balance--
				if x.balance == 0 {
					// z height increase absorbed
					break
				} else {
					// z height increases
					continue
				}
			}
		}
		if g == nil {
			// n is the new tree root
			*rootPtrPtr = n
			break
		} else {
			if x == g.left {
				// x is a left child
				g.left = n
			} else if x == g.right {
				g.right = n
			} else {
				panic(fmt.Sprintf("x = %v; g.left = %v; g.right = %v", x, g.left, g.right))
			}
			break
		}
	}
}

func rotateRL(x, z *gtCountNode) *gtCountNode {
	// z is by 2 higher than it's sibling
	y := z.left // inner child of z
	t3 := y.right
	z.left = t3
	y.right = z
	t2 := y.left
	x.right = t2
	y.left = x
	if y.balance > 0 { // t3 was higher
		x.balance = -1 // t1 now higher
		z.balance = 0
	} else {
		if y.balance == 0 {
			x.balance = 0
			z.balance = 0
		} else {
			// t2 was higher
			x.balance = 0
			z.balance = 1
		}
	}
	y.balance = 0

	x.gtCount = x.gtCount - y.gtCount - z.gtCount - y.eqCount - z.eqCount
	y.gtCount = y.gtCount + z.gtCount + z.eqCount

	return y
}

func rotateL(x, z *gtCountNode) *gtCountNode {
	// z is by 2 higher than its sibling
	t23 := z.left // inner child of z
	x.right = t23
	z.left = x
	if z.balance == 0 {
		x.balance = 1  // t23 is now higher
		z.balance = -1 // t3 is now lother than x
	} else {
		x.balance = 0
		z.balance = 0
	}

	x.gtCount = x.gtCount - z.gtCount - z.eqCount

	return z // new root of rotated value
}

func rotateLR(x, z *gtCountNode) *gtCountNode {
	// z is by 2 higher than it's sibling
	y := z.right // inner child of z
	t3 := y.left
	z.right = t3
	y.left = z
	t2 := y.right
	x.left = t2
	y.right = x
	if y.balance < 0 { // t3 was higher
		x.balance = 1 // t1 now higher
		z.balance = 0
	} else {
		if y.balance == 0 {
			x.balance = 0
			z.balance = 0
		} else {
			// t2 was higher
			x.balance = 0
			z.balance = -1
		}
	}
	y.balance = 0

	z.gtCount = z.gtCount - y.gtCount - y.eqCount
	y.gtCount = y.gtCount + x.gtCount + x.eqCount

	return y
}

func rotateR(x, z *gtCountNode) *gtCountNode {
	// z is by 2 higher than its sibling
	t23 := z.right // inner child of z
	x.left = t23
	z.right = x
	if z.balance == 0 {
		x.balance = -1 // t23 is now higher
		z.balance = 1  // t3 is now lother than x
	} else {
		x.balance = 0
		z.balance = 0
	}

	z.gtCount = z.gtCount + x.gtCount + x.eqCount

	return z // new root of rotated value
}

func CountGT(node *gtCountNode, value int) uint64 {
	if node == nil {
		return 0
	} else if value < node.value {
		return CountGT(node.left, value) + node.gtCount + node.eqCount
	} else if value > node.value {
		return CountGT(node.right, value)
	} else {
		return node.gtCount
	}
}
