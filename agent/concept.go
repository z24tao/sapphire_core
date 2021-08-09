package agent

/*
	The highest abstraction level of knowledge, the purpose of this type is to enable association
	across information category, i.e. seeing <an object> reminds me of <an action>.
*/
type concept interface {
	// associated concept -> association strength
	getAssocs() map[concept]float64
	addAssoc(other concept, strength float64)
	match(other concept) bool

	// delete all assocs from self, and delete self from all assocs
	deprecate()

	// used for debug
	toString(indent string, recursive, indentFirstLine bool) string
}

/*
	Knowledge is often based on types rather than individual instances, this type is used to
	generalize knowledge from instances to corresponding types, so that even with no knowledge
	about an individual instance, we can still make assumptions based on its type. Example:

	Person A: "I have an apple at home."
	Person B: "I would assume it is sweet."
*/
type conceptType interface {
	concept
	instantiate() concept // create instance from type
}

// the purpose of this struct is to remove duplicated code from implementations
type commonConcept struct {
	agent  *Agent
	assocs map[concept]float64
}

func (c *commonConcept) addAssoc(other concept, strength float64) {
	for existingAssoc := range c.assocs {
		if existingAssoc.match(other) {
			if c.assocs[existingAssoc] < strength {
				c.assocs[existingAssoc] = strength
			}
			return
		}
	}

	c.assocs[other] = strength
}

func (c *commonConcept) getAssocs() map[concept]float64 {
	result := map[concept]float64{}

	for assoc, strength := range c.assocs {
		result[assoc] = strength
	}

	return result
}

func (c *commonConcept) deprecate() {
	for assoc := range c.assocs {
		delete(assoc.getAssocs(), c)
	}
}

// these functions exists for commonConcept to implement concept, in order to allow deprecate to access concept assocs,
//   should not be called directly
func (c *commonConcept) match(_ concept) bool {
	panic("implement me")
}

func (c *commonConcept) toString(_ string, _, _ bool) string {
	panic("implement me")
}

func (a *Agent) newCommonConcept() *commonConcept {
	return &commonConcept{
		agent:  a,
		assocs: map[concept]float64{}, // associated concept -> association strength
	}
}
