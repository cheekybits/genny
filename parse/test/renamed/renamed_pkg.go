package renamed

import (
	"fmt"

	"github.com/mauricelam/genny/generic"

	testpkg "github.com/mauricelam/genny/parse/test/renamed/subpkg"
)

type _t_ generic.Type

func someFunc_t_() {
	var t _t_
	fmt.Println(t)
	fmt.Println(testpkg.Bar)
}
