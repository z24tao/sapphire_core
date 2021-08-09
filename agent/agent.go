package agent

import (
	"fmt"
	"github.com/z24tao/sapphire_core/config"
	"github.com/z24tao/sapphire_core/util"
	"github.com/z24tao/sapphire_core/world"
	"math/rand"
)

const imageDefaultImportance = 0.5
const changeDefaultImportance = 0.3
const curiosityValue = 50.0

// Agent - Singular Actor.
type Agent struct {
	state    *agentState
	mind     *mind
	activity *activity
	memory   *memory
	unitId   int
}

// TimeStep - One unit of time passing: updates Agent state, observes environment, and acts.
func (a *Agent) TimeStep() {
	a.updateState()
	a.observe() // observe and learn
	a.act()     // think and act

	if rand.Float64() < 0.01 {
		fmt.Println(a.mind.toString())
	}
}

// updateState - Updates the state of agent.
func (a *Agent) updateState() {
	for stateKey, stateVal := range a.state.states {
		if stateInfo, seen := agentStateInfos[stateKey]; seen {
			a.state.states[stateKey] += stateInfo.perTurn
		} else if stateVal > 0 { // experience fades away if remaining number of turns is positive
			a.state.states[stateKey]--
		}
	}
}

// observe - The Agent perceives the environment (discrete).
func (a *Agent) observe() {
	imgs := world.Look(a.unitId)
	for _, img := range imgs {
		a.processImage(img)
	}
	a.updateActionCausations()
}

// processImage - Process the environment attributes and objects, add to mind if applicable.
func (a *Agent) processImage(img *world.Image) {
	attrs := a.identifyAttributes(img)
	obj := a.identifyObjectInst(attrs)

	if obj == nil {
		a.createObject(attrs, img.DebugName)
	}
}

// identifyAttributes - Identifies Attributes of the world image:
// Includes: Color, Shape, Distance (X + Z), Direction (X , Z).
func (a *Agent) identifyAttributes(img *world.Image) map[int]int {
	attrs := make(map[int]int)

	attrs[attrTypeColor] = img.Color
	attrs[attrTypeShape] = img.Shape
	attrs[attrTypeDistance] = util.Abs(img.XDist) + util.Abs(img.ZDist)

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

// identifyObjectInst - Match object instance in environment to object instance existing in mind based on attributes.
func (a *Agent) identifyObjectInst(attrs map[int]int) object {
	mindObjs := a.mind.objects()
	challengeObj := &simpleObject{
		commonObject: a.newCommonObject(),
		objectType:   nil,
	}
	challengeObj.attrs = attrs

	for _, mindObj := range mindObjs {
		if challengeObj.match(mindObj) {
			changes := mindObj.setAttrs(a, attrs)
			a.mind.addItem(mindObj, imageDefaultImportance)
			a.recordActionChanges(changes)
			return mindObj
		}
	}

	return nil
}

// createObject - Instantiates a new Object type if not exists, then creates an instance in mind.
func (a *Agent) createObject(attrs map[int]int, debugName string) {
	newType := a.identifyObjectType(attrs)
	if newType == nil {
		newType = a.createObjectType(attrs, debugName)
	}

	a.createObjectInst(attrs, newType)
}

// identifyObjectType - Create a new temporary object type and attempts to match with existing object type.
// Return existing otherwise nil.
func (a *Agent) identifyObjectType(attrs map[int]int) objectType {
	mindObjTypes := a.mind.objectTypes()
	challengeObjType := a.newSimpleObjectType("")
	challengeObjType.attrs = attrs
	defer a.memory.remove(challengeObjType)

	for _, objType := range mindObjTypes {
		if objType.match(challengeObjType) {
			return objType
		}
	}

	return nil
}

// createObjectType - Creates a new Simple Object Type and adds it to mind with imageDefaultImportance.
// Uses Color, Shape
func (a *Agent) createObjectType(attrs map[int]int, debugName string) objectType {
	objType := a.newSimpleObjectType(debugName)
	for attrType, attrVal := range attrs {
		if visualAttrTypes[attrType] {
			objType.attrs[attrType] = attrVal
		}
	}

	a.mind.addItem(objType, imageDefaultImportance)
	return objType
}

// createObjectInst - Instantiates a new object instance and adds it to mind.
func (a *Agent) createObjectInst(attrs map[int]int, newType objectType) object {
	newInst := newType.instantiate().(object)
	newInst.setAttrs(a, attrs)
	a.mind.addItem(newInst, imageDefaultImportance)
	return newInst
}

// act - Agent action cycle: spawn, process thoughts -> start action.
func (a *Agent) act() {
	// assumes every object visited during observe is already in mind with importance
	// uses object information and changes to update activity
	a.mind.spawnThoughts()
	a.mind.mergeNewThoughts()
	a.startNewAction()
	a.stepActions()
}

// startNewAction - Chooses action and begins it.
// Defaults to highest value action that is not currently being done.
func (a *Agent) startNewAction() {
	actionTypes := a.mind.actionTypes()

	highestValue := 0.0
	var bestActionTypes []actionType
	for _, t := range actionTypes {
		isActive := false
		// if we currently have an active action, we do not want to start a new action
		for _, ac := range a.activity.activeActions {
			if ac.getState() == actionStateActive {
				isActive = true
				break
			}
		}

		if isActive {
			return
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
	a.mind.addItem(bestActionType, 1.0)

	// add pre-action conditions for hypothesis training
	for cond := range a.getConditions() {
		preActionConditions := newAction.getType().getConditions()[actionConditionTypeObservedAtStart]
		preActionConditions[cond] = true
		newAction.getPreConditions()[cond] = true
	}
}

// stepActions - Check status of actions taken in step, removes Idle or Done actions.
func (a *Agent) stepActions() {
	// actions that are still active after current step
	filteredActions := make([]action, 0)

	var prevAc, currAc action

	for _, ac := range a.activity.activeActions {
		if ac.getState() == actionStateDone {
			prevAc = ac
			continue
		}

		if ac.getState() == actionStateIdle {
			if !ac.start(a) {
				continue
			} else {
				ac.getType().attempt()
			}
		}

		if ac.getState() == actionStateActive {
			ac.step(a)
			currAc = ac

			if ac.getState() == actionStateDone {
				acPostConds := ac.getPostConditions()
				for cond := range a.getConditions() {
					acPostConds[cond] = true
				}
				ac.buildCausations()
			}
		}

		filteredActions = append(filteredActions, ac)
	}

	a.activity.activeActions = filteredActions
	a.buildSequentialActions(prevAc, currAc)
}

// recordActionChanges - Saves current changes of mind into culminate changes.
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
				actionCausations[a.newCausation(c, ac.getType())] = true
			}
		}

		a.mind.addItem(c, changeDefaultImportance)
	}

	for _, ac := range a.activity.activeActions {
		if ac.getState() != actionStateDone {
			continue
		}

		a.buildActionHypotheses(ac)
		//a.evaluateActionHypotheses(ac)
		a.buildConditionalActions(ac)
	}

	if config.Cfg.DebugMindChanges {
		fmt.Println(len(a.mind.changes), "changes")
		for _, ch := range a.mind.changes {
			fmt.Println(ch.toString("", true, true))
		}
	}

	a.mind.changes = make([]change, 0)
}

