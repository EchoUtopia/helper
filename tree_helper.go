package leetcode

import (
	"encoding/json"
	"fmt"
	"github.com/goccy/go-graphviz/cgraph"
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func ParseTreeFromInput(inputStr string) *TreeNode {
	var input []*int
	if err := json.Unmarshal([]byte(inputStr), &input); err != nil {
		panic(err)
	}
	if len(input) == 0 {
		return nil
	} else if len(input) == 1 {
		return &TreeNode{*input[0], nil, nil}
	}
	r := &TreeNode{*input[0], nil, nil}
	input = input[1:]
	stack := []*TreeNode{r}
	stack1 := []*TreeNode{}
	for len(stack) != 0 {
		p := stack[0]
		stack = stack[1:]
		if len(input) != 0 {
			i := input[0]
			input = input[1:]
			if i != nil {
				p.Left = &TreeNode{*i, nil, nil}
				stack1 = append(stack1, p.Left)
			}
		}
		if len(input) != 0 {
			i := input[0]
			input = input[1:]
			if i != nil {
				p.Right = &TreeNode{*i, nil, nil}
				stack1 = append(stack1, p.Right)
			}
		}
		if len(stack) == 0 {
			stack = stack1
			stack1 = []*TreeNode{}
		}
	}
	return r
}

func PrintTree(prefix string, n *TreeNode) {
	if n != nil {
		PrintTree(prefix+"    ", n.Right)
		fmt.Printf("%s |-- %d\n", prefix, n.Val)
		PrintTree(prefix+"    ", n.Left)
	}
}

func balanceWhiteSpace(val string, cnt int) string {
	prefixCnt := cnt / 2
	SuffixCnt := cnt - prefixCnt
	b := make([]byte, 0, len(val)+cnt)
	for i := 0; i < prefixCnt; i++ {
		b = append(b, ' ')
	}
	b = append(b, val...)
	for i := 0; i < SuffixCnt; i++ {
		b = append(b, ' ')
	}
	return string(b)
}

func RenderTree(n *TreeNode) {
	if n == nil {
		return
	}

	stack := []*TreeNode{n}
	nextStack := []*TreeNode{}

	render(func(graph *cgraph.Graph) {
		// cgraph will treat nodes with the same name as the same node
		seq := map[int]int{}
		nodeMap := map[*TreeNode]*cgraph.Node{}
		createTreeNode := func(node *TreeNode) (*cgraph.Node, error) {
			var err error
			gNode, ok := nodeMap[node]
			if !ok {
				gNode, err = graph.CreateNode(balanceWhiteSpace(strconv.Itoa(node.Val), seq[node.Val]))
				seq[node.Val]++
				if err != nil {
					return nil, err
				}
				nodeMap[node] = gNode
			}
			return gNode, nil
		}

		for len(stack) != 0 {
			for _, v := range stack {
				s, err := createTreeNode(v)
				if err != nil {
					panic(err)
				}
				if v.Left != nil {
					d1, err := createTreeNode(v.Left)
					if err != nil {
						panic(err)
					}
					edge, err := graph.CreateEdge(``, s, d1)
					if err != nil {
						panic(err)
					}
					if v.Right == nil {
						edge.SetLabel(`left`)
					}
					nextStack = append(nextStack, v.Left)
				}
				if v.Right != nil {
					d2, err := createTreeNode(v.Right)
					if err != nil {
						panic(err)
					}
					nextStack = append(nextStack, v.Right)
					edge, err := graph.CreateEdge(``, s, d2)
					if err != nil {
						panic(err)
					}
					if v.Left == nil {
						edge.SetLabel(`right`)
					}
				}
			}
			stack, nextStack = nextStack, stack
			nextStack = nextStack[:0]
		}
	})
}
