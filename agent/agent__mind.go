package agent

import (
	"fmt"
	"sort"
)

const mindCapacity = 100

// mind - Agent's mind.
// Currently holds thoughts, newThoughts, and list of changes.
type mind struct {
	thoughts    map[concept]float64
	newThoughts map[concept][]float64
	changes     []change
}

// objects - Parses through the mind's thoughts for objects and returns a list of objects.
func (m *mind) objects() []object {
	var items []object
	for t := range m.thoughts {
		if o, ok := t.(object); ok {
			items = append(items, o)
		}
	}

	return items
}

// objectTypes - Parses through the mind's thoughts for objectTypes and returns a list of objectTypes.
func (m *mind) objectTypes() []objectType {
	var items []objectType

	for t := range m.thoughts {
		if o, ok := t.(objectType); ok {
			items = append(items, o)
		}
	}

	return items
}

// actions - Parses through the mind's thoughts for actions and returns a list of actions.
func (m *mind) actions() []action {
	var items []action

	for t := range m.thoughts {
		if o, ok := t.(action); ok {
			items = append(items, o)
		}
	}

	return items
}

// actionTypes - Parses through the mind's thoughts for actionTypes and returns a list of actionTypes.
func (m *mind) actionTypes() []actionType {
	var items []actionType

	for t := range m.thoughts {
		if o, ok := t.(actionType); ok {
			items = append(items, o)
		}
	}

	return items
}

// addItem - Adds a concept to the mind with a given importance as a newThought.
func (m *mind) addItem(c concept, importance float64) {
	for existingThought := range m.newThoughts {
		if existingThought.match(c) {
			m.newThoughts[existingThought] = append(m.newThoughts[existingThought], importance)
			return
		}
	}

	m.newThoughts[c] = []float64{importance}
}

// spawnThoughts - Generate thoughts related to current existing thoughts as well as spontaneous new thoughts.
func (m *mind) spawnThoughts() {
	m.spawnRelatedThoughts()
	m.spawnSpontaneousThoughts()
}

// spawnRelatedThoughts - Generate thoughts related to existing thoughts.
func (m *mind) spawnRelatedThoughts() {
	for thoughtConcept, thoughtImportance := range m.thoughts {
		for assoc, assocStrength := range thoughtConcept.getAssocs() {
			m.addItem(assoc, thoughtImportance*assocStrength)
		}
	}
}

// TODO: Complete spawnSpontaneousThoughts

// spawnSpontaneousThoughts -
func (m *mind) spawnSpontaneousThoughts() {
	// do nothing for now
}

// mergeNewThoughts - Compare new thoughts and old thoughts and merge importance of thoughts if pre-existing.
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

// sortableThought - Structure to hold thoughts to sport by importance.
type sortableThought struct {
	c          concept
	importance float64
}

// filterThoughts - Sort thoughts by importance and remove non-important thoughts.
// Currently removes until there are <= mindCapacity thoughts. Possible TODO: Different way of filtering
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

// toString - Debug only function.
func (m *mind) toString() string {
	result := "thoughts: [\n"
	for c, i := range m.thoughts {
		result += c.toString("  ", true, true) + ", importance: " + fmt.Sprintf("%.2f", i) + "\n"
	}
	result += "]"
	return result
}

// newMind - New instance of mind.
func newMind() *mind {
	return &mind{
		thoughts:    map[concept]float64{},
		newThoughts: map[concept][]float64{},
	}
}
