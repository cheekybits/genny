package bugreports

import "github.com/cheekybits/genny/generic"

// Tests for issues raised in github issues #36 and #49
type NumberType generic.Number

type ObjNumberType struct {
	v NumberType
}

func NewNumberTypes() (*ObjNumberType, []ObjNumberType) {
	n := new(ObjNumberType)
	m := make([]ObjNumberType, 0)
	return n, m
}
