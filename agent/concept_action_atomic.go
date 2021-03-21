package agent

import (
	"../world"
	"fmt"
)

type atomicAction struct {
	__type   *atomicActionType
	__assocs []concept
	__object object
	__state  int
}

func (a *atomicAction) _type() actionType {
	return a.__type
}

func (a *atomicAction) _assocs() []concept {
	return a.__assocs
}

func (a *atomicAction) _toString() string {
	return "atomicAction"
}

func (a *atomicAction) _getObject() object {
	return a.__object
}

func (a *atomicAction) _setObject(o object) {
	if o == nil {
		a.__object = nil
		return
	}

	if !a.__type.__objType._match(o._attrs(), nil) {
		return
	}

	a.__object = o
}

func (a *atomicAction) _start(agent *Agent) bool {
	if a.__state != actionStateIdle {
		return false
	}

	if a.__type.__objType != nil && a.__object == nil {
		return false
	}

	for _, cond := range a.__type._startConditions() {
		if !cond._isSatisfied(agent) {
			return false
		}
	}

	a.__state = actionStateActive
	a.__type.__attempts ++
	return true
}

func (a *atomicAction) _step(agent *Agent) {
	if a.__type.__objType != nil && a.__object == nil {
		return
	}

	if a.__state == actionStateActive && a.__type.__interface.Enabled {
		responses := a.__type.__interface.Step()

		for _, response := range responses {
			agent._processActionResponse(response)
		}

		a.__state = actionStateComplete
	}
}

func (a *atomicAction) _pause(agent *Agent) bool {
	if a.__state != actionStateActive {
		return false
	}

	for _, cond := range a.__type._pauseConditions() {
		if !cond._isSatisfied(agent) {
			return false
		}
	}

	a.__state = actionStatePaused
	return true
}

func (a *atomicAction) _stop(agent *Agent) bool {
	if a.__state != actionStateActive {
		return false
	}

	for _, cond := range a.__type._stopConditions() {
		if !cond._isSatisfied(agent) {
			return false
		}
	}

	a.__state = actionStateComplete
	return true
}

func (a *atomicAction) _state() int {
	return a.__state
}

type atomicActionType struct {
	__assocs     []concept
	__objType    objectType
	__causations []*causation
	__attempts   int
	__interface  *world.AtomicActionInterface
}

func (t *atomicActionType) _assocs() []concept {
	return t.__assocs
}

func (t *atomicActionType) _toString() string {
	return fmt.Sprintf("atomicActionType (%d causations)", len(t.__causations))
}

func (t *atomicActionType) _instantiate() concept {
	return &atomicAction{
		__type:   t,
		__assocs: t.__assocs,
		__object: nil,
		__state:  actionStateIdle,
	}
}

func (t *atomicActionType) _objectType() objectType {
	return t.__objType
}

func (t *atomicActionType) _startConditions() []condition {
	return []condition{&aaiStartCondition{
		__interface: t.__interface,
	}}
}

func (t *atomicActionType) _pauseConditions() []condition {
	return []condition{&contradictionCondition{}}
}

func (t *atomicActionType) _stopConditions() []condition {
	return []condition{&contradictionCondition{}}
}

func (t *atomicActionType) _endConditions() []condition {
	return []condition{&contradictionCondition{}}
}

func (t *atomicActionType) _getCausations() []*causation {
	return t.__causations
}

func (t *atomicActionType) _setCausations(cs []*causation) {
	t.__causations = cs
}

func (t *atomicActionType) _attempts() int {
	return t.__attempts
}

type atomicActionInterfaceChange struct {
	__assocs    []concept
	__interface *world.AtomicActionInterface
	__value     float64
	__enabling  bool // true: disabled -> enabled, false: enabled -> disabled
}

func (c *atomicActionInterfaceChange) _assocs() []concept {
	return c.__assocs
}

func (c *atomicActionInterfaceChange) _toString() string {
	return "atomicActionInterfaceChange"
}

func (c *atomicActionInterfaceChange) _match(n change) bool {
	nn, ok := n.(*atomicActionInterfaceChange)
	return ok && c.__interface == nn.__interface && c.__enabling == nn.__enabling
}

func (c *atomicActionInterfaceChange) _matchOld(a *Agent, oldItem concept) bool {
	// interface currently disabled -> c.__interface.Enabled == false
	// this is an enabling change -> c.__enabling == true
	// -> matched old
	// complete opposite for disabling change
	return a._activity._atomicActionInterfaces[c.__interface] &&
		c.__interface.Enabled != c.__enabling
}

func (c *atomicActionInterfaceChange) _matchNew(a *Agent, newItem concept) bool {
	return a._activity._atomicActionInterfaces[c.__interface] &&
		c.__interface.Enabled == c.__enabling
}

func (c *atomicActionInterfaceChange) _getValue() float64 {
	return c.__value
}

func (c *atomicActionInterfaceChange) _setValue(newValue float64) {
	c.__value = newValue
}

type aaiStartCondition struct {
	__interface *world.AtomicActionInterface
}

func (c *aaiStartCondition) _isSatisfied(a *Agent) bool {
	return c.__interface.Enabled
}
