// Code generated with https://github.com/cheekybits/genny DO NOT EDIT.
// Any changes will be lost if this file is regenerated.

package multipletypes

type InterfaceIntMap map[interface{}]int

func (m InterfaceIntMap) Has(key interface{}) bool {
	_, ok := m[key]
	return ok
}

func (m InterfaceIntMap) Get(key interface{}) int {
	return m[key]
}

func (m InterfaceIntMap) Set(key interface{}, value int) InterfaceIntMap {
	m[key] = value
	return m
}
