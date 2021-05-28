package extrapackages

import "github.com/cheekybits/genny/generic"

type ForeignType generic.Type

func ForeignTypeSayHello(a ForeignType) ForeignType {
	return a
}
