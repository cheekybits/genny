package bugreports

import "github.com/cheekybits/genny/generic"

// NumberType will be replaced in tests
type NumberType generic.Number

// ObjNumberType is the struct used for tests.
type ObjNumberType struct {
	v NumberType
}

// NewNumberTypes calls new on ObjNumberType and instantiates slice.
// Targets github issues #36 and #49
func NewNumberTypes() (*ObjNumberType, []ObjNumberType) {
	n := new(ObjNumberType)
	m := make([]ObjNumberType, 0)
	return n, m
}
