package agent

func (a *Agent) _act() {
	// assumes every object visited during observe is already in mind with importance
	// uses object information and changes to update activity
	a._mind._spawnThoughts()
	a._mind._mergeNewThoughts()
	a._startNewAction()
	a._stepActions()
}

func (a *Agent) _startNewAction() {
	actionTypes := a._mind._actionTypes()

	highestValue := 0.0
	var bestActionType actionType = nil
	for _, t := range actionTypes {
		isActive := false
		for _, ac := range a._activity._activeActions {
			if ac._type() == t {
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
		if v > highestValue {
			highestValue = v
			bestActionType = t
		}
	}

	if bestActionType == nil {
		return
	}

	for _, startCondition := range bestActionType._startConditions() {
		if !startCondition._isSatisfied(a) {
			return
		}
	}

	newAction := bestActionType._instantiate().(action)
	a._activity._activeActions = append(a._activity._activeActions, newAction)
}

func (a *Agent) _stepActions() {
	filteredActions := make([]action, 0)
	for _, ac := range a._activity._activeActions {
		switch ac._state() {
		case actionStateComplete:
			continue
		case actionStateIdle:
			if !ac._start(a) {
				continue
			}
		case actionStateActive:
			ac._step(a)
		}

		filteredActions = append(filteredActions, ac)
	}

	a._activity._activeActions = filteredActions
}
