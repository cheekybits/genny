package math

import "github.com/metabition/genny/generic"

type NumberType generic.Number

type NumberTypeCompare func(a, b NumberType) bool

func Max(fn NumberTypeCompare, a, b NumberType) NumberType {
	if fn(a, b) {
		return a
	}
	return b
}
