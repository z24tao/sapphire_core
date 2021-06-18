package agent

import (
	"reflect"
)

// the agent's memory, where singleton concepts are stored
// singleton concepts: every concept that cannot be constructed via some conceptType.instantiate()
// the purpose of this struct is to prevent the same struct from being created repeatedly, e.g. condition
type memory struct {
	singletons map[reflect.Type]map[concept]bool
}

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

func (m *memory) find(c concept) concept {
	for singleton := range m.singletons[reflect.TypeOf(c)] {
		if singleton.match(c) {
			return singleton
		}
	}
	return nil
}

func (m *memory) remove(c concept) {
	if m.find(c) != nil {
		delete(m.singletons[reflect.TypeOf(c)], c)
	}
}

func newMemory() *memory {
	return &memory{
		singletons: map[reflect.Type]map[concept]bool{},
	}
}
