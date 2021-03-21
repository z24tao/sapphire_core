package agent

const (
	attrTypeColor = iota
	attrTypeShape
	attrTypeDirection
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
