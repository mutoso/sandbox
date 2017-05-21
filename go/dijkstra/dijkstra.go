// © 2017 Alastair Feille
// License: CC0

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/emirpasic/gods/trees/binaryheap"
	"github.com/emirpasic/gods/utils"
)

type Vertex struct {
	name string
	dist float64
	pred *Vertex
}

type Edge struct {
	dest   *Vertex
	weight float64
}

type Graph struct {
	vertexes map[string]*Vertex
	edges    map[string][]Edge
}

func main() {
	if len(os.Args) < 3 {
		log.Fatal("Please provide both source and destination airport codes.")
	}
	src := os.Args[1]
	dest := os.Args[2]
	path, totalCost := runDijkstra(src, dest)
	fmt.Println(path, totalCost)
}

func runDijkstra(src, dest string) (path []string, totalCost float64) {
	graph := buildGraph("airports.txt")
	graph.dijkstra(src)
	return graph.getResult(dest)
}

func buildGraph(filename string) Graph {

	// read adjacency list file
	lines, err := readlines(filename)
	if err != nil {
		log.Fatal(err)
	}

	g := Graph{
		vertexes: make(map[string]*Vertex),
		edges:    make(map[string][]Edge),
	}

	// populate vertex list
	for _, line := range lines {
		splitLine := strings.Fields(line)
		initialVertex := splitLine[0]
		g.vertexes[initialVertex] = &Vertex{name: initialVertex}
	}

	// populate edge list
	for _, line := range lines {
		splitLine := strings.Fields(line)
		initialVertex := splitLine[0]
		connections := splitLine[1:]

		// for each entry in the adjacency list
		for i := 0; i < len(connections); i += 2 {
			dest := connections[i]
			weight, _ := strconv.ParseFloat(connections[i+1], 64)
			g.edges[initialVertex] = append(g.edges[initialVertex], Edge{dest: g.vertexes[dest], weight: weight})
		}
	}

	return g
}

func (g *Graph) getResult(dest string) (path []string, totalCost float64) {
	for cur := g.vertexes[dest]; cur != nil; cur = cur.pred {
		path = append(path, cur.name)
	}
	a := path
	// reverse the slice
	for i := len(a)/2 - 1; i >= 0; i-- {
		opp := len(a) - 1 - i
		a[i], a[opp] = a[opp], a[i]
	}
	return a, g.vertexes[dest].dist
}

func readlines(filename string) ([]string, error) {
	inputbytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return []string{}, err
	}
	inputtext := string(inputbytes)
	lines := strings.Split(inputtext, "\n")

	// return array without empty last line
	return lines[:len(lines)-1], nil
}

func (g *Graph) dijkstra(src string) {
	vertexComparator := func(a, b interface{}) int {
		v1 := a.(*Vertex)
		v2 := b.(*Vertex)

		// min-heap based on the vertex's distance
		return utils.Float64Comparator(v1.dist, v2.dist)
	}
	heap := binaryheap.NewWith(vertexComparator)

	// initialize each vertex to inf dist and no predecessor
	for _, v := range g.vertexes {
		// v.dist = ∞
		v.dist = math.MaxFloat64
		v.pred = nil
	}

	// set the starting vertex's distance to zero
	s := g.vertexes[src]
	s.dist = 0

	// push all of them onto the heap
	for _, v := range g.vertexes {
		heap.Push(v)
	}

	for !heap.Empty() {
		o, _ := heap.Pop()
		u := o.(*Vertex)
		// for each vertex v adjacent to u
		for _, edge := range g.edges[u.name] {
			v := edge.dest
			g.relax(u, v)
		}
	}
}

func (g *Graph) relax(u *Vertex, v *Vertex) {
	if u.dist+g.w(u, v) < v.dist {
		v.dist = u.dist + g.w(u, v)
		v.pred = u
	}
}

func (g *Graph) w(u *Vertex, v *Vertex) float64 {
	for _, connection := range g.edges[u.name] {
		if connection.dest.name == v.name {
			return connection.weight
		}
	}
	return 0
}
