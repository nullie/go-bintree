package bintree

import (
	"math/rand"
	"sort"
	"testing"
)

func checkDepth(node *gtCountNode) (int, bool) {
	if node == nil {
		return 0, true
	}
	left, ok := checkDepth(node.left)
	if !ok {
		return 0, false
	}
	right, ok := checkDepth(node.right)
	if !ok {
		return 0, false
	}
	if left > right+1 || right > left+1 {
		return 0, false
	}
	if left > right {
		return left + 1, true
	} else {
		return right + 1, true
	}
}

func inOrder(node *gtCountNode) (res []int) {
	if node == nil {
		return
	}
	res = append(res, inOrder(node.left)...)
	res = append(res, node.value)
	res = append(res, inOrder(node.right)...)
	return
}

func TestAVL(t *testing.T) {
	testCases := []struct {
		name      string
		generator func() []int
	}{
		{"Ascending", func() []int {
			n := 1000
			res := make([]int, n)
			for i := 0; i < n; i++ {
				res[i] = int(i)
			}
			return res
		}},
		{"Descending", func() []int {
			n := 1000
			res := make([]int, n)
			for i := 0; i < n; i++ {
				res[i] = int(n - i)
			}
			return res
		}},
		{"Random", func() []int {
			n := 1000
			res := make([]int, n)
			for i := 0; i < n; i++ {
				res[i] = rand.Int()
			}
			return res
		}},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			a := tc.generator()
			var tree *gtCountNode
			for _, value := range a {
				insertBalanced(&tree, value)
			}
			if _, ok := checkDepth(tree); !ok {
				t.Error("Unbalanced tree")
			}
			sort.Ints(a)
			b := inOrder(tree)
			for i, x := range a {
				if b[i] != x {
					t.Error("In order not equals sorted")
				}
			}
		})
	}
}
