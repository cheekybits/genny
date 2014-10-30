package gogenerate

import "github.com/metabition/genny/generic"

//go:generate genny gen -f=$GOFILE "KeyType=string,int ValueType=string,int"

type KeyType generic.Type
type ValueType generic.Type

type KeyTypeValueTypeMap map[KeyType]ValueType

func NewKeyTypeValueTypeMap() map[KeyType]ValueType {
	return make(map[KeyType]ValueType)
}
