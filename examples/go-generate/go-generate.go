package gogenerate

import "github.com/metabition/genny/generic"

//go:generate genny gen "KeyType=string,int ValueType=string,int" -f=$GOFILE

type KeyType generic.Type
type ValueType generic.Type

type KeyTypeValueTypeMap map[KeyType]ValueType

func NewKeyTypeValueTypeMap() map[KeyType]ValueType {
	return make(map[KeyType]ValueType)
}
