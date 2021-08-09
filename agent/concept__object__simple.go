package agent

import "fmt"

type simpleObject struct {
	*commonObject
	objectType *simpleObjectType
}

func (o *simpleObject) toString(indent string, _, indentFirstLine bool) string {
	result := ""
	if indentFirstLine {
		result += indent
	}
	result += fmt.Sprintf("simpleObject: %s\n", o.objectType.debugName)
	result += fmt.Sprintf(indent+"  attributes: (%d) [\n", len(o.attrs))

	for attrType, attrVal := range o.attrs {
		if qualitativeAttrTypes[attrType] {
			result += fmt.Sprintf(indent+"    %s: %s\n", attrTypes[attrType], attrVals[attrType][attrVal])
		} else {
			result += fmt.Sprintf(indent+"    %s: %d\n", attrTypes[attrType], attrVal)
		}
	}
	result += indent + "  ]"

	return result
}

func (o *simpleObject) getType() objectType {
	return o.objectType
}

// attrs: attribute type -> attribute value
// updates object attributes, construct and return appropriate changes
func (o *simpleObject) setAttrs(a *Agent, attrs map[int]int) []change {
	var changes []change

	for newType, newVal := range attrs {
		oldVal, seen := o.attrs[newType]

		if seen && oldVal == newVal {
			continue
		}

		// update object attribute
		o.attrs[newType] = newVal

		// for current iteration, we do not record a change for introducing a new attribute type
		if !seen {
			continue
		}

		// for existing attribute type, record appropriate change
		// if attribute type is qualitative (discrete), record qualitative change
		if qualitativeAttrTypes[newType] {
			changes = append(changes, a.newQualitativeAttributeChange(o.getType(), newType, oldVal, newVal))
		}

		// if attribute type is quantitative (continuous), record quantitative change
		if quantitativeAttrTypes[newType] {
			changes = append(changes, a.newQuantitativeAttributeChange(o.getType(), newType, oldVal < newVal))
		}
	}

	return changes
}
