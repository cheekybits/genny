package bugreports

import "github.com/cheekybits/genny/generic"

// X is the type generic type used in tests
type X generic.Type

// CellX is result of generating code via genny for type X
type CellX struct {
	Value X
}

const constantX = 1

func funcX(p CellX) {}

// exampleX does some instantation and function calls for types inclueded in this file.
// Targets github issue 15
func exampleX() {
	aCellX := CellX{}
	anotherCellX := CellX{}
	if aCellX != anotherCellX {
		println(constantX)
		panic(constantX)
	}
	funcX(CellX{})
}
