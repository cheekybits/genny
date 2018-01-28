package main

import (
	"fmt"

	"github.com/mauricelam/genny/generic"
)

type Stringer interface {
	generic.Type
	fmt.Stringer
}

func JoinStringers(list []Stringer, sep string) (result string) {
	for i, elem := range list {
		if i > 0 {
			result += sep
		}
		result += elem.String()
	}
	return
}
