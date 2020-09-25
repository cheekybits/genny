package bugreports

import "github.com/cheekybits/genny/generic"

type SomeThing generic.Type

// foo=SomeThing,OtherThing
func foo(x SomeThing) {}
