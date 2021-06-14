package agent

import "github.com/z24tao/sapphire_core/world"

/*
	In the current version, instead of being implemented as a concept, attributes are simply
	kept as constants for simplicity.
*/

const (
	attrTypeColor = iota
	attrTypeShape
	attrTypeDirection
	attrTypeDistance
)

const (
	directionOrigin = iota
	directionXPos
	directionXNeg
	directionZPos
	directionZNeg
	directionXPosZPos
	directionXPosZNeg
	directionXNegZPos
	directionXNegZNeg
)

var visualAttrTypes = map[int]bool{
	attrTypeColor: true,
	attrTypeShape: true,
}

// type -> name
var attrTypes = map[int]string{
	attrTypeColor:     "color",
	attrTypeShape:     "shape",
	attrTypeDirection: "direction",
	attrTypeDistance: "distance",
}

// type -> value -> name
var attrVals = map[int]map[int]string{
	attrTypeColor: world.Colors,
	attrTypeShape: world.Shapes,
	attrTypeDirection: {
		directionOrigin:   "origin",
		directionXPos:     "right",
		directionXNeg:     "left",
		directionZPos:     "front",
		directionZNeg:     "back",
		directionXPosZPos: "right/front",
		directionXPosZNeg: "right/back",
		directionXNegZPos: "left/front",
		directionXNegZNeg: "left/back",
	},
}

var qualitativeAttrTypes = map[int]bool{
	attrTypeColor:     true,
	attrTypeShape:     true,
	attrTypeDirection: true,
}

var quantitativeAttrTypes = map[int]bool{
	attrTypeDistance: true,
}
