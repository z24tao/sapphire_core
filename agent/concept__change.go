package agent

type change interface {
	concept
	match(other concept) bool
	before() condition
	after() condition
	precedes(other change) bool
	getValue() float64
	setValue(value float64)
}

type commonChange struct {
	*commonConcept
	value float64
}

func (c *commonChange) getValue() float64 {
	return c.value
}

func (c *commonChange) setValue(value float64) {
	c.value = value
}

func (a *Agent) newCommonChange() *commonChange {
	return &commonChange{
		commonConcept: a.newCommonConcept(),
		value:         0,
	}
}
