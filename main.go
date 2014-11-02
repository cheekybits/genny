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

  source | genny gen [-in=""] [-out=""] "KeyType=string,int ValueType=string,int"

*/

const (
	exitcodeInvalidArgs       = 1
	exitcodeInvalidTypeSet    = 2
	exitcodeStdinFailed       = 3
	exitcodeGenFailed         = 4
	exitcodeSourceFileInvalid = 5
	exitcodeDestFileFailed    = 6
)

func main() {
	var (
		in  = flag.String("in", "", "file to parse instead of stdin")
		out = flag.String("out", "", "file to save output to instead of stdout")
	)
	flag.Parse()
	args := flag.Args()

	if len(args) != 2 {
		usage()
		os.Exit(exitcodeInvalidArgs)
	}
	if strings.ToLower(args[0]) != "gen" {
		usage()
		os.Exit(exitcodeInvalidArgs)
	}

	// parse the typesets
	typeSets, err := parse.TypeSet(args[1])
	if err != nil {
		fatal(exitcodeInvalidTypeSet, err)
	}

	var outWriter io.Writer
	if len(*out) > 0 {
		outFile, err := os.Create(*out)
		if err != nil {
			fatal(exitcodeDestFileFailed, err)
		}
		defer outFile.Close()
		outWriter = outFile
	} else {
		outWriter = os.Stdout
	}

	if len(*in) > 0 {
		file, err := os.Open(*in)
		if err != nil {
			fatal(exitcodeSourceFileInvalid, err)
		}
		defer file.Close()
		err = gen(*in, file, typeSets, outWriter)
	} else {
		source, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			fatal(exitcodeStdinFailed, err)
		}
		reader := bytes.NewReader(source)
		err = gen("stdin", reader, typeSets, outWriter)
	}

	// do the work
	if err != nil {
		fatal(exitcodeGenFailed, err)
	}

}

func usage() {
	fmt.Fprintln(os.Stderr, `usage: genny [{flags}] gen "{types}"

gen - generates type specific code from generic code.

{types}  - (optional) Command line flags (see below)
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
