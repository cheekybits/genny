package buildtags

import (
	"fmt"

	"github.com/moriyoshi/genny/generic"
)

// +build x,y z
// +build genny

type _t_ generic.Type

func _t_Print(t _t_) {
	fmt.Println(t)
}
