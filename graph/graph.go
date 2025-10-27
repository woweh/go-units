package graph

import (
	"fmt"
	"strings"
)

var (
	Debug = false // enable debug message printing
)

func dbg(format string, a ...interface{}) {
	if Debug {
		fmt.Printf(format, a...)
	}
}

// Edge represents a directed connection between two nodes.
type Edge interface {
	From() string
	To() string
}

// Graph is a simple graph structure for BFS pathfinding.
type Graph struct {
	edges []Edge
}

// New creates a new Graph with optional initial edges.
func New(edges ...Edge) *Graph {
	b := &Graph{}
	for _, e := range edges {
		if e != nil {
			b.AddEdge(e)
		}
	}
	return b
}

func (b *Graph) Len() int      { return len(b.edges) }
func (b *Graph) Edges() []Edge { return b.edges }

// AddEdge adds an edge to the tree, skipping nils.
func (b *Graph) AddEdge(e Edge) {
	if e == nil {
		return
	}
	b.edges = append(b.edges, e)
}

// HasEdge returns true if the tree contains the given edge.
func (b *Graph) HasEdge(e Edge) bool {
	for _, edge := range b.edges {
		if edge == e {
			return true
		}
	}
	return false
}

// HasNode returns true if the tree contains the given node name.
func (b *Graph) HasNode(name string) bool {
	for _, e := range b.edges {
		if e.From() == name || e.To() == name {
			return true
		}
	}
	return false
}

// Reset removes all edges from the tree.
func (b *Graph) Reset() {
	b.edges = nil
}

// Copy returns a deep copy of the tree.
func (b *Graph) Copy() *Graph {
	newEdges := make([]Edge, len(b.edges))
	copy(newEdges, b.edges)
	return &Graph{edges: newEdges}
}

// Nodes returns unique node names in the tree.
func (b *Graph) Nodes() []string {
	names := make(map[string]struct{})
	for _, e := range b.edges {
		if e == nil {
			continue
		}
		names[e.From()] = struct{}{}
		names[e.To()] = struct{}{}
	}
	result := make([]string, 0, len(names))
	for n := range names {
		result = append(result, n)
	}
	return result
}

// fromNode returns all edges starting from a given node.
func (b *Graph) fromNode(start string) []Edge {
	var res []Edge
	for _, e := range b.edges {
		if e != nil && e.From() == start {
			res = append(res, e)
		}
	}
	return res
}

// FindPath finds a path from start to end using BFS.
func (b *Graph) FindPath(start, end string) (*Path, error) {
	if start == end {
		return nil, fmt.Errorf("start and end nodes are the same: %q", start)
	}
	var iter int
	var paths []*Path

	for _, e := range b.fromNode(start) {
		if e == nil {
			continue
		}
		p := newPath(e)
		if e.To() == end {
			return p, nil
		}
		paths = append(paths, p)
	}

	for len(paths) > 0 {
		var newPaths []*Path
		dbg("iter %d\n", iter)
		for _, p := range paths {
			dbg("  %s\n", p)
			last, err := p.Last()
			if err != nil {
				continue
			}
			children := b.fromNode(last.To())
			if len(children) == 0 {
				dbg("    dropped path (no children): %s\n", p)
				continue
			}
			for _, e := range children {
				if e == nil || p.IsCircular(e) {
					dbg("    dropped circular or nil child: %v\n", e)
					continue
				}
				np := newPath(p.edges...)
				np.AddEdge(e)
				dbg("    new path branch: %s\n", np)
				if e.To() == end {
					return np, nil
				}
				newPaths = append(newPaths, np)
			}
		}
		iter++
		paths = newPaths
	}
	return nil, fmt.Errorf("no path found from %q to %q", start, end)
}

// Path represents a sequence of edges.
type Path struct {
	*Graph
}

func newPath(edges ...Edge) *Path {
	p := &Path{&Graph{}}
	for _, e := range edges {
		if e != nil {
			p.AddEdge(e)
		}
	}
	return p
}

// String returns the path as a string of node names separated by '->'.
func (p *Path) String() string {
	nodes := p.Nodes()
	if len(nodes) == 0 {
		return ""
	}
	var sb strings.Builder
	sb.WriteString(nodes[0])
	for _, n := range nodes[1:] {
		sb.WriteString("->")
		sb.WriteString(n)
	}
	return sb.String()
}

// Last returns the last edge in the path, or error if empty.
func (p *Path) Last() (Edge, error) {
	if len(p.edges) == 0 {
		return nil, fmt.Errorf("path is empty")
	}
	return p.edges[len(p.edges)-1], nil
}

// Nodes returns names for all path nodes in order.
func (p *Path) Nodes() []string {
	if len(p.edges) == 0 {
		return nil
	}
	names := make([]string, 0, len(p.edges)+1)
	names = append(names, p.edges[0].From())
	for _, e := range p.edges {
		names = append(names, e.To())
	}
	return names
}

// IsCircular returns true if adding edge would create a cycle.
func (p *Path) IsCircular(edge Edge) bool {
	child := edge.To()
	for _, e := range p.edges {
		if e.From() == child || e.To() == child {
			return true
		}
	}
	return false
}

// HasNode returns true if the path traverses the given node.
func (p *Path) HasNode(s string) bool {
	for _, e := range p.edges {
		if e.From() == s || e.To() == s {
			return true
		}
	}
	return false
}
