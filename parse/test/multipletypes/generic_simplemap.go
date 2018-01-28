package multipletypes

import "github.com/mauricelam/genny/generic"

type KeyType generic.Type
type ValueType generic.Type

type KeyTypeValueTypeMap map[KeyType]ValueType

func (m KeyTypeValueTypeMap) Has(key KeyType) bool {
	_, ok := m[key]
	return ok
}

func (m KeyTypeValueTypeMap) Get(key KeyType) ValueType {
	return m[key]
}

func (m KeyTypeValueTypeMap) Set(key KeyType, value ValueType) KeyTypeValueTypeMap {
	m[key] = value
	return m
}
