package parse_test

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"testing"

	"github.com/mauricelam/genny/parse"
	"github.com/stretchr/testify/assert"
)

var tests = []struct {
	// input
	filename string
	pkgName  string
	in       string
	tag      string
	imports  []string
	types    []map[string]string

	// expectations
	expectedOut string
	expectedErr error
}{
	{
		filename:    "generic_queue.go",
		in:          `test/queue/generic_queue.go`,
		types:       []map[string]string{{"Something": "int"}},
		expectedOut: `test/queue/int_queue.go`,
	},
	{
		filename:    "generic_queue.go",
		pkgName:     "changed",
		in:          `test/queue/generic_queue.go`,
		types:       []map[string]string{{"Something": "int"}},
		expectedOut: `test/queue/changed/int_queue_newpkg.go`,
	},
	{
		filename:    "generic_queue.go",
		in:          `test/queue/generic_queue.go`,
		types:       []map[string]string{{"Something": "float32"}},
		expectedOut: `test/queue/float32_queue.go`,
	},
	{
		filename:    "generic_simplemap.go",
		in:          `test/multipletypes/generic_simplemap.go`,
		types:       []map[string]string{{"KeyType": "string", "ValueType": "int"}},
		expectedOut: `test/multipletypes/string_int_simplemap.go`,
	},
	{
		filename:    "generic_simplemap.go",
		in:          `test/multipletypes/generic_simplemap.go`,
		types:       []map[string]string{{"KeyType": "interface{}", "ValueType": "int"}},
		expectedOut: `test/multipletypes/interface_int_simplemap.go`,
	},
	{
		filename:    "generic_simplemap.go",
		in:          `test/multipletypes/generic_simplemap.go`,
		types:       []map[string]string{{"KeyType": "*MyType1", "ValueType": "*MyOtherType"}},
		expectedOut: `test/multipletypes/custom_types_simplemap.go`,
	},
	{
		filename:    "generic_internal.go",
		in:          `test/unexported/generic_internal.go`,
		types:       []map[string]string{{"secret": "*myType"}},
		expectedOut: `test/unexported/mytype_internal.go`,
	},
	{
		filename: "generic_simplemap.go",
		in:       `test/multipletypesets/generic_simplemap.go`,
		types: []map[string]string{
			{"KeyType": "int", "ValueType": "string"},
			{"KeyType": "float64", "ValueType": "bool"},
		},
		expectedOut: `test/multipletypesets/many_simplemaps.go`,
	},
	{
		filename:    "generic_number.go",
		in:          `test/numbers/generic_number.go`,
		types:       []map[string]string{{"NumberType": "int"}},
		expectedOut: `test/numbers/int_number.go`,
	},
	{
		filename:    "generic_digraph.go",
		in:          `test/bugreports/generic_digraph.go`,
		types:       []map[string]string{{"Node": "int"}},
		expectedOut: `test/bugreports/int_digraph.go`,
	},
	{
		filename:    "renamed_pkg.go",
		in:          `test/renamed/renamed_pkg.go`,
		types:       []map[string]string{{"_t_": "int"}},
		expectedOut: `test/renamed/renamed_pkg_int.go`,
	},
	{
		filename:    "buildtags.go",
		in:          `test/buildtags/buildtags.go`,
		types:       []map[string]string{{"_t_": "int"}},
		expectedOut: `test/buildtags/buildtags_expected.go`,
		tag:         "genny",
	},
	{
		filename:    "buildtags.go",
		in:          `test/buildtags/buildtags.go`,
		types:       []map[string]string{{"_t_": "string"}},
		expectedOut: `test/buildtags/buildtags_expected_nostrip.go`,
		tag:         "",
	},
}

func TestParse(t *testing.T) {

	for testNo, test := range tests {

		t.Run(fmt.Sprintf("%d:%s", testNo, test.expectedOut), func(t *testing.T) {
			test.in = contents(test.in)
			test.expectedOut = contents(test.expectedOut)

			bytes, err := parse.Generics(test.filename, test.pkgName, strings.NewReader(test.in), test.types, test.imports, test.tag)

			// check the error
			if test.expectedErr == nil {
				assert.NoError(t, err, "(%d: %s) No error was expected but got: %s", testNo, test.filename, err)
			} else {
				assert.NotNil(t, err, "(%d: %s) No error was returned by one was expected: %s", testNo, test.filename, test.expectedErr)
				assert.IsType(t, test.expectedErr, err, "(%d: %s) Generate should return object of type %v", testNo, test.filename, test.expectedErr)
			}

			// assert the response
			if !assert.Equal(t, test.expectedOut, string(bytes), "Parse didn't generate the expected output.") {
				log.Println("EXPECTED: " + test.expectedOut)
				log.Println("ACTUAL: " + string(bytes))
			}
		})

	}

}

func contents(s string) string {
	if strings.HasSuffix(s, "go") {
		file, err := ioutil.ReadFile(s)
		if err != nil {
			panic(err)
		}
		return string(file)
	}
	return s
}
