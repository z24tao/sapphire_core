package agent

// TODO match should return a float for match rate
// ideally, if the agent is 50% hungry but only has experience (_matchOld) for
// being 30% hungry and being 70% hungry, perhaps the corresponding actions
// should be added with modified importance
type change interface {
	concept
	_match(n change) bool
	_matchOld(a *Agent, oldItem concept) bool
	_matchNew(a *Agent, newItem concept) bool
	_getValue() float64
	_setValue(newValue float64)
}

// TODO matchNew implies the agent imagined a future it rather be in, which requires the object
// interface to be able to distinguish between present and future, add this.
