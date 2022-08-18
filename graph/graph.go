package graph

type Graph interface {
	AddVertex(v int)
	AddEdge(u, v int)
	Directed() bool
	Dfs(func(u, v int))
	InDegree(v int) int
	OutDegree(v int) int
	GetByInDegree(degree int) []int
	GetByOutDegree(degree int) []int
	Topological() []int
}

type graph struct {
	v        int
	e        int
	edges    map[int]map[int]int
	directed bool

	inDegrees  map[int]int
	outDegrees map[int]int
	Graph
}

func New() *graph {
	return &graph{
		edges:      make(map[int]map[int]int),
		directed:   false,
		inDegrees:  make(map[int]int),
		outDegrees: make(map[int]int),
	}
}

func (g *graph) AddEdge(u, v int) {
	g.AddVertex(u)
	g.AddVertex(v)
	g.edges[u][v] = 1
	if !g.Directed() {
		g.edges[v][u] = 1
	} else {
		g.inDegrees[v] += 1
		g.outDegrees[u] += 1
	}
}

func (g *graph) AddVertex(v int) {
	if _, exist := g.edges[v]; !exist {
		g.edges[v] = make(map[int]int)
		g.inDegrees[v] = 0
	}
}

func (g *graph) Directed() bool {
	return g.directed
}

func (g *graph) Dfs(func(u, v int)) {

}

func (g *graph) GetByInDegree(degree int) []int {
	vertices := make([]int, 0)
	for v, d := range g.inDegrees {
		if d == degree {
			vertices = append(vertices, v)
		}
	}
	return vertices
}

func (g *graph) GetByOutDegree(degree int) []int {
	vertices := make([]int, 0)
	for v, d := range g.outDegrees {
		if d == degree {
			vertices = append(vertices, v)
		}
	}
	return vertices
}

func (g *graph) Topological() []int {
	res := make([]int, 0)
	inDegrees := g.getInDegrees()

	vertices := g.GetByInDegree(0)

	for len(vertices) != 0 {
		u := vertices[0]
		res = append(res, u)
		vertices = vertices[1:]
		for v, _ := range g.edges[u] {
			inDegrees[v] -= 1
			if inDegrees[v] == 0 {
				vertices = append(vertices, v)
			}
		}
	}

	return res
}

func (g *graph) getInDegrees() map[int]int {
	newInDegree := make(map[int]int)
	for v, d := range g.inDegrees {
		newInDegree[v] = d
	}
	return newInDegree
}