// buildActionHypotheses - generate atomic action hypotheses for conditional actions.
func (a *Agent) buildActionHypotheses(ac action) {
	// for now, only allow hypotheses on atomic actions
	if _, ok := ac.(*atomicAction); !ok {
		return
	}

	forwardHypotheses, backwardHypotheses := ac.getType().getHypotheses()
	preconditions := ac.getPreConditions()
	for cond := range preconditions {
		if _, seen := forwardHypotheses[cond]; !seen {
			forwardHypotheses[cond] = map[change]*hypothesis{}
		}
	}

	for _, currChange := range a.mind.changes {
		if _, seen := backwardHypotheses[currChange]; !seen {
			backwardHypotheses[currChange] = map[condition]*hypothesis{}
		}
	}

	for cond := range preconditions {
		for _, currChange := range a.mind.changes {
			if _, seen := forwardHypotheses[cond][currChange]; !seen {
				forwardHypotheses[cond][currChange] = newHypothesis(cond, currChange, true)
			}

			if _, seen := backwardHypotheses[currChange][cond]; !seen {
				backwardHypotheses[currChange][cond] = newHypothesis(cond, currChange, false)
			}
		}
	}

	for cond := range preconditions {
		for _, currHypothesis := range forwardHypotheses[cond] {
			currHypothesis.conditionMatch++
		}
	}

	for _, currChange := range a.mind.changes {
		for _, currHypothesis := range backwardHypotheses[currChange] {
			currHypothesis.changeMatch++
		}
	}

	for cond := range preconditions {
		for _, currChange := range a.mind.changes {
			forwardHypotheses[cond][currChange].changeMatch++
			forwardHypotheses[cond][currChange].bothMatch++

			backwardHypotheses[currChange][cond].conditionMatch++
			backwardHypotheses[currChange][cond].bothMatch++
		}
	}
}

