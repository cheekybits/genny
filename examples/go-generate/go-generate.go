package gogenerate

import "github.com/mauricelam/genny/generic"

//go:generate genny -in=$GOFILE -out=gen-$GOFILE gen "KeyType=string,int ValueType=string,int"

type KeyType generic.Type
type ValueType generic.Type

type KeyTypeValueTypeMap map[KeyType]ValueType

func NewKeyTypeValueTypeMap() map[KeyType]ValueType {
	return make(map[KeyType]ValueType)
}
