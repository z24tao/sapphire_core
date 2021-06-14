package agent

type condition interface {
	concept
	isSatisfied(a *Agent) bool
	match(other singletonConcept) bool
}