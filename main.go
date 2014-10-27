package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"

	"github.com/metabition/genny/parse"
)

/*

  source | genny gen "KeyType=string,int ValueType=string,int"

*/

const (
	_ = iota
	exitcodeInvalidArgs
	exitcodeInvalidTypeSet
	exitcodeStdinFailed
	exitcodeGenFailed
)

func main() {

	if len(os.Args) != 3 {
		usage()
		os.Exit(exitcodeInvalidArgs)
	}
	if strings.ToLower(os.Args[1]) != "gen" {
		usage()
		os.Exit(exitcodeInvalidArgs)
	}

	// parse the typesets
	typeSets, err := parse.TypeSet(os.Args[2])
	if err != nil {
		fatal(exitcodeInvalidTypeSet, err)
	}

	source, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fatal(exitcodeStdinFailed, err)
	}
	reader := bytes.NewReader(source)

	// do the work
	if err := gen("stdin", reader, typeSets, os.Stdout); err != nil {
		fatal(exitcodeGenFailed, err)
	}

}

func usage() {
	fmt.Fprintln(os.Stderr, `usage: genny gen "{types}"

gen - generates type specific code (to stdout) from generic code (via stdin)

{types}  - (required) Specific types for each generic type in the source
{types} format:  {generic}={specific}[,another][ {generic2}={specific2}]
Examples:
  Generic=Specific
  Generic1=Specific1 Generic2=Specific2
  Generic1=Specific1,Specific2 Generic2=Specific3,Specific4
`)
}

func fatal(code int, a ...interface{}) {
	fmt.Println(a...)
	os.Exit(code)
}

// gen performs the generic generation.
func gen(filename string, in io.ReadSeeker, typesets []map[string]string, out io.Writer) error {

	var output []byte
	var err error

	output, err = parse.Types(filename, in, typesets)
	if err != nil {
		return err
	}

	out.Write(output)
	return nil
}
