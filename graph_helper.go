package leetcode

import (
	"encoding/json"
	"github.com/goccy/go-graphviz/cgraph"
	"strconv"
)

func ParseGraphInput(in string) [][]int {
	graph := [][]int{}
	if err := json.Unmarshal([]byte(in), &graph); err != nil {
		panic(err)
	}
	return graph
}

func ParseEdgesInput(in string) [][]int {

	edges := [][]int{}
	if err := json.Unmarshal([]byte(in), &edges); err != nil {
		panic(err)
	}
	return edges
}

func RenderGraph(graph [][]int) {
	edges := [][]int{}
	for s, ds := range graph {
		for _, d := range ds {
			edges = append(edges, []int{s, d})
		}
	}
	RenderGraphByEdges(edges)
}

func RenderGraphByEdges(edges [][]int) {
	render(func(graph *cgraph.Graph) {
		var err error
		nodes := map[int]*cgraph.Node{}
		for _, l := range edges {
			s, d := l[0], l[1]
			sn := nodes[s]
			if sn == nil {
				sn, err = graph.CreateNode(strconv.Itoa(s))
				if err != nil {
					panic(err)
				}
				nodes[s] = sn
			}
			dn := nodes[d]
			if dn == nil {
				dn, err = graph.CreateNode(strconv.Itoa(d))
				if err != nil {
					panic(err)
				}
				nodes[d] = dn
			}
			if _, err := graph.CreateEdge(``, sn, dn); err != nil {
				panic(err)
			}
		}
	})
}

func RenderGraphByEdgesMap(edges map[int][]int) {

	render(func(graph *cgraph.Graph) {
		var err error
		nodes := map[int]*cgraph.Node{}
		for s, l := range edges {
			for _, d := range l {
				sn := nodes[s]
				if sn == nil {
					sn, err = graph.CreateNode(strconv.Itoa(s))
					if err != nil {
						panic(err)
					}
					nodes[s] = sn
				}
				dn := nodes[d]
				if dn == nil {
					dn, err = graph.CreateNode(strconv.Itoa(d))
					if err != nil {
						panic(err)
					}
					nodes[d] = dn
				}
				if _, err := graph.CreateEdge(``, sn, dn); err != nil {
					panic(err)
				}
			}

		}
	})
}
