package treemap

import (
	"math"
	"sort"
)

type Box struct {
	X float64
	Y float64
	W float64
	H float64
}

// Squarify partitions box into parts by using Squarify algorithm.
// Returns boxes in same order as areas.
// As described in "Squarified Treemaps", Mark Bruls, Kees Huizing, and Jarke J. van Wijk., 2000
func Squarify(box Box, areas []float64) []Box {
	// sort
	sortedAreas := make([]float64, len(areas))
	copy(sortedAreas, areas)
	sort.Slice(sortedAreas, func(i, j int) bool { return sortedAreas[i] > sortedAreas[j] })

	// normalize
	var totalArea float64
	for _, s := range sortedAreas {
		totalArea += s
	}
	if boxArea := box.W * box.H; totalArea != boxArea {
		for i, s := range sortedAreas {
			sortedAreas[i] = boxArea * s / totalArea
		}
	}

	layout := squarifyBoxLayout{
		boxes:     nil,
		freeSpace: box,
	}

	layout.squarify(sortedAreas, nil, math.Min(layout.freeSpace.W, layout.freeSpace.H))

	boxes := layout.boxes
	cutoffOverflows(box, layout.boxes)

	return boxes
}

// squarifyBoxLayout defines how to partition BoundingBox into boxes
type squarifyBoxLayout struct {
	boxes     []Box // fixed boxes that have been positioned and sized already
	freeSpace Box   // free space that is left. will be used to fill out with new boxes
}

// squarify expects normalized areas that add up to free space
func (l *squarifyBoxLayout) squarify(unassignedAreas []float64, stackAreas []float64, w float64) {
	if len(unassignedAreas) == 0 {
		l.stackBoxes(stackAreas)
		return
	}

	if len(stackAreas) == 0 {
		l.squarify(unassignedAreas[1:], []float64{unassignedAreas[0]}, w)
		return
	}

	c := unassignedAreas[0]
	if stackc := append(stackAreas, c); highestAspectRatio(stackAreas, w) > highestAspectRatio(stackc, w) {
		// aspect ratio improves, add it to current stack
		l.squarify(unassignedAreas[1:], stackc, w)
	} else {
		// aspect ratio does not improve
		l.stackBoxes(stackAreas)
		l.squarify(unassignedAreas, nil, math.Min(l.freeSpace.W, l.freeSpace.H))
	}
}

// stackBoxes makes new boxes accordingly to areas and fix them into freeSpacelayout within bounding box
func (l *squarifyBoxLayout) stackBoxes(stackAreas []float64) {
	if l.freeSpace.W < l.freeSpace.H {
		l.stackBoxesHorizontal(stackAreas)
	} else {
		l.stackBoxesVertical(stackAreas)
	}
}

// stackBoxesVertical takes vertical chunk of free space of bounding box and partitiones it into areas
func (l *squarifyBoxLayout) stackBoxesVertical(areas []float64) {
	if len(areas) == 0 {
		return
	}

	stackArea := 0.0
	for _, s := range areas {
		stackArea += s
	}
	if stackArea == 0 {
		return
	}

	totalArea := l.freeSpace.W * l.freeSpace.H
	if totalArea == 0 {
		return
	}

	// stack
	offset := l.freeSpace.Y
	for _, s := range areas {
		h := l.freeSpace.H * s / stackArea
		b := Box{
			X: l.freeSpace.X,
			W: l.freeSpace.W * stackArea / totalArea,
			Y: offset,
			H: h,
		}
		offset += h
		l.boxes = append(l.boxes, b)
	}

	// shrink free space
	l.freeSpace = Box{
		X: l.freeSpace.X + (l.freeSpace.W * stackArea / totalArea),
		W: l.freeSpace.W * (1 - (stackArea / totalArea)),
		Y: l.freeSpace.Y,
		H: l.freeSpace.H,
	}
}

// stackBoxesHorizontal takes horizontal chunk of free space of bounding box and partitiones it into areas
func (l *squarifyBoxLayout) stackBoxesHorizontal(areas []float64) {
	if len(areas) == 0 {
		return
	}

	stackArea := 0.0
	for _, s := range areas {
		stackArea += s
	}
	if stackArea == 0 {
		return
	}

	totalArea := l.freeSpace.W * l.freeSpace.H
	if totalArea == 0 {
		return
	}

	// stack
	offset := l.freeSpace.X
	for _, s := range areas {
		w := l.freeSpace.W * s / stackArea
		b := Box{
			X: offset,
			W: w,
			Y: l.freeSpace.Y,
			H: l.freeSpace.H * stackArea / totalArea,
		}
		offset += w
		l.boxes = append(l.boxes, b)
	}

	// shrink free space
	l.freeSpace = Box{
		X: l.freeSpace.X,
		W: l.freeSpace.W,
		Y: l.freeSpace.Y + (l.freeSpace.H * stackArea / totalArea),
		H: l.freeSpace.H * (1 - (stackArea / totalArea)),
	}
}

// highestAspectRatio of a list of rectangles's areas, given the length of the side along which they are to be laid out
func highestAspectRatio(areas []float64, w float64) float64 {
	var minArea, maxArea, totalArea float64
	for i, s := range areas {
		totalArea += s
		if i == 0 || s < minArea {
			minArea = s
		}
		if i == 0 || s > maxArea {
			maxArea = s
		}
	}

	v1 := w * w * maxArea / (totalArea * totalArea)
	v2 := totalArea * totalArea / (w * w * minArea)

	return math.Max(v1, v2)
}

// cutoffOverflows will set boxes that overflow to fit into bounding box.
// This is useful for numerical stability on the borders.
func cutoffOverflows(boundingBox Box, boxes []Box) {
	maxX := boundingBox.X + boundingBox.W
	maxY := boundingBox.Y + boundingBox.H

	for i, b := range boxes {
		if delta := (b.X + b.W) - maxX; delta > 0 {
			boxes[i].W -= delta
		}
		if delta := (b.Y + b.H) - maxY; delta > 0 {
			boxes[i].H -= delta
		}
	}
}
