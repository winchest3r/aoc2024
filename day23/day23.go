package day23

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"sort"
	"strings"
)

type Node struct {
	Name string
	Cons []*Node
}

type Network map[string]*Node

func ReadInput(fname string) Network {
	file, err := os.Open(fname)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	sc := bufio.NewScanner(file)
	net := make(Network)
	for sc.Scan() {
		nodes := strings.Split(sc.Text(), "-")
		a, b := nodes[0], nodes[1]
		_, ok := net[a]
		if !ok {
			net[a] = &Node{a, make([]*Node, 0)}
		}
		_, ok = net[b]
		if !ok {
			net[b] = &Node{b, make([]*Node, 0)}
		}
		net[a].Cons = append(net[a].Cons, net[b])
		net[b].Cons = append(net[b].Cons, net[a])
	}
	return net
}

func FindThreeCons(net Network) int {
	result := make(map[string]bool)
	for aName, aNode := range net {
		for _, bNode := range aNode.Cons {
			bName := bNode.Name
			for _, cNode := range bNode.Cons {
				cName := cNode.Name
				if slices.Contains(aNode.Cons, cNode) {
					names := []string{aName, bName, cName}
					sort.Strings(names)
					result[strings.Join(names, ",")] = true
				}
			}
		}
	}
	return len(result)
}

func IsFullyConnected(nodes []*Node) bool {
	for i := 0; i < len(nodes); i++ {
		for j := i + 1; j < len(nodes); j++ {
			if !isConnected(nodes[i], nodes[j]) {
				return false
			}
		}
	}
	return true
}

func isConnected(a, b *Node) bool {
	for _, neighbor := range a.Cons {
		if neighbor == b {
			return true
		}
	}
	return false
}

func FindLargestNetwork(net Network) []string {
	var nodes []*Node
	for _, node := range net {
		nodes = append(nodes, node)
	}

	n := len(nodes)
	var result []string

	for subset := 1; subset < (1 << n); subset++ {
		var subsetNodes []*Node
		for i := 0; i < n; i++ {
			if subset&(1<<i) != 0 {
				subsetNodes = append(subsetNodes, nodes[i])
			}
		}

		if IsFullyConnected(subsetNodes) && len(subsetNodes) > len(result) {
			result = nil
			for _, node := range subsetNodes {
				result = append(result, node.Name)
			}
		}
	}

	return result
}

func FindThreeConsStartsWithT(net Network) int {
	result := make(map[string]bool)
	for aName, aNode := range net {
		for _, bNode := range aNode.Cons {
			bName := bNode.Name
			for _, cNode := range bNode.Cons {
				cName := cNode.Name
				if slices.Contains(aNode.Cons, cNode) {
					names := []string{aName, bName, cName}
					for _, n := range names {
						if strings.HasPrefix(n, "t") {
							sort.Strings(names)
							result[strings.Join(names, ",")] = true
							break
						}
					}
				}
			}
		}
	}
	return len(result)
}

func SolvePartOne(fname string) {
	fmt.Println(FindThreeConsStartsWithT(ReadInput(fname)))
}

func SolvePartTwo(fname string) {
	names := FindLargestNetwork(ReadInput(fname))
	sort.Strings(names)
	fmt.Println(strings.Join(names, ","))
}
