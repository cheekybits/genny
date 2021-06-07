package parse

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_removeFullyQualifiedImportPathAndAddToMap(t *testing.T) {
	type args struct {
		typeSet          map[string]string
		extraPackagesSet map[string]struct{}
	}
	tests := []struct {
		name                 string
		args                 args
		want                 map[string]string
		wantExtraPackagesSet map[string]struct{}
		wantErr              bool
	}{
		{
			name: "one field example",
			args: args{
				typeSet: map[string]string{
					"GenericField": "example.com/a/b.MyType",
				},
				extraPackagesSet: map[string]struct{}{},
			},
			want: map[string]string{
				"GenericField": "b.MyType",
			},
			wantExtraPackagesSet: map[string]struct{}{
				"example.com/a/b": {},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := removeFullyQualifiedImportPathAndAddToMap(tt.args.typeSet, tt.args.extraPackagesSet)
			if (err != nil) != tt.wantErr {
				t.Errorf("removeFullyQualifiedImportPathAndAddToMap() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantExtraPackagesSet, tt.args.extraPackagesSet)
		})
	}
}

func Test_addExtraImports(t *testing.T) {
	type testDef struct {
		Name         string
		In           string
		Out          string
		ExtraImports []string
	}

	tests := []testDef{
		{
			Name: "single import",
			In: `package x

import "fmt"

func sayHello(user userpkg.User) {
	return fmt.Sprintf("hello %s", user.Name)
}
`,
			Out: `package x

import (
	"fmt"
	"example.com/me/userpkg"
)

func sayHello(user userpkg.User) {
	return fmt.Sprintf("hello %s", user.Name)
}
`,
			ExtraImports: []string{"example.com/me/userpkg"},
		}, {
			Name: "no imports",
			In: `package x

func sayHello(user userpkg.User) {
	return "hello " + user.Name
}
`,
			Out: `package x

import (
	"example.com/me/userpkg"
)

func sayHello(user userpkg.User) {
	return "hello " + user.Name
}
`,
			ExtraImports: []string{"example.com/me/userpkg"},
		}, {
			Name: "multiple imports",
			In: `package x

import (
	"fmt"
	"io"
)

func sayHello(writer io.Writer, user userpkg.User) {
	return fmt.Fprintf(writer, "hello %s", user.Name)
}
`,
			Out: `package x

import (
	"fmt"
	"io"
	"example.com/me/userpkg"
)

func sayHello(writer io.Writer, user userpkg.User) {
	return fmt.Fprintf(writer, "hello %s", user.Name)
}
`,
			ExtraImports: []string{"example.com/me/userpkg"},
		},
	}
	for _, test := range tests {

		out, err := addExtraImports(bytes.NewReader([]byte(test.In)), test.ExtraImports)
		assert.NoError(t, err)

		assert.Equal(t, test.Out, string(out), test.Name)
	}
}
