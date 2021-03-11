package leetcode

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestPrintTree(t *testing.T) {
	input := `[9,6,-3,null,null,-6,2,null,null,2,null,-6,-6,-6]`
	tree := ParseTreeFromInput(input)
	PrintTree(``, tree)
}

func TestRenderTree(t *testing.T) {
	input := `[9,6,-3,null,null,-6,2,null,null,2,null,-6,-6,-6]`
	tree := ParseTreeFromInput(input)
	RenderTree(tree)
}

func TestGraph(t *testing.T) {
	input := `[[4,3,1],[3,2,4],[3],[4],[]]`
	graph := ParseGraphInput(input)
	RenderGraph(graph)
}

func TestGraphEdges(t *testing.T) {
	input := `[[0,1],[0,2],[2,5],[3,4],[4,2]]`
	edges := ParseEdgesInput(input)
	RenderGraphByEdges(edges)
}

func TestLinkedList(t *testing.T) {
	input := `1,2,3,4,5,6`
	head := ParseLinkedListFromStr(input)
	// get int linked list instead of string
	//head := GetIntLinkedListFromStr(input)
	PrintLinkedListNode(head)
}

func TestIntMatrix(t *testing.T) {
	input := `[[0,0,1,1],[1,0,1,0],[1,1,0,0]]`
	matrix := ParseIntMatrix(input)
	PrintIntMatrix(matrix)
}

func TestIntArray(t *testing.T) {
	input := `[0,0,1,1]`
	ParseIntSlice(input)
}

func goodNodes(root *TreeNode) int {
	return 1 + goodNodesHelper(root, root.Val)
}

func goodNodesHelper(p *TreeNode, max int) int {
	if p == nil {
		return 0
	}
	if p.Left != nil {
		lm := max
		cnt := 0
		if p.Left.Val >= lm {
			lm = p.Left.Val
			cnt = 1
		}
		return cnt + goodNodesHelper(p.Left, lm)
	}
	if p.Right != nil {
		rm := max
		cnt := 0
		if p.Right.Val >= rm {
			rm = p.Right.Val
			cnt = 1
		}
		return cnt + goodNodesHelper(p.Right, rm)
	}
	return 0
}

func TestTmp(t *testing.T) {
	inputs := []string{
		`[3,1,4,3,null,1,5]`,
		`[3,3,null,4,2]`,
		`[1]`,
	}
	res := []int{4, 3, 1}
	for k, v := range inputs {
		in := ParseTreeFromInput(v)
		require.Equal(t, res[k], goodNodes(in))
	}
}

func TestTmp2(t *testing.T) {

}
