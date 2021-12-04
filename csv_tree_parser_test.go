package treemap

import (
	"sort"
	"strings"
	"testing"
)

func TestMakeTree(t *testing.T) {
	tests := []struct {
		name    string
		nodes   []Node
		expTree *Tree
		expErr  string
	}{
		{
			name: "when one node with virtual nodes, then root is correct and no virtual nodes and edges is correct",
			nodes: []Node{
				{Path: "a/b/c"},
			},
			expTree: &Tree{
				Root: "a",
				To: map[string][]string{
					"a/b": {
						"a/b/c",
					},
					"a": {
						"a/b",
					},
				},
				Nodes: map[string]Node{
					"a/b/c": {Path: "a/b/c"},
				},
			},
		},
		{
			name: "when multiple nodes same root and same virtual node, then root is correct and no virtual nodes and edges is correct",
			nodes: []Node{
				{Path: "a/b/c"},
				{Path: "a/b/c/d"},
				{Path: "a/b/d"},
			},
			expTree: &Tree{
				Root: "a",
				To: map[string][]string{
					"a": {
						"a/b",
					},
					"a/b": {
						"a/b/c",
						"a/b/d",
					},
					"a/b/c": {
						"a/b/c/d",
					},
				},
				Nodes: map[string]Node{
					"a/b/c":   {Path: "a/b/c"},
					"a/b/c/d": {Path: "a/b/c/d"},
					"a/b/d":   {Path: "a/b/d"},
				},
			},
		},
		{
			name: "when has leading /, then has empty string as root",
			nodes: []Node{
				{Path: "/a/b/c"},
			},
			expTree: &Tree{
				Root: "",
				To: map[string][]string{
					"": {
						"/a",
					},
					"/a": {
						"/a/b",
					},
					"/a/b": {
						"/a/b/c",
					},
				},
				Nodes: map[string]Node{
					"/a/b/c": {Path: "/a/b/c"},
				},
			},
		},
		{
			name:    "when no roots, then error",
			nodes:   []Node{},
			expTree: nil,
			expErr:  "cycle",
		},
		{
			name: "when two roots, then making fake root",
			nodes: []Node{
				{Path: "a/b"},
				{Path: "b/d"},
			},
			expTree: &Tree{
				Root: "<some-secret-string>",
				To: map[string][]string{
					"<some-secret-string>": {
						"a",
						"b",
					},
					"a": {
						"a/b",
					},
					"b": {
						"b/d",
					},
				},
				Nodes: map[string]Node{
					"a/b": {Path: "a/b"},
					"b/d": {Path: "b/d"},
				},
			},
		},
		{
			name: "when duplicate nodes, then error",
			nodes: []Node{
				{Path: "a/b"},
				{Path: "a/b"},
			},
			expErr: "duplicate",
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tree, err := makeTree(tc.nodes)

			// error
			if tc.expErr == "" && err != nil {
				t.Error(err)
			}
			if tc.expErr != "" {
				if err == nil {
					t.Error("expected error, got nil")
				}
				if !strings.Contains(err.Error(), tc.expErr) {
					t.Error("error does not contain expected string")
				}
			}

			// tree
			if tc.expTree == nil && tree != nil {
				t.Error("got tree not nil, expected nil")
			}
			if tc.expTree != nil {
				if tree == nil {
					t.Error("got tree nil, expected not nil")
					return
				}
				if !eqTree(*tc.expTree, *tree) {
					t.Errorf("tree: exp(%#v) != got(%#v)", tc.expTree, tree)
				}
			}
		})
	}
}

func eqTree(a, b Tree) bool {
	if a.Root != b.Root {
		return false
	}

	// nodes
	if len(a.Nodes) != len(b.Nodes) {
		return false
	}
	for k, v := range a.Nodes {
		if b.Nodes[k] != v {
			return false
		}
	}

	if len(a.To) != len(b.To) {
		return false
	}
	for k, ato := range a.To {
		bto := b.To[k]

		if len(ato) != len(bto) {
			return false
		}
		sort.Slice(ato, func(i, j int) bool { return ato[i] < bto[j] })
		sort.Slice(bto, func(i, j int) bool { return bto[i] < bto[j] })

		for i := range ato {
			if ato[i] != bto[i] {
				return false
			}
		}
	}

	return true
}
