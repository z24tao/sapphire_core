package agent

import (
	"../world"
	"fmt"
)

type Agent struct {
	_self  *simpleObject
	_state *agentState

	_mind     *mind
	_activity *activity

	_unitId int
}

func (a *Agent) TimeStep() {
	a._updateState()
	a._observe() // observe and learn
	a._act() // think and act

	fmt.Println(a._mind._toString())
}

func NewAgent() *Agent {
	a := &Agent{
		_mind:     newMind(),
		_unitId:   world.NewActor(),
	}

	a._activity = newActivity(a)
	a._state = newAgentState(a)
	return a
}

func init() {

}
