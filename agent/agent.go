package agent

import (
	"fmt"
	"github.com/z24tao/sapphire_core/world"
	"math/rand"
)

const imageDefaultImportance = 0.5
const changeDefaultImportance = 0.3
const curiosityValue = 50.0

type Agent struct {
	state    *agentState
	mind     *mind
	activity *activity
	unitId   int
}

func (a *Agent) TimeStep() {
	a.updateState()
	a.observe() // observe and learn
	a.act()     // think and act
	fmt.Println(a.mind.toString())
}

func (a *Agent) updateState() {
	for stateKey, stateVal := range a.state.states {
		if stateType, seen := agentStateTypes[stateKey]; seen {
			a.state.states[stateKey] += stateType.perTurn
		} else if stateVal > 0 { // experience fades away if remaining number of turns is positive
			a.state.states[stateKey]--
		}
	}
}

func (a *Agent) observe() {
	imgs := world.Look(a.unitId)
	for _, img := range imgs {
		a.processImage(img)
	}
	a.updateActionCausations()
}

func (a *Agent) processImage(img *world.Image) {
	attrs := a.identifyAttributes(img)
	obj := a.identifyObjectInst(attrs)

	if obj == nil {
		a.createObject(attrs, img.DebugName)
	}
}

func (a *Agent) identifyAttributes(img *world.Image) map[int]int {
	attrs := make(map[int]int)

	attrs[attrTypeColor] = img.Color
	attrs[attrTypeShape] = img.Shape

	if img.XDist == 0 && img.ZDist == 0 {
		attrs[attrTypeDirection] = directionOrigin
	} else if img.XDist > 0 && img.ZDist == 0 {
		attrs[attrTypeDirection] = directionXPos
	} else if img.XDist < 0 && img.ZDist == 0 {
		attrs[attrTypeDirection] = directionXNeg
	} else if img.XDist == 0 && img.ZDist > 0 {
		attrs[attrTypeDirection] = directionZPos
	} else if img.XDist == 0 && img.ZDist < 0 {
		attrs[attrTypeDirection] = directionZNeg
	} else if img.XDist > 0 && img.ZDist > 0 {
		attrs[attrTypeDirection] = directionXPosZPos
	} else if img.XDist > 0 && img.ZDist < 0 {
		attrs[attrTypeDirection] = directionXPosZNeg
	} else if img.XDist < 0 && img.ZDist > 0 {
		attrs[attrTypeDirection] = directionXNegZPos
	} else if img.XDist < 0 && img.ZDist < 0 {
		attrs[attrTypeDirection] = directionXNegZNeg
	}

	return attrs
}

func (a *Agent) identifyObjectInst(attrs map[int]int) object {
	mindObjs := a.mind.objects()
	for _, obj := range mindObjs {
		if obj.match(attrs, visualAttrTypes) {
			changes := obj.setAttrs(attrs)
			a.recordActionChanges(changes)
			return obj
		}
	}

	return nil
}

func (a *Agent) createObject(attrs map[int]int, debugName string) {
	newType := a.identifyObjectType(attrs)
	if newType == nil {
		newType = a.createObjectType(attrs, debugName)
	}

	a.createObjectInst(attrs, newType)
}

func (a *Agent) identifyObjectType(attrs map[int]int) objectType {
	mindObjTypes := a.mind.objectTypes()
	for _, objType := range mindObjTypes {
		if objType.match(attrs, visualAttrTypes) {
			return objType
		}
	}

	return nil
}

func (a *Agent) createObjectType(attrs map[int]int, debugName string) objectType {
	objType := newSimpleObjectType(debugName)

	for attrType, attrVal := range attrs {
		if visualAttrTypes[attrType] {
			objType.attrs[attrType] = attrVal
		}
	}

	a.mind.addItem(objType, imageDefaultImportance)
	return objType
}

func (a *Agent) createObjectInst(attrs map[int]int, newType objectType) object {
	newInst := newType.instantiate().(object)
	newInst.setAttrs(attrs)
	a.mind.addItem(newInst, imageDefaultImportance)
	return newInst
}

