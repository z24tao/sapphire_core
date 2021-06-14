package agent

/*
	The highest abstraction level of knowledge, the purpose of this type is to enable association
	across information category, i.e. seeing <an object> reminds me of <an action>.
*/
type concept interface {
	addAssoc(other concept, strength float64)
	getAssocs() map[concept]float64 // associated concept -> association strength
	//deprecate()                                          // delete all assocs from self, and delete self from all assocs
	toString(indent string, indentFirstLine bool) string // used for debug
}

/*
	Knowledge is often based on types rather than individual instances, this type is used to
	generalize knowledge from instances to corresponding types, so that even with no knowledge
	about an individual instance, we can still make assumptions based on its type. Example:

	Person A: "I have an apple at home."
	Person B: "I would assume it is sweet."
*/
type conceptType interface {
	singletonConcept
	instantiate() concept // create instance from type
}

// the purpose of this struct is to remove duplicated code from implementations
type commonConcept struct {
	assocs map[concept]float64
}

func (c *commonConcept) addAssoc(other concept, strength float64) {
	if c.assocs[other] < strength {
		c.assocs[other] = strength
	}
}

func (c *commonConcept) getAssocs() map[concept]float64 {
	return c.assocs
}

func (c *commonConcept) deprecate() {
	for assoc := range c.assocs {
		delete(assoc.getAssocs(), c)
	}
}

//
// this function exists for commonConcept to implement concept, in order to allow deprecate to access concept assocs,
//   should not be called directly
func (c *commonConcept) toString(indent string, indentFirstLine bool) string {
	return ""
}

func newCommonConcept() *commonConcept {
	return &commonConcept{
		assocs: map[concept]float64{}, // associated concept -> association strength
	}
}
