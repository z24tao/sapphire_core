package world

import "fmt"

type AtomicActionInterface struct {
	Name        string
	Enabled     bool
	ObjectImage Image
	Step        func() []interface{}
}

type AtomicActionInterfaceChange struct {
	// TODO rename this to sound more like an info
	Interface *AtomicActionInterface
	Enabling  bool
}

func (a *AtomicActionInterfaceChange) ToString() string {
	return "AtomicActionInterfaceChange " + a.Interface.Name + ", Enabling: " + fmt.Sprint(a.Enabling)
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

	basicAAIs[newAAI(unitId, aaiXPos, "right")] = true
	basicAAIs[newAAI(unitId, aaiXNeg, "left")] = true
	basicAAIs[newAAI(unitId, aaiZPos, "front")] = true
	basicAAIs[newAAI(unitId, aaiZNeg, "back")] = true
	basicAAIs[newAAI(unitId, aaiEat, "eat")] = true

	unitAAI[unitId] = basicAAIs
	defaultBoard.updateAAI(units[unitId])
	return basicAAIs
}

func newAAI(unitId, actionId int, name string) *AtomicActionInterface {
	aai := &AtomicActionInterface{
		Name:    name,
		Enabled: true,
		Step:    newAAIStep(unitId, actionId),
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
			response = append(response, defaultBoard.updateAAI(units[unitId])...)
			return response
		}
	}
	return nil
}
