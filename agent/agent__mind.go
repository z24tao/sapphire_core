package agent

import (
	"fmt"
	"sort"
)

const mindCapacity = 30

type mind struct {
	thoughts    map[concept]float64
	newThoughts map[concept][]float64
	changes     []change
}

func (m *mind) objects() []object {
	var items []object
	for t := range m.thoughts {
		if o, ok := t.(object); ok {
			items = append(items, o)
		}
	}

	return items
}

func (m *mind) objectTypes() []objectType {
	var items []objectType

	for t := range m.thoughts {
		if o, ok := t.(objectType); ok {
			items = append(items, o)
		}
	}

	return items
}

func (m *mind) actions() []action {
	var items []action

	for t := range m.thoughts {
		if o, ok := t.(action); ok {
			items = append(items, o)
		}
	}

	return items
}

func (m *mind) actionTypes() []actionType {
	var items []actionType

	for t := range m.thoughts {
		if o, ok := t.(actionType); ok {
			items = append(items, o)
		}
	}

	return items
}

func (m *mind) addItem(c concept, importance float64) {
	m.newThoughts[c] = append(m.newThoughts[c], importance)
}

func (m *mind) spawnThoughts() {
	m.spawnRelatedThoughts()
	m.spawnSpontaneousThoughts()
}

func (m *mind) spawnRelatedThoughts() {
	for thoughtConcept, thoughtImportance := range m.thoughts {
		for assoc, assocStrength := range thoughtConcept.getAssocs() {
			m.addItem(assoc, thoughtImportance*assocStrength)
		}
	}
}

func (m *mind) spawnSpontaneousThoughts() {
	// do nothing for now
}

func (m *mind) mergeNewThoughts() {
	newThoughts := make(map[concept]float64)
	for newConcept, newImportances := range m.newThoughts {
		newImportances = append(newImportances, m.thoughts[newConcept]*0.8) // add old importance
		unimportance := 1.0
		for _, importance := range newImportances {
			unimportance *= 1 - importance
		}
		newThoughts[newConcept] = 1 - unimportance
	}

	m.thoughts = newThoughts
	m.newThoughts = map[concept][]float64{}
	m.filterThoughts()
}

type sortableThought struct {
	c          concept
	importance float64
}

func (m *mind) filterThoughts() {
	thoughtList := make([]sortableThought, 0)
	for c, importance := range m.thoughts {
		thoughtList = append(thoughtList, sortableThought{c, importance})
	}

	sort.SliceStable(thoughtList, func(i, j int) bool {
		return thoughtList[i].importance > thoughtList[j].importance
	})

	i := 0
	m.thoughts = make(map[concept]float64)
	for _, t := range thoughtList {
		if i > mindCapacity {
			break
		}

		m.thoughts[t.c] = t.importance
		i++
	}
}

func (m *mind) toString() string {
	result := "thoughts: [\n"
	for c, i := range m.thoughts {
		result += c.toString("  ", true) + ", importance: " + fmt.Sprintf("%.2f", i) + "\n"
	}
	result += "]"
	return result
}

func newMind() *mind {
	return &mind{
		thoughts:    map[concept]float64{},
		newThoughts: map[concept][]float64{},
	}
}
