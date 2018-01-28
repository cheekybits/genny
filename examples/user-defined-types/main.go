package main

import (
	"fmt"

	"github.com/mauricelam/genny/examples/user-defined-types/person"
	"github.com/mauricelam/genny/examples/user-defined-types/pet"
)

//go:generate genny -pkg=main -in=pair/pair.go -out=gen-$GOFILE -imp "github.com/mauricelam/genny/examples/user-defined-types/person" -imp "github.com/mauricelam/genny/examples/user-defined-types/pet" gen "FirstType=Person:person.Person SecondType=Dog:pet.Dog"

func main() {
	p := PairPersonDog{
		person.Person{"John", "Doe"},
		pet.Dog{"ThePet"},
	}
	fmt.Printf("%v, %v\n", p.Left(), p.Right().Name)
}
