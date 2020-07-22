// Code generated with https://github.com/cheekybits/genny DO NOT EDIT.
// Any changes will be lost if this file is regenerated.

package bugreports

// CellInt is result of generating code via genny for type int
type CellInt struct {
	Value int
}

const constantInt = 1

func funcInt(p CellInt) {}

// exampleInt does some instantation and function calls for types inclueded in this file.
// Targets github issue 15
func exampleInt() {
	aCellInt := CellInt{}
	anotherCellInt := CellInt{}
	if aCellInt != anotherCellInt {
		println(constantInt)
		panic(constantInt)
	}
	funcInt(CellInt{})
}
