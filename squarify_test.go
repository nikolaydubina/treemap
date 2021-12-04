package treemap

import (
	"fmt"
	"testing"
)

func TestSquarifyStackBoxesVertical(t *testing.T) {
	tests := []struct {
		name      string
		layout    squarifyBoxLayout
		areas     []float64
		leaveArea float64
		expLayout squarifyBoxLayout
	}{
		{
			name: "when leaving 50% of space and equal parts, then space is set correctly and stacked correctly",
			layout: squarifyBoxLayout{
				boxes:     nil,
				freeSpace: Box{W: 6, H: 4},
			},
			areas: []float64{6, 6},
			expLayout: squarifyBoxLayout{
				boxes: []Box{
					{X: 0, Y: 0, W: 3, H: 2},
					{X: 0, Y: 2, W: 3, H: 2},
				},
				freeSpace: Box{X: 3, Y: 0, W: 3, H: 4},
			},
		},
		{
			name: "when leaving 50% of space and one part, then space is correctly and one part takes all hight",
			layout: squarifyBoxLayout{
				boxes:     nil,
				freeSpace: Box{W: 6, H: 4},
			},
			areas: []float64{12},
			expLayout: squarifyBoxLayout{
				boxes: []Box{
					{X: 0, Y: 0, W: 3, H: 4},
				},
				freeSpace: Box{X: 3, Y: 0, W: 3, H: 4},
			},
		},
		{
			name: "when no areas, then leave everything empty",
			layout: squarifyBoxLayout{
				boxes:     nil,
				freeSpace: Box{W: 6, H: 4},
			},
			areas: nil,
			expLayout: squarifyBoxLayout{
				boxes:     nil,
				freeSpace: Box{W: 6, H: 4},
			},
		},
		{
			name: "when area sum is zero, then leave everything empty",
			layout: squarifyBoxLayout{
				boxes:     nil,
				freeSpace: Box{W: 6, H: 4},
			},
			areas: []float64{0, 0, 0},
			expLayout: squarifyBoxLayout{
				boxes:     nil,
				freeSpace: Box{W: 6, H: 4},
			},
		},
		{
			name: "when no free area, then skip",
			layout: squarifyBoxLayout{
				freeSpace: Box{W: 0, H: 4},
			},
			areas: []float64{1},
			expLayout: squarifyBoxLayout{
				freeSpace: Box{W: 0, H: 4},
			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc.layout.stackBoxesVertical(tc.areas)
			if !eqSliceBox(tc.expLayout.boxes, tc.layout.boxes) {
				t.Errorf("wrong boxes: exp(%#v) != got(%#v)", tc.expLayout.boxes, tc.layout.boxes)
			}
			if tc.expLayout.freeSpace != tc.layout.freeSpace {
				t.Errorf("wrong free space: exp(%#v) != got(%#v)", tc.expLayout.freeSpace, tc.layout.freeSpace)
			}
		})
	}
}

func TestSquarifyStackBoxesHorizontal(t *testing.T) {
	tests := []struct {
		name      string
		layout    squarifyBoxLayout
		areas     []float64
		leaveArea float64
		expLayout squarifyBoxLayout
	}{
		{
			name: "when leaving 50% of space and equal parts, then space is set correctly and stacked correctly",
			layout: squarifyBoxLayout{
				boxes:     nil,
				freeSpace: Box{W: 4, H: 6},
			},
			areas: []float64{6, 6},
			expLayout: squarifyBoxLayout{
				boxes: []Box{
					{X: 0, Y: 0, W: 2, H: 3},
					{X: 2, Y: 0, W: 2, H: 3},
				},
				freeSpace: Box{X: 0, Y: 3, W: 4, H: 3},
			},
		},
		{
			name: "when leaving 50% of space and one part, then space is correctly and one part takes all hight",
			layout: squarifyBoxLayout{
				boxes:     nil,
				freeSpace: Box{W: 4, H: 6},
			},
			areas: []float64{12},
			expLayout: squarifyBoxLayout{
				boxes: []Box{
					{X: 0, Y: 0, W: 4, H: 3},
				},
				freeSpace: Box{X: 0, Y: 3, W: 4, H: 3},
			},
		},
		{
			name: "when no areas, then leave everything empty",
			layout: squarifyBoxLayout{
				boxes:     nil,
				freeSpace: Box{W: 4, H: 6},
			},
			areas: nil,
			expLayout: squarifyBoxLayout{
				boxes:     nil,
				freeSpace: Box{W: 4, H: 6},
			},
		},
		{
			name: "when area sum is zero, then leave everything empty",
			layout: squarifyBoxLayout{
				boxes:     nil,
				freeSpace: Box{W: 4, H: 6},
			},
			areas: []float64{0, 0, 0},
			expLayout: squarifyBoxLayout{
				boxes:     nil,
				freeSpace: Box{W: 4, H: 6},
			},
		},
		{
			name: "when no free area, then skip",
			layout: squarifyBoxLayout{
				freeSpace: Box{W: 0, H: 4},
			},
			areas: []float64{1},
			expLayout: squarifyBoxLayout{
				freeSpace: Box{W: 0, H: 4},
			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc.layout.stackBoxesHorizontal(tc.areas)
			if !eqSliceBox(tc.expLayout.boxes, tc.layout.boxes) {
				t.Errorf("wrong boxes: exp(%#v) != got(%#v)", tc.expLayout.boxes, tc.layout.boxes)
			}
			if tc.expLayout.freeSpace != tc.layout.freeSpace {
				t.Errorf("wrong free space: exp(%#v) != got(%#v)", tc.expLayout.freeSpace, tc.layout.freeSpace)
			}
		})
	}
}

