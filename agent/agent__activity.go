package agent

import (
	"github.com/z24tao/sapphire_core/world"
)

type activity struct {
	atomicActionInterfaces map[*world.AtomicActionInterface]*atomicActionType
	activeActions          []action
}

func newActivity(a *Agent) *activity {
	newAAIs := world.NewAAIs(a.unitId)

	ac := &activity{
		atomicActionInterfaces: map[*world.AtomicActionInterface]*atomicActionType{},
	}

	for aai := range newAAIs {
		aat := newAtomicActionType(aai)
		ac.atomicActionInterfaces[aai] = aat
		a.mind.addItem(aat, 1)
	}
	a.mind.mergeNewThoughts()

	return ac
}
