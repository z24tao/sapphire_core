package agent

type condition interface {
	_isSatisfied(a *Agent) bool
}

type tautologyCondition struct {}

func (c *tautologyCondition) _isSatisfied(a *Agent) bool {
	return true
}

type contradictionCondition struct {}

func (c *contradictionCondition) _isSatisfied(a *Agent) bool {
	return false
}

type andCondition struct {
	__conditions []condition
}

func (c *andCondition) _isSatisfied(a *Agent) bool {
	for _, condition := range c.__conditions {
		if !condition._isSatisfied(a) {
			return false
		}
	}
	return true
}

type orCondition struct {
	__conditions []condition
}

func (c *orCondition) _isSatisfied(a *Agent) bool {
	for _, condition := range c.__conditions {
		if condition._isSatisfied(a) {
			return true
		}
	}
	return false
}

type notCondition struct {
	__condition condition
}

func (c *notCondition) _isSatisfied(a *Agent) bool {
	return !c.__condition._isSatisfied(a)
}

type hasObjectCondition struct {
	__object object
}

func (c *hasObjectCondition) _isSatisfied(a *Agent) bool {
	c.__object._attrs()[attrTypeDirection] = directionOrigin
	for _, mindObj := range a._mind._objects() {
		if c.__object._match(mindObj._type(), mindObj._attrs(), nil) {
			return true
		}
	}
	return false
}