func TestSquarifyHighestAspectRatioPaper(t *testing.T) {
	tests := []struct {
		areasA []float64
		areasB []float64
		w      float64
	}{
		{
			areasA: []float64{6, 6},
			areasB: []float64{6},
			w:      4,
		},
		{
			areasA: []float64{6, 6},
			areasB: []float64{6, 6, 4},
			w:      4,
		},
	}
	for i, tc := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			a := highestAspectRatio(tc.areasA, tc.w)
			b := highestAspectRatio(tc.areasB, tc.w)
			if a > b {
				t.Errorf("a(%f) ! < b(%f)", a, b)
			}
		})
	}
}

func TestSquarifyCutoffOverflows(t *testing.T) {
	tests := []struct {
		boundingBox Box
		boxes       []Box
		expBoxes    []Box
	}{
		{
			boundingBox: Box{W: 4, Y: 4},
			boxes:       []Box{{X: 2.9999, W: 1.1, Y: 2.9999, H: 1.1}},
			expBoxes:    []Box{{X: 2.9999, W: 1.0001000000000002, Y: 2.9999, H: 1.0001000000000002}},
		},
	}
	for i, tc := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			cutoffOverflows(tc.boundingBox, tc.boxes)
			if !eqSliceBox(tc.expBoxes, tc.boxes) {
				t.Errorf("wrong boxes: exp(%#v) != got(%#v)", tc.expBoxes, tc.boxes)
			}
		})
	}
}

func TestSquarify(t *testing.T) {
	tests := []struct {
		name     string
		box      Box
		areas    []float64
		expBoxes []Box
	}{
		{
			name:  "when example from original paper, then it has same layout",
			box:   Box{W: 6, H: 4},
			areas: []float64{6, 6, 4, 3, 2, 2, 1},
			expBoxes: []Box{
				{X: 0, Y: 0, W: 3, H: 2},                                                      // stack-1 vertical
				{X: 0, Y: 2, W: 3, H: 2},                                                      // stack-1 vertical
				{X: 3, Y: 0, W: 1.7142857142857142, H: 2.3333333333333335},                    // stack-2 horizontal
				{X: 4.714285714285714, Y: 0, W: 1.2857142857142858, H: 2.3333333333333335},    // stack-2 horizontal
				{X: 3, Y: 2.3333333333333335, W: 1.2, H: 1.6666666666666665},                  // stack-3 horizontal
				{X: 4.2, Y: 2.3333333333333335, W: 1.2, H: 1.6666666666666665},                // stack-3 horizontal
				{X: 5.4, Y: 2.3333333333333335, W: 0.5999999999999998, H: 1.6666666666666663}, // stack-3 horizontal
			},
		},
		{
			name:  "when one box, then take all area",
			box:   Box{W: 6, H: 4},
			areas: []float64{24},
			expBoxes: []Box{
				{X: 0, Y: 0, W: 6, H: 4},
			},
		},
		{
			name:  "when two boxes same area, then split into half and half horizontally",
			box:   Box{W: 6, H: 4},
			areas: []float64{12, 12},
			expBoxes: []Box{
				{X: 0, Y: 0, W: 3, H: 4},
				{X: 3, Y: 0, W: 3, H: 4},
			},
		},
		{
			name:  "when three boxes same area and wide bounding, then split equally along long dimension",
			box:   Box{W: 12, H: 3},
			areas: []float64{12, 12, 12},
			expBoxes: []Box{
				{X: 0, Y: 0, W: 4, H: 3},
				{X: 4, Y: 0, W: 4, H: 3},
				{X: 8, Y: 0, W: 4, H: 3},
			},
		},
		{
			name:  "when need to normalize, then normalization happens for correct result",
			box:   Box{W: 12, H: 3},
			areas: []float64{2, 2, 2},
			expBoxes: []Box{
				{X: 0, Y: 0, W: 4, H: 3},
				{X: 4, Y: 0, W: 4, H: 3},
				{X: 8, Y: 0, W: 4, H: 3},
			},
		},
		{
			name:  "when need to normalize and sort, then normalization and sorting happens to make big first for correct result",
			box:   Box{W: 12, H: 3},
			areas: []float64{1, 3},
			expBoxes: []Box{
				{X: 0, Y: 0, W: 9, H: 3},
				{X: 9, Y: 0, W: 3, H: 3},
			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			boxes := Squarify(tc.box, tc.areas)

			if !eqSliceBox(tc.expBoxes, boxes) {
				t.Errorf("wrong boxes: exp(%#v) != got(%#v)", tc.expBoxes, boxes)
			}
			for i, b := range boxes {
				if (b.H * b.W) < 0.1 {
					t.Errorf("got wrong size for box(%d: %#v)", i, b)
				}
				if b.X < 0 || b.Y < 0 || ((b.X + b.W) > (tc.box.X + tc.box.W)) || ((b.Y + b.H) > (tc.box.Y + tc.box.H)) {
					t.Errorf("box(%d: %#v) overflows", i, b)
				}
			}
		})
	}
}

func eqSliceBox(a, b []Box) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range b {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
