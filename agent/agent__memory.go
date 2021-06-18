package agent

import (
	"reflect"
)

// memory - where singleton concepts are stored.
// singleton concepts: every concept that cannot be constructed via some conceptType.instantiate().
// the purpose of this struct is to prevent the same struct from being created repeatedly, e.g. condition.
type memory struct {
	singletons map[reflect.Type]map[concept]bool
}

// add - Adds a concept into memory based on concept type, if concept type does not exist, it is created.
func (m *memory) add(c concept) concept {
	if existing := m.find(c); existing != nil {
		return existing
	}

	cType := reflect.TypeOf(c)
	if _, seen := m.singletons[cType]; !seen {
		m.singletons[cType] = map[concept]bool{}
	}

	m.singletons[cType][c] = true
	return c
}

// find - Attempts to find a match of a concept, otherwise returns nil.
func (m *memory) find(c concept) concept {
	for singleton := range m.singletons[reflect.TypeOf(c)] {
		if singleton.match(c) {
			return singleton
		}
	}
	return nil
}

// remove - Removes a concept from memory if it exists.
func (m *memory) remove(c concept) {
	if m.find(c) != nil {
		delete(m.singletons[reflect.TypeOf(c)], c)
	}
}

// newMemory - Constructor for memory
func newMemory() *memory {
	return &memory{
		singletons: map[reflect.Type]map[concept]bool{},
	}
}
