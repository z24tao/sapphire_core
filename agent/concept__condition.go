package agent

type condition interface {
	concept
	isSatisfied(a *Agent) bool
	match(other concept) bool
	buildChanges(other condition) []change
}
