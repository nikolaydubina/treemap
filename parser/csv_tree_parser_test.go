package parser

import (
	"sort"
	"strings"
	"testing"

	"github.com/nikolaydubina/treemap"
)

func TestMakeTree(t *testing.T) {
	tests := []struct {
		name    string
		nodes   []treemap.Node
		expTree *treemap.Tree
		expErr  string
	}{
		{
			name: "one deep node",
			nodes: []treemap.Node{
				{Path: "a/b/c"},
			},
			expTree: &treemap.Tree{
				Nodes: map[string]treemap.Node{
					"a":     {Path: "a"},
					"a/b":   {Path: "a/b"},
					"a/b/c": {Path: "a/b/c"},
				},
				To: map[string][]string{
					"a":   {"a/b"},
					"a/b": {"a/b/c"},
				},
				Root: "a",
			},
		},
		{
			name: "multiple deep nodes",
			nodes: []treemap.Node{
				{Path: "a/b/c"},
				{Path: "a/b/c/d"},
				{Path: "a/b/d"},
			},
			expTree: &treemap.Tree{
				Nodes: map[string]treemap.Node{
					"a":       {Path: "a"},
					"a/b":     {Path: "a/b"},
					"a/b/c":   {Path: "a/b/c"},
					"a/b/c/d": {Path: "a/b/c/d"},
					"a/b/d":   {Path: "a/b/d"},
				},
				To: map[string][]string{
					"a":     {"a/b"},
					"a/b":   {"a/b/c", "a/b/d"},
					"a/b/c": {"a/b/c/d"},
				},
				Root: "a",
			},
		},
		{
			name: "when has leading slash, then has empty string as root",
			nodes: []treemap.Node{
				{Path: "/a/b/c"},
			},
			expTree: &treemap.Tree{
				Nodes: map[string]treemap.Node{
					"":       {Path: ""},
					"/a":     {Path: "/a"},
					"/a/b":   {Path: "/a/b"},
					"/a/b/c": {Path: "/a/b/c"},
				},
				To: map[string][]string{
					"":     {"/a"},
					"/a":   {"/a/b"},
					"/a/b": {"/a/b/c"},
				},
				Root: "",
			},
		},
		{
			name:    "when no roots, then error",
			nodes:   []treemap.Node{},
			expTree: nil,
			expErr:  "cycle",
		},
		{
			name: "when two roots, then making fake root",
			nodes: []treemap.Node{
				{Path: "a/b"},
				{Path: "b/d"},
			},
			expTree: &treemap.Tree{
				Nodes: map[string]treemap.Node{
					"a":   {Path: "a"},
					"a/b": {Path: "a/b"},
					"b":   {Path: "b"},
					"b/d": {Path: "b/d"},
				},
				To: map[string][]string{
					"a":                  {"a/b"},
					"b":                  {"b/d"},
					"some-secret-string": {"a", "b"},
				},
				Root: "some-secret-string",
			},
		},
		{
			name: "when duplicate nodes, then overrides latest",
			nodes: []treemap.Node{
				{Path: "a/b"},
				{Path: "a/b"},
			},
			expTree: &treemap.Tree{
				Nodes: map[string]treemap.Node{
					"a":   {Path: "a", Name: "", Size: 0, Heat: 0, HasHeat: false},
					"a/b": {Path: "a/b", Name: "", Size: 0, Heat: 0, HasHeat: false},
				},
				To: map[string][]string{
					"a": {"a/b"},
				},
				Root: "a",
			},
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
				} else {
					if !strings.Contains(err.Error(), tc.expErr) {
						t.Error("error does not contain expected string")
					}
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

func TestParseNodes(t *testing.T) {
	tests := []struct {
		name     string
		in       string
		expNodes []treemap.Node
		expErr   string
	}{
		{
			name: "when basic case, then works",
			in:   "a/b/c,10,11",
			expNodes: []treemap.Node{
				{
					Path:    "a/b/c",
					Size:    10,
					Heat:    11,
					HasHeat: true,
				},
			},
		},
		{
			name:   "when wrong number, then error",
			in:     ",,\n\n",
			expErr: "is not float",
		},
		{
			name:   "when wrong number, then error",
			in:     ",1,\n\n",
			expErr: "is not float",
		},
		{
			name:   "when wrong quotation, then error",
			in:     "\"\n\n",
			expErr: "can not parse",
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			nodes, err := parseNodes(tc.in)

			assertError(t, err, tc.expErr)

			if len(tc.expNodes) != len(nodes) {
				t.Error("wrong len")
			}
			for i := range nodes {
				if tc.expNodes[i] != nodes[i] {
					t.Error("wrong node")
				}
			}
		})
	}
}

func assertError(t *testing.T, err error, expErr string) {
	if expErr == "" && err != nil {
		t.Error(err)
	}
	if expErr != "" {
		if err == nil {
			t.Error("expected error, got nil")
		}
		if !strings.Contains(err.Error(), expErr) {
			t.Errorf("error does not contain expected string got(%s)", err)
		}
	}
}

func eqTree(a, b treemap.Tree) bool {
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
