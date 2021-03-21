package agent

import (
	"fmt"
	"sort"
)

const assocDefaultImportanceMultiplier = 0.5
const mindCapacity = 10

type mind struct {
	_thoughts map[concept]float64
	_newThoughts map[concept][]float64
}

func (m *mind) _objects() []object {
	var items []object

	for t := range m._thoughts {
		if o, ok := t.(object); ok {
			items = append(items, o)
		}
	}

	return items
}

func (m *mind) _objectTypes() []objectType {
	var items []objectType

	for t := range m._thoughts {
		if o, ok := t.(objectType); ok {
			items = append(items, o)
		}
	}

	return items
}

func (m *mind) _actions() []action {
	var items []action

	for t := range m._thoughts {
		if o, ok := t.(action); ok {
			items = append(items, o)
		}
	}

	return items
}

func (m *mind) _actionTypes() []actionType {
	var items []actionType

	for t := range m._thoughts {
		if o, ok := t.(actionType); ok {
			items = append(items, o)
		}
	}

	return items
}

func (m *mind) _addItem(c concept, importance float64) {
	m._newThoughts[c] = append(m._newThoughts[c], importance)
}

func (m *mind) _spawnThoughts() {
	m._spawnRelatedThoughts()
	m._spawnSpontaneousThoughts()
}

func (m *mind) _spawnRelatedThoughts() {
	for thoughtConcept, thoughtImportance := range m._thoughts {
		for _, assoc := range thoughtConcept._assocs() {
			m._addItem(assoc, thoughtImportance * assocDefaultImportanceMultiplier)
		}
	}
}

func (m *mind) _spawnSpontaneousThoughts() {
	// do nothing for now
}

func (m *mind) _mergeNewThoughts() {
	newThoughts := make(map[concept]float64)
	for c, importances := range m._newThoughts {
		importances = append(importances, m._thoughts[c])
		d := 1.0
		for _, importance := range importances {
			d *= importance
		}
		newThoughts[c] = 1 - d
	}

	m._thoughts = newThoughts
	m._filterThoughts()
}

type sortableThought struct {
	c concept
	importance float64
}

func (m *mind) _filterThoughts() {
	thoughtList := make([]sortableThought, 0)
	for c, importance := range m._thoughts {
		thoughtList = append(thoughtList, sortableThought{c, importance})
	}

	sort.SliceStable(thoughtList, func(i, j int) bool {
		return thoughtList[i].importance > thoughtList[j].importance
	})

	i := 0
	m._thoughts = make(map[concept]float64)
	for _, t := range thoughtList {
		if i > mindCapacity {
			break
		}

		m._thoughts[t.c] = t.importance
		i ++
	}
}

func (m *mind) _toString() string {
	result := "thoughts: [\n"
	for c, i := range m._thoughts {
		result += "  " + c._toString() + ": " + fmt.Sprintf("%v", i) + "\n"
	}
	//result += "]\nnew thoughts: [\n"
	//for c, i := range m._newThoughts {
	//	result += "  " + c._toString() + ": " + fmt.Sprintf("%v", i) + "\n"
	//}
	result += "]"
	return result
}

func newMind() *mind {
	return &mind{
		_thoughts: map[concept]float64{},
		_newThoughts: map[concept][]float64{},
	}
}
