package agent

import (
	"github.com/z24tao/sapphire_core/world"
)

// activity - Active actions and atomic action types
type activity struct {
	atomicActionInterfaces map[*world.AtomicActionInterface]*atomicActionType
	activeActions          []action
}

//TODO: figure out what this does...

// newActivity -
func newActivity(a *Agent) *activity {
	newAAIs := world.NewAAIs(a.unitId)

	ac := &activity{
		atomicActionInterfaces: map[*world.AtomicActionInterface]*atomicActionType{},
	}

	for aai := range newAAIs {
		aat := a.newAtomicActionType(aai)
		ac.atomicActionInterfaces[aai] = aat
		a.mind.addItem(aat, 1.0)
	}
	a.mind.mergeNewThoughts()

	return ac
}
