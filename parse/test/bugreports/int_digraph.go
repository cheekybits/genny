// Code generated with https://github.com/cheekybits/genny DO NOT EDIT.
// Any changes will be lost if this file is regenerated.

package bugreports

type DigraphInt struct {
	nodes map[int][]int
}

func NewDigraphInt() *DigraphInt {
	return &DigraphInt{
		nodes: make(map[int][]int),
	}
}

func (dig *DigraphInt) Add(n int) {
	if _, exists := dig.nodes[n]; exists {
		return
	}

	dig.nodes[n] = nil
}

func (dig *DigraphInt) Connect(a, b int) {
	dig.Add(a)
	dig.Add(b)

	dig.nodes[a] = append(dig.nodes[a], b)
}
