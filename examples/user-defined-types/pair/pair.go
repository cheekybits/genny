package pair

import (
	"github.com/mauricelam/genny/generic"
)

type FirstType generic.Type
type SecondType generic.Type

type PairFirstTypeSecondType struct {
	First  FirstType
	Second SecondType
}

func (p PairFirstTypeSecondType) Left() FirstType {
	return p.First
}

func (p PairFirstTypeSecondType) Right() SecondType {
	return p.Second
}
