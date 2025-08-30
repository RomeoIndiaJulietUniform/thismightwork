package index

import (
	"container/heap"
	"math"
	"math/rand"
	"sort"
)

type Node struct {
	Vector    []float64
	Level     int
	Neighbors [][]*Node
}

type NodeDist struct {
	Node *Node
	Dist float64
}

type PriorityQueue []*NodeDist

func (pq PriorityQueue) Len() int            { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool  { return pq[i].Dist < pq[j].Dist }
func (pq PriorityQueue) Swap(i, j int)       { pq[i], pq[j] = pq[j], pq[i] }
func (pq *PriorityQueue) Push(x interface{}) { *pq = append(*pq, x.(*NodeDist)) }
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[:n-1]
	return item
}

type HNSW struct {
	M              int
	EF             int
	EFConstruction int
	Layers         [][]*Node
	Nodes          []*Node
}

func NewHNSW(M, EF, EFConstruction int) *HNSW {
	return &HNSW{
		M:              M,
		EF:             EF,
		EFConstruction: EFConstruction,
		Layers:         make([][]*Node, 0),
		Nodes:          make([]*Node, 0),
	}
}

func distance(a, b []float64) float64 {
	sum := 0.0
	for i := range a {
		diff := a[i] - b[i]
		sum += diff * diff
	}
	return math.Sqrt(sum)
}

func randomLevel(maxLevel int) int {
	level := 0
	p := 0.5
	for rand.Float64() < p && level < maxLevel {
		level++
	}
	return level
}

func (h *HNSW) AddNode(vector []float64) {
	level := randomLevel(5)
	node := &Node{
		Vector:    vector,
		Level:     level,
		Neighbors: make([][]*Node, level+1),
	}
	h.Nodes = append(h.Nodes, node)

	for len(h.Layers) <= level {
		h.Layers = append(h.Layers, make([]*Node, 0))
	}
	h.Layers[level] = append(h.Layers[level], node)

	if len(h.Nodes) == 1 {
		return
	}

	entry := h.Nodes[0]
	for l := level; l >= 0; l-- {
		entry = h.searchLayer(vector, entry, l, h.EFConstruction)
		h.connectNeighbors(node, entry, l)
	}
}

func (h *HNSW) searchLayer(query []float64, entry *Node, level, ef int) *Node {
	visited := make(map[*Node]bool)
	pq := &PriorityQueue{}
	heap.Init(pq)
	bestDist := distance(query, entry.Vector)
	best := entry
	heap.Push(pq, &NodeDist{Node: entry, Dist: bestDist})
	visited[entry] = true

	for pq.Len() > 0 {
		nd := heap.Pop(pq).(*NodeDist)
		curr := nd.Node
		if nd.Dist > bestDist {
			break
		}

		for _, neighbor := range curr.Neighbors[level] {
			if !visited[neighbor] {
				visited[neighbor] = true
				d := distance(query, neighbor.Vector)
				if d < bestDist {
					bestDist = d
					best = neighbor
				}
				heap.Push(pq, &NodeDist{Node: neighbor, Dist: d})
				if pq.Len() > ef {
					heap.Pop(pq)
				}
			}
		}
	}
	return best
}

func (h *HNSW) connectNeighbors(node, entry *Node, level int) {
	candidates := append([]*Node{}, entry.Neighbors[level]...)
	candidates = append(candidates, entry)

	sort.Slice(candidates, func(i, j int) bool {
		return distance(candidates[i].Vector, node.Vector) < distance(candidates[j].Vector, node.Vector)
	})

	limit := min(h.M, len(candidates))
	node.Neighbors[level] = append([]*Node{}, candidates[:limit]...)

	for _, neighbor := range node.Neighbors[level] {
		neighbor.Neighbors[level] = append(neighbor.Neighbors[level], node)
		if len(neighbor.Neighbors[level]) > h.M {
			neighbor.Neighbors[level] = neighbor.Neighbors[level][:h.M]
		}
	}
}

func (h *HNSW) SearchKNN(query []float64, k int) [][]float64 {
	if len(h.Nodes) == 0 {
		return nil
	}

	entry := h.Nodes[0]
	for l := len(h.Layers) - 1; l >= 0; l-- {
		entry = h.searchLayer(query, entry, l, h.EF)
	}

	pq := &PriorityQueue{}
	heap.Init(pq)
	heap.Push(pq, &NodeDist{Node: entry, Dist: distance(query, entry.Vector)})

	results := make([][]float64, 0, k)
	for pq.Len() > 0 && len(results) < k {
		nd := heap.Pop(pq).(*NodeDist)
		results = append(results, nd.Node.Vector)
	}
	return results
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
