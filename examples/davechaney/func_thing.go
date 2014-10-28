package math

import "github.com/metabition/genny/generic"

type ThisNumberType generic.Number

type ThisNumberTypeCompare func(a, b ThisNumberType) bool

func ThisNumberTypeMax(fn ThisNumberTypeCompare, a, b ThisNumberType) ThisNumberType {
	if fn(a, b) {
		return a
	}
	return b
}
