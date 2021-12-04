package render

import (
	"testing"

	"github.com/nikolaydubina/treemap"
)

func TestTreeHues(t *testing.T) {
	tests := []struct {
		name    string
		tree    treemap.Tree
		offset  float64
		expHues map[string]float64
	}{
		{
			name:   "when single child, then same as parent",
			offset: 0,
			tree: treemap.Tree{
				To: map[string][]string{
					"a": {"a/b"},
				},
				Nodes: map[string]treemap.Node{
					"a/b": {Path: "a/b"},
				},
				Root: "a",
			},
			expHues: map[string]float64{
				"a":   180,
				"a/b": 180,
			},
		},
		{
			name:   "when two, then diametrally opposites",
			offset: 0,
			tree: treemap.Tree{
				To: map[string][]string{
					"a": {"a/b", "a/c"},
				},
				Nodes: map[string]treemap.Node{
					"a/b": {Path: "a/b"},
					"a/c": {Path: "a/c"},
				},
				Root: "a",
			},
			expHues: map[string]float64{
				"a":   180,
				"a/b": 90,
				"a/c": 270,
			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			hues := TreeHues(tc.tree, tc.offset)

			if len(hues) != len(tc.expHues) {
				t.Errorf("len: exp(%d) != got(%d)", len(tc.expHues), len(hues))
			}
			for i := range hues {
				if tc.expHues[i] != hues[i] {
					t.Errorf("%v: exp(%v) != got(%v)", i, tc.expHues[i], hues[i])
				}
			}
		})
	}
}
