package world

type AtomicActionInterface struct {
	Enabled     bool
	ObjectImage Image
	Step        func() []interface{}
}

type AtomicActionInterfaceChange struct {
	// TODO rename this to sound more like an info
	Interface *AtomicActionInterface
	Enabling bool
}

const (
	aaiXPos = iota
	aaiXNeg
	aaiZPos
	aaiZNeg
	aaiEat
)

var unitAAI = make(map[int]map[*AtomicActionInterface]bool)
var unitAAIAction = make(map[*AtomicActionInterface]int)

func NewAAIs(unitId int) map[*AtomicActionInterface]bool {
	basicAAIs := make(map[*AtomicActionInterface]bool)

	basicAAIs[newAAI(unitId, aaiXPos)] = true
	basicAAIs[newAAI(unitId, aaiXNeg)] = true
	basicAAIs[newAAI(unitId, aaiZPos)] = true
	basicAAIs[newAAI(unitId, aaiZNeg)] = true
	basicAAIs[newAAI(unitId, aaiEat)] = true

	unitAAI[unitId] = basicAAIs
	defaultBoard.updateAAI(units[unitId])
	return basicAAIs
}

func newAAI(unitId, actionId int) *AtomicActionInterface {
	aai := &AtomicActionInterface{
		Enabled: true,
		Step: newAAIStep(unitId, actionId),
	}

	unitAAIAction[aai] = actionId
	return aai
}

func newAAIStep(unitId, actionId int) func() []interface{} {
	if actionId == aaiXPos {
		return func() []interface{} {
			defaultBoard.moveUnit(units[unitId], 1, 0)
			return defaultBoard.updateAAI(units[unitId])
		}
	}

	if actionId == aaiXNeg {
		return func() []interface{} {
			defaultBoard.moveUnit(units[unitId], -1, 0)
			return defaultBoard.updateAAI(units[unitId])
		}
	}

	if actionId == aaiZPos {
		return func() []interface{} {
			defaultBoard.moveUnit(units[unitId], 0, 1)
			return defaultBoard.updateAAI(units[unitId])
		}
	}

	if actionId == aaiZNeg {
		return func() []interface{} {
			defaultBoard.moveUnit(units[unitId], 0, -1)
			return defaultBoard.updateAAI(units[unitId])
		}
	}

	if actionId == aaiEat {
		return func() []interface{} {
			response := defaultBoard.unitEat(units[unitId])
			response = append(response, defaultBoard.updateAAI(units[unitId]))
			return response
		}
	}
	return nil
}
