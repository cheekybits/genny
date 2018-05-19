package bugreports

import "github.com/justnoise/genny/generic"

type NumberType generic.Number

type ObjNumberType struct {
	v NumberType
}

func NewNumberTypes() (*ObjNumberType, []ObjNumberType) {
	n := new(ObjNumberType)
	m := make([]ObjNumberType, 0)
	return n, m
}
