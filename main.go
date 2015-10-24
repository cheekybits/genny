package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"runtime/debug"
	"strings"

	"github.com/avoronkov/genny/parse"
)

/*

  source | genny gen [-in=""] [-out=""] [-pkg=""] "KeyType=string,int ValueType=string,int"

*/

const (
	_ = iota
	exitcodeInvalidArgs
	exitcodeInvalidTypeSet
	exitcodeStdinFailed
	exitcodeGenFailed
	exitcodeGetFailed
	exitcodeSourceFileInvalid
	exitcodeDestFileFailed
	exitcodeInternalError
)

func main() {
	var (
		mainErr  error
		exitCode int
	)

	defer func() {
		if r := recover(); r != nil {
			fmt.Fprintf(os.Stderr, "Runtime panic: %v: %s\n", r, debug.Stack())
			exitCode = exitcodeInternalError
		}
		if mainErr != nil {
			fmt.Fprintf(os.Stderr, "Error 111: %v\n", mainErr)
		}
		os.Exit(exitCode)
	}()

	var (
		in      = flag.String("in", "", "file to parse instead of stdin")
		out     = flag.String("out", "", "file to save output to instead of stdout")
		pkgName = flag.String("pkg", "", "package name for generated files")
		imports Strings
		prefix  = "https://github.com/metabition/gennylib/raw/master/"
	)
	flag.Var(&imports, "imp", "spcify import explicitly (can be specified multiple times)")
	flag.Usage = usage
	flag.Parse()
	args := flag.Args()

	if len(args) < 2 {
		usage()
		os.Exit(exitcodeInvalidArgs)
	}

	if strings.ToLower(args[0]) != "gen" && strings.ToLower(args[0]) != "get" {
		usage()
		os.Exit(exitcodeInvalidArgs)
	}

	// parse the typesets
	var setsArg = args[1]
	if strings.ToLower(args[0]) == "get" {
		setsArg = args[2]
	}
	typeSets, err := parse.TypeSet(setsArg)
	if err != nil {
		exitCode, mainErr = exitcodeInvalidTypeSet, err
		return
	}

	var outWriter io.Writer
	if len(*out) > 0 {
		outFile, err := os.Create(*out)
		if err != nil {
			exitCode, mainErr = exitcodeDestFileFailed, err
			return
		}
		defer func(outPath string) {
			outFile.Close()
			if mainErr != nil {
				os.RemoveAll(outPath)
			}
		}(*out)
		outWriter = outFile
	} else {
		outWriter = os.Stdout
	}

	if strings.ToLower(args[0]) == "get" {
		if len(args) != 3 {
			fmt.Println("not enough arguments to get")
			usage()
			os.Exit(exitcodeInvalidArgs)
		}
		r, err := http.Get(prefix + args[1])
		if err != nil {
			exitCode, mainErr = exitcodeGetFailed, err
			return
		}
		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			exitCode, mainErr = exitcodeGetFailed, err
			return
		}
		r.Body.Close()
		br := bytes.NewReader(b)
		err = gen(*in, *pkgName, br, typeSets, imports, outWriter)
	} else if len(*in) > 0 {
		var file *os.File
		file, err = os.Open(*in)
		if err != nil {
			exitCode = exitcodeSourceFileInvalid
			return
		}
		defer file.Close()
		err = gen(*in, *pkgName, file, typeSets, imports, outWriter)
	} else {
		var source []byte
		source, err = ioutil.ReadAll(os.Stdin)
		if err != nil {
			exitCode = exitcodeStdinFailed
			return
		}
		reader := bytes.NewReader(source)
		err = gen("stdin", *pkgName, reader, typeSets, imports, outWriter)
	}

	// do the work
	if err != nil {
		exitCode, mainErr = exitcodeGenFailed, err
	}
}

func usage() {
	fmt.Fprintln(os.Stderr, `usage: genny [{flags}] gen "{types}"

gen - generates type specific code from generic code.
get <package/file> - fetch a generic template from the online library and gen it.

{types}  - (optional) Command line flags (see below)
{types}  - (required) Specific types for each generic type in the source
{types} format:  {generic}={specific}[,another][ {generic2}={specific2}]

Examples:
  Generic=Specific
  Generic1=Specific1 Generic2=Specific2
  Generic1=Specific1,Specific2 Generic2=Specific3,Specific4
  Generic=SpecificTitle:package.Type,AnotherSpecific

Flags:`)
	flag.PrintDefaults()
}

// gen performs the generic generation.
func gen(filename, pkgName string, in io.ReadSeeker, typesets []map[string]string, imports []string, out io.Writer) error {

	var output []byte
	var err error

	output, err = parse.Generics(filename, pkgName, in, typesets, imports)
	if err != nil {
		return err
	}

	out.Write(output)
	return nil
}

// List of strings for flag
type Strings []string

func (i Strings) String() string {
	return strings.Join([]string(i), ", ")
}

func (i *Strings) Set(value string) error {
	*i = append(*i, value)
	return nil
}
