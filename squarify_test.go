package treemap

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSquarifystackBoxesVertical(t *testing.T) {
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
			areas:     []float64{5, 5},
			leaveArea: 6 * 4 * 0.5,
			expLayout: squarifyBoxLayout{
				boxes: []Box{
					{X: 0, Y: 0, W: 2.5, H: 2},
					{X: 0, Y: 2, W: 2.5, H: 2},
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
			areas:     []float64{5},
			leaveArea: 6 * 4 * 0.5,
			expLayout: squarifyBoxLayout{
				boxes: []Box{
					{X: 0, Y: 0, W: 1.25, H: 4},
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
			areas:     nil,
			leaveArea: 6 * 4 * 0.5,
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
			areas:     []float64{0, 0, 0},
			leaveArea: 6 * 4 * 0.5,
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
			areas:     []float64{1},
			leaveArea: 1,
			expLayout: squarifyBoxLayout{
				freeSpace: Box{W: 0, H: 4},
			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc.layout.stackBoxesVertical(tc.areas, tc.leaveArea)
			assert.Equal(t, tc.expLayout, tc.layout)
		})
	}
}

func TestSquarifystackBoxesHorizontal(t *testing.T) {
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
			areas:     []float64{5, 5},
			leaveArea: 6 * 4 * 0.5,
			expLayout: squarifyBoxLayout{
				boxes: []Box{
					{X: 0, Y: 0, W: 2, H: 2.5},
					{X: 2, Y: 0, W: 2, H: 2.5},
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
			areas:     []float64{5},
			leaveArea: 6 * 4 * 0.5,
			expLayout: squarifyBoxLayout{
				boxes: []Box{
					{X: 0, Y: 0, W: 4, H: 1.25},
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
			areas:     nil,
			leaveArea: 6 * 4 * 0.5,
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
			areas:     []float64{0, 0, 0},
			leaveArea: 6 * 4 * 0.5,
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
			areas:     []float64{1},
			leaveArea: 1,
			expLayout: squarifyBoxLayout{
				freeSpace: Box{W: 0, H: 4},
			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc.layout.stackBoxesHorizontal(tc.areas, tc.leaveArea)
			assert.Equal(t, tc.expLayout, tc.layout)
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
				{X: 3, Y: 1.6666666666666667, W: 1.2, H: 1.6666666666666667},                  // stack-3 horizontal
				{X: 4.8, Y: 1.6666666666666667, W: 1.2, H: 1.6666666666666667},                // stack-3 horizontal
				{X: 5.3999999999999995, Y: 1.6666666666666667, W: 0.6, H: 1.6666666666666667}, // stack-3 horizontal
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
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			boxes := Squarify(tc.box, tc.areas)

			if len(tc.expBoxes) != len(boxes) {
				t.Errorf("exp(%#v) != got(%#v)", tc.expBoxes, boxes)
			}
			for i, b := range tc.expBoxes {
				if tc.expBoxes[i] != boxes[i] {
					t.Errorf("exp(%#v) != got(%#v)", tc.expBoxes, boxes)
				}
				if (b.H * b.W) < 0.1 {
					t.Errorf("got wrong size for box(%d: %#v)", i, b)
				}
			}
		})
	}
}
