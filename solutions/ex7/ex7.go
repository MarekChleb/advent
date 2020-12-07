package ex7

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

const Path = "./solutions/ex7"

type Edge struct {
	to    string
	count int
}

type Node struct {
	edges []Edge
}

func A(lines []string) {
	graph := map[string]Node{}
	count := 0
	hasShinyBag := map[string]bool{}

	for _, l := range lines {
		node, name := CreateNodeFromSentence(l)
		graph[name] = node
	}

	for name, _ := range graph {
		has, ok := hasShinyBag[name]
		if !ok {
			traverseGraph(name, graph, &hasShinyBag)
			has = hasShinyBag[name]
		}

		if has {
			count++
		}
	}

	fmt.Println(count)
}

func B(lines []string) {
	count := 0
	graph := map[string]Node{}
	counts := map[string]int{}

	for _, l := range lines {
		node, name := CreateNodeFromSentence(l)
		graph[name] = node
	}

	count = howManyBags("shiny gold", &counts, graph) - 1
	fmt.Println(count)
}

func CreateNodeFromSentence(line string) (Node, string) {
	re := regexp.MustCompile(`^(?P<whichBag>.*) bags contain (?P<whatsInside>.*)[.]$`)
	re2 := regexp.MustCompile(`^(?P<count>\d+) (?P<color>.*) bags?$`)

	info := re.FindStringSubmatch(line)

	whichBag := info[1]
	whatsInside := info[2]

	if whatsInside == "no other bags" {
		return Node{}, whichBag
	}

	containsSplitted := strings.Split(whatsInside, ", ")

	edges := []Edge{}

	for _, contain := range containsSplitted {
		info = re2.FindStringSubmatch(contain)

		count := info[1]
		color := info[2]

		countInt64, _ := strconv.ParseInt(count, 10, 64)

		edges = append(edges, Edge{to: color, count: int(countInt64)})
	}

	node := Node{
		edges: edges,
	}

	return node, whichBag
}

func traverseGraph(name string, g map[string]Node, h *map[string]bool) {
	node := g[name]
	for _, e := range node.edges {
		to := e.to
		if to == "shiny gold" {
			(*h)[name] = true
			return
		}
		has, ok := (*h)[to]
		if !ok {
			traverseGraph(to, g, h)
			has = (*h)[to]
		}
		if has {
			(*h)[name] = true
			return
		}
	}
	(*h)[name] = false
}

func howManyBags(name string, counts *map[string]int, g map[string]Node) int {
	if v, ok := (*counts)[name]; ok {
		return v
	}

	node := g[name]
	count := 1
	for _, e := range node.edges {
		count += howManyBags(e.to, counts, g) * e.count
	}

	(*counts)[name] = count
	return count
}
