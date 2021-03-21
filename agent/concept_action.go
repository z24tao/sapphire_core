package agent

import (
	"../world"
	"fmt"
)

const (
	actionStateIdle = iota
	actionStateActive
	actionStatePaused
	actionStateComplete
)

const curiosityValue = 50.0

type action interface {
	concept
	// assume subject is self -- TODO add subject
	_type() actionType
	_getObject() object
	_setObject(o object)
	_start(agent *Agent) bool // returns success
	_step(agent *Agent)
	_pause(agent *Agent) bool // returns success
	_stop(agent *Agent) bool  // returns success
	_state() int
}

type actionType interface {
	concept
	conceptType
	_objectType() objectType
	_startConditions() []condition
	_pauseConditions() []condition
	_stopConditions() []condition
	_endConditions() []condition
	_getCausations() []*causation
	_setCausations([]*causation)
	_attempts() int
}

func (a *Agent) _processActionResponse(response interface{}) {
	fmt.Println("_processActionResponse")
	if taste, ok := response.(*world.Taste); ok {
		a._processTaste(taste)
	} else if aaiChange, ok := response.(*world.AtomicActionInterfaceChange); ok {
		a._processAAIChange(aaiChange)
	}
}

func (a *Agent) _processTaste(taste *world.Taste) {
	fmt.Println("_processTaste")
	fmt.Println("taste.sweet", taste.Sweet)
	fmt.Println("taste.nutrition", taste.Nutrition)
	a._state._update(agentStateTypeHunger, -taste.Nutrition)
	if taste.Sweet {
		fmt.Println("oh it's also sweet")
		a._state._update(agentExperienceTypeSweet, 0)
	}
}

func (a *Agent) _processAAIChange(c *world.AtomicActionInterfaceChange) {
	fmt.Println("_processAAIChange")
	a._updateActionOutcome([]change{&atomicActionInterfaceChange{
		 __interface: c.Interface,
		 __enabling: c.Enabling,
		 //__value:
	}})
}

func actionTypeValue(t actionType) float64 {
	if len(t._getCausations()) <= 0 {
		return curiosityValue
	}

	v := 0.0
	for _, c := range t._getCausations() {
		v += c._change._getValue() * float64(c._occurrences) / float64(t._attempts())
	}

	return v
}
