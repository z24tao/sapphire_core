package agent

import (
	"../world"
)

type activity struct {
	_atomicActionInterfaces map[*world.AtomicActionInterface]bool
	_activeActions          []action
}

func newActivity(a *Agent) *activity {
	ac := &activity{
		_atomicActionInterfaces: world.NewAAIs(a._unitId),
	}

	for aai := range ac._atomicActionInterfaces {
		aat := &atomicActionType {
			__assocs: []concept{},
			__objType: nil,
			__causations: []*causation{},
			__attempts: 0,
			__interface: aai,
		}
		a._mind._addItem(aat, 1)
	}
	a._mind._mergeNewThoughts()

	return ac
}
