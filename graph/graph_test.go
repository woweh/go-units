package graph

import (
	"reflect"
	"testing"
)

type TestEdge struct {
	from string
	to   string
}

func (t TestEdge) To() string   { return t.to }
func (t TestEdge) From() string { return t.from }

func TestBFSTree_Basic(t *testing.T) {
	tree := New(
		TestEdge{"a", "b"},
		TestEdge{"b", "d"},
		TestEdge{"d", "e"},
		TestEdge{"e", "f"},
		TestEdge{"f", "g"},
		TestEdge{"c", "d"},
		TestEdge{"c", "a"},
		TestEdge{"a", "c"},
		TestEdge{"b", "c"},
	)
	if tree.Len() != 9 {
		t.Errorf("expected 9 edges, got %d", tree.Len())
	}
	nodes := tree.Nodes()
	if len(nodes) != 7 {
		t.Errorf("expected 7 nodes, got %d", len(nodes))
	}
	if !tree.HasNode("a") || !tree.HasNode("g") {
		t.Error("expected HasNode to find 'a' and 'g'")
	}
	if !tree.HasEdge(TestEdge{"a", "b"}) {
		t.Error("expected HasEdge to find edge a->b")
	}
}

func TestBFSTree_FindPath(t *testing.T) {
	tree := New(
		TestEdge{"a", "b"},
		TestEdge{"b", "c"},
		TestEdge{"c", "d"},
		TestEdge{"d", "e"},
	)
	path, err := tree.FindPath("a", "e")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	exp := []string{"a", "b", "c", "d", "e"}
	if got := path.Nodes(); !reflect.DeepEqual(got, exp) {
		t.Errorf("expected path %v, got %v", exp, got)
	}
	if path.String() != "a->b->c->d->e" {
		t.Errorf("unexpected path string: %s", path.String())
	}
}

func TestBFSTree_FindShortPath(t *testing.T) {
	tree := New(TestEdge{"x", "y"})
	path, err := tree.FindPath("x", "y")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if path.String() != "x->y" {
		t.Errorf("unexpected path string: %s", path.String())
	}
}

func TestBFSTree_FindNoPath(t *testing.T) {
	tree := New(TestEdge{"a", "b"})
	_, err := tree.FindPath("a", "z")
	if err == nil {
		t.Error("expected error for missing path")
	}
}

func TestBFSTree_Circular(t *testing.T) {
	tree := New(
		TestEdge{"a", "b"},
		TestEdge{"b", "a"},
	)
	path, err := tree.FindPath("a", "b")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if path.IsCircular(TestEdge{"b", "a"}) != true {
		t.Error("expected IsCircular to detect cycle")
	}
}

func TestBFSTree_ResetAndCopy(t *testing.T) {
	tree := New(TestEdge{"a", "b"}, TestEdge{"b", "c"})
	copyTree := tree.Copy()
	if !reflect.DeepEqual(tree.Edges(), copyTree.Edges()) {
		t.Error("expected Copy to produce identical edges")
	}
	tree.Reset()
	if tree.Len() != 0 {
		t.Error("expected Reset to clear edges")
	}
	if copyTree.Len() != 2 {
		t.Error("expected Copy to be unaffected by Reset")
	}
}

func TestBFSTree_NilEdgeHandling(t *testing.T) {
	tree := New(nil, TestEdge{"a", "b"}, nil)
	if tree.Len() != 1 {
		t.Errorf("expected 1 edge, got %d", tree.Len())
	}
}

func TestPath_Last_Error(t *testing.T) {
	p := New()
	path := New()
	_, err := path.FindPath("a", "b")
	if err == nil {
		t.Error("expected error for missing path")
	}
	_, err = p.FindPath("a", "a")
	if err == nil {
		t.Error("expected error for same start and end")
	}
}
