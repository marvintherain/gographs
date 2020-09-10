package main

import "fmt"

const maxv = 1000

type edgenode struct {
	y      int
	weight int
	next   *edgenode
}

type graph struct {
	edges     [maxv + 1]*edgenode
	degree    [maxv + 1]int
	nvertices int
	nedges    int
	directed  bool
}

func initializegraph(directed bool) *graph {
	g := new(graph)

	g.nvertices = 0
	g.nedges = 0
	g.directed = directed

	for i := 1; i <= maxv; i++ {
		g.degree[i] = 0
		g.edges[i] = nil
	}

	return g
}

func readgraph(g *graph, directed bool) {
	var m int
	var x, y int

	fmt.Scanf("%d %d", &g.nvertices, &m)

	for i := 1; i <= m; i++ {
		fmt.Scanf("%d %d", &x, &y)
		insertedge(g, x, y, directed)

	}

}

func insertedge(g *graph, x int, y int, directed bool) {
	var p *edgenode = new(edgenode)

	p.weight = 0
	p.y = y
	p.next = g.edges[x]

	g.edges[x] = p

	g.degree[x]++

	if directed == false {
		insertedge(g, y, x, true)
	} else {
		g.nedges++
	}
}

func printgraph(g *graph) {
	var p *edgenode = new(edgenode)

	for i := 1; i <= g.nvertices; i++ {
		fmt.Printf("%d: ", i)
		p = g.edges[i]
		for p != nil {
			fmt.Printf(" %d", p.y)
			p = p.next
		}
		fmt.Printf("\n")
	}
}

func main() {
	mygraph := initializegraph(true)

	readgraph(mygraph, true)

	printgraph(mygraph)

}
