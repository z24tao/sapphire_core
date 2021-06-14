package agent

type singletonConcept interface {
	concept
	match(other singletonConcept) bool
}
