package bugreports

import "github.com/cheekybits/genny/generic"

type SomeThing generic.Type

func ContainsSomeThing(slice []SomeThing, element SomeThing) bool {
	return false
}

// ContainsAllSomeThing targets github issue 36
func ContainsAllSomeThing(slice []SomeThing, other []SomeThing) bool {
	for _, e := range other {
		if !ContainsSomeThing(slice, e) {
			return false
		}
	}
	return true
}
