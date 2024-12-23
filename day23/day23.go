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

func GetConnectedNodes(node *Node) map[string]bool {
	set := map[string]bool{node.Name: true}
	for _, n := range node.Cons {
		set[n.Name] = true
	}
	for _, n := range node.Cons {
		if !set[n.Name] {
			continue
		}
		newSet := map[string]bool{n.Name: true}
		for _, nei := range n.Cons {
			newSet[nei.Name] = true
		}
		for name := range set {
			set[name] = set[name] && newSet[name]
		}
	}
	return set
}

func FindLargestNetwork(net Network) []string {
	result := make([]string, 0)
	for _, node := range net {
		data := GetConnectedNodes(node)
		for key, val := range data {
			if !val {
				delete(data, key)
			}
		}
		if len(result) < len(data) {
			result = make([]string, 0)
			for key := range data {
				result = append(result, key)
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
