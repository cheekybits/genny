// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/falun/genny

package buildtags

import "fmt"

// +build x,y z
// +build genny

func stringPrint(t string) {
	fmt.Println(t)
}
