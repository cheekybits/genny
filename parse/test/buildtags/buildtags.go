package buildtags

import (
	"fmt"

	"github.com/falun/genny/generic"
)

// +build x,y z
// +build genny

type _t_ generic.Type

func _t_Print(t _t_) {
	fmt.Println(t)
}