// evaluateActionHypotheses - evaluate atomic action hypotheses for conditions and causation.
func (a *Agent) evaluateActionHypotheses(ac action) {
	// for now, only allow hypotheses on atomic actions
	if _, ok := ac.(*atomicAction); !ok {
		return
	}

	if rand.Intn(10) == 0 {
		if config.Cfg.DebugActionHypotheses {
			fmt.Println("===== evaluate action hypotheses =====")
			fmt.Println(ac.getType().toString("", true, true))
			fmt.Println()
		}
		forwardHypotheses, backwardHypotheses := ac.getType().getHypotheses()
		for _, row := range forwardHypotheses {
			for _, h := range row {
				fmt.Println(h.toString("  ", true, true))
			}
		}
		for _, row := range backwardHypotheses {
			for _, h := range row {
				fmt.Println(h.toString("  ", true, true))
			}
		}
	}
}

// buildConditionalActions - Using hypotheses, generate conditional actions.
func (a *Agent) buildConditionalActions(ac action) {
	// for now, only allow hypotheses on atomic actions
	if _, ok := ac.(*atomicAction); !ok {
		return
	}

	// only use forward hypotheses for now
	forwardHypotheses, _ := ac.getType().getHypotheses()
	for _, row := range forwardHypotheses {
		for _, h := range row {
			// if a condition is always satisfied, ignore it
			//   this is to dodge "if apple is red walk towards it"
			if h.conditionMatch > ac.getType().getAttempts()*9/10 {
				continue
			}

			if h.evaluate() > 0.9 {
				ca := a.newConditionalActionType(h.condition, ac.getType(), emptyActionTypeSingleton)
				a.mind.addItem(ca, 1.0)
			}
		}
	}
}

func (a *Agent) buildSequentialActions(prevAc, currAc action) {
	if prevAc == nil || currAc == nil {
		return
	}
	// build sequential actions if:
	//   - two actions reliably (at least 5 attempts, at least 90% success) cause the same quantitative change
	//   - prev action reliably cause a change, the after condition of which matches a start condition of curr action
	prevAttempts := prevAc.getType().getAttempts()
	currAttempts := currAc.getType().getAttempts()
	if prevAttempts < 5 || currAttempts < 5 {
		return
	}

	prevCausations := prevAc.getType().getCausations()
	currCausations := currAc.getType().getCausations()
	for prevCausation := range prevCausations {
		if prevCausation.occurrences*10/9 < prevAttempts {
			return
		}

		for currCausation := range currCausations {
			if currCausation.occurrences*10/9 < currAttempts {
				return
			}

			// if the changes match and is quantitative, build self-recursive sequential action
			if prevCausation.change.match(currCausation.change) {
				if _, ok := prevCausation.change.(*quantitativeAttributeChange); ok {
					sa := a.newSequentialActionType(currAc.getType(), nil)
					sa.next = sa

					for cond := range currAc.getType().getConditions()[actionConditionTypeStart] {
						sa.getConditions()[actionConditionTypeStart][cond] = true
					}
					a.mind.addItem(sa, 1.0)
				}
			}
		}
	}
}

// processActionResponse - Processes environment response to action previously taken.
func (a *Agent) processActionResponse(response interface{}) {
	if taste, ok := response.(*world.Taste); ok {
		a.processTaste(taste)
	} else if aaiChange, ok := response.(*world.AtomicActionInterfaceChange); ok {
		a.processAAIChange(aaiChange)
	}
}

// processTaste - Updates agent state hunger and taste.
func (a *Agent) processTaste(taste *world.Taste) {
	a.state.update(agentStateTypeHunger, -taste.Nutrition)
	if taste.Sweet {
		a.state.update(agentExperienceTypeSweet, 0)
	}
}

// processAAIChange - records action changes for atomic actions.
func (a *Agent) processAAIChange(c *world.AtomicActionInterfaceChange) {
	a.recordActionChanges([]change{a.newAAIChange(a.activity.atomicActionInterfaces[c.Interface], c.Enabling)})
}

// getConditions - Returns conditions of Agent and objects in mind.
func (a *Agent) getConditions() map[condition]bool {
	result := map[condition]bool{}

	for _, o := range a.mind.objects() {
		for attrType, attrVal := range o.getAttrs() {
			if _, seen := qualitativeAttrTypes[attrType]; !seen {
				continue
			}

			result[a.newAttributeCondition(o.getType(), attrType, attrVal)] = true
		}
	}

	for stateType, stateVal := range a.state.states {
		if stateInfo, seen := agentStateInfos[stateType]; seen {
			result[a.newAgentCondition(stateType, stateVal > stateInfo.threshold)] = true
		} else {
			result[a.newAgentCondition(stateType, stateVal > 0)] = true
		}
	}

	return result
}

// NewAgent - New instance of Agent Actor
func NewAgent() *Agent {
	a := &Agent{
		mind:   newMind(),
		memory: newMemory(),
		unitId: world.NewActor(),
	}

	a.activity = newActivity(a)
	a.state = newAgentState(a)

	emptyActionTypeSingleton = &emptyActionType{
		commonActionType: a.newCommonActionType(),
	}

	return a
}

func init() {

}