func (a *Agent) act() {
	// assumes every object visited during observe is already in mind with importance
	// uses object information and changes to update activity
	a.mind.spawnThoughts()
	a.mind.mergeNewThoughts()
	a.startNewAction()
	a.stepActions()
}

func (a *Agent) startNewAction() {
	actionTypes := a.mind.actionTypes()

	highestValue := 0.0
	var bestActionTypes []actionType
	for _, t := range actionTypes {
		isActive := false
		// if we currently have an active action of this type, we do not want to start this action
		for _, ac := range a.activity.activeActions {
			if ac.getType() == t {
				isActive = true
				break
			}
		}

		if isActive {
			continue
		}

		// TODO what if an action cannot be started
		// highest value is to eat an apple, but there is no apple, we should somehow start thinking
		// about how to obtain an apple

		v := actionTypeValue(t)
		if v >= highestValue {
			canStart := true
			for startCond := range t.getConditions()[actionConditionTypeStart] {
				if !startCond.isSatisfied(a) {
					canStart = false
				}
			}

			if canStart {
				if v > highestValue {
					highestValue = v
					bestActionTypes = []actionType{}
				}
				bestActionTypes = append(bestActionTypes, t)
			}
		}
	}

	if len(bestActionTypes) == 0 {
		return
	}

	bestActionType := bestActionTypes[rand.Intn(len(bestActionTypes))]
	for startCondition := range bestActionType.getConditions()[actionConditionTypeStart] {
		if !startCondition.isSatisfied(a) {
			return
		}
	}

	newAction := bestActionType.instantiate().(action)
	a.activity.activeActions = append(a.activity.activeActions, newAction)
}

func (a *Agent) stepActions() {
	// actions that are still active after current step
	filteredActions := make([]action, 0)
	for _, ac := range a.activity.activeActions {
		if ac.getState() == actionStateDone {
			continue
		}

		if ac.getState() == actionStateIdle {
			if !ac.start(a) {
				continue
			}
		}

		if ac.getState() == actionStateActive {
			ac.step(a)
		}

		filteredActions = append(filteredActions, ac)
	}

	a.activity.activeActions = filteredActions
}

func (a *Agent) recordActionChanges(changes []change) {
	a.mind.changes = append(a.mind.changes, changes...)
}

func (a *Agent) updateActionCausations() {
	for _, c := range a.mind.changes {
		for _, ac := range a.activity.activeActions {
			if ac.getState() != actionStateDone {
				continue
			}

			matched := false
			actionCausations := ac.getType().getCausations()
			for currCausation := range actionCausations {
				if currCausation.change.match(c) {
					matched = true
					currCausation.occurrences++
					break
				}
			}
			if !matched {
				actionCausations[newCausation(c, ac.getType())] = true
			}
		}

		a.mind.addItem(c, changeDefaultImportance)
	}

	a.mind.changes = make([]change, 0)
}

func (a *Agent) processActionResponse(response interface{}) {
	if taste, ok := response.(*world.Taste); ok {
		a.processTaste(taste)
	} else if aaiChange, ok := response.(*world.AtomicActionInterfaceChange); ok {
		a.processAAIChange(aaiChange)
	}
}

func (a *Agent) processTaste(taste *world.Taste) {
	a.state._update(agentStateTypeHunger, -taste.Nutrition)
	if taste.Sweet {
		a.state._update(agentExperienceTypeSweet, 0)
	}
}

func (a *Agent) processAAIChange(c *world.AtomicActionInterfaceChange) {
	a.recordActionChanges([]change{newAAIChange(a.activity.atomicActionInterfaces[c.Interface], c.Enabling)})
}

func NewAgent() *Agent {
	a := &Agent{
		mind:   newMind(),
		unitId: world.NewActor(),
	}

	a.activity = newActivity(a)
	a.state = newAgentState(a)
	return a
}

func init() {

}
