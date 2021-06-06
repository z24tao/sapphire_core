package agent

type condition interface {
	concept
	isSatisfied(a *Agent) bool
	match(other condition) bool
}
