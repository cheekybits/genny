package main

import (
	"bytes"
	"flag"
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
	exitcodeInvalidArgs       = 1
	exitcodeInvalidTypeSet    = 2
	exitcodeStdinFailed       = 3
	exitcodeGenFailed         = 4
	exitcodeSourceFileInvalid = 5
)

func main() {
	var (
		filename = flag.String("f", "", "file to parse instead of stdin")
	)
	flag.Parse()
	args := flag.Args()
	if len(args) != 3 {
		usage()
		os.Exit(exitcodeInvalidArgs)
	}
	if strings.ToLower(args[1]) != "gen" {
		usage()
		os.Exit(exitcodeInvalidArgs)
	}

	// parse the typesets
	typeSets, err := parse.TypeSet(args[2])
	if err != nil {
		fatal(exitcodeInvalidTypeSet, err)
	}

	if len(*filename) > 0 {
		file, err := os.Open(*filename)
		if err != nil {
			fatal(exitcodeSourceFileInvalid, err)
		}
		defer file.Close()
		err = gen(*filename, file, typeSets, os.Stdout)
	} else {
		source, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			fatal(exitcodeStdinFailed, err)
		}
		reader := bytes.NewReader(source)
		err = gen("stdin", reader, typeSets, os.Stdout)
	}

	// do the work
	if err != nil {
		fatal(exitcodeGenFailed, err)
	}

}

func usage() {
	fmt.Fprintln(os.Stderr, `usage: genny gen "{types}"

gen - generates type specific code (to stdout) from generic code (via stdin) or the file specified with the -f flag.

{types}  - (required) Specific types for each generic type in the source
{types} format:  {generic}={specific}[,another][ {generic2}={specific2}]

Examples:
  Generic=Specific
  Generic1=Specific1 Generic2=Specific2
  Generic1=Specific1,Specific2 Generic2=Specific3,Specific4

Flags:`)
	flag.PrintDefaults()
}

func fatal(code int, a ...interface{}) {
	fmt.Println(a...)
	os.Exit(code)
}

// gen performs the generic generation.
func gen(filename string, in io.ReadSeeker, typesets []map[string]string, out io.Writer) error {

	var output []byte
	var err error

	output, err = parse.Generics(filename, in, typesets)
	if err != nil {
		return err
	}

	out.Write(output)
	return nil
}
