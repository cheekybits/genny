package out_test

import (
	"github.com/cheekybits/genny/out"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"testing"
)

const testFileName = "test-file.go"

func tearDown() {
	var err = os.Remove(testFileName)
	if err != nil && !os.IsNotExist(err) {
		panic("Could not delete test file")
	}
}

func assertFileContains(t *testing.T, expected string) {
	file, err := os.Open(testFileName)
	if err != nil {
		panic(err)
	}
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}
	assert.Equal(t, expected, string(fileBytes), "File contents not written properly")
}

func TestMultipleWrites(t *testing.T) {
	defer tearDown()
	lf := out.LazyFile{FileName: testFileName}
	defer lf.Close()
	lf.Write([]byte("Word1"))
	lf.Write([]byte("Word2"))
	assertFileContains(t, "Word1Word2")
}

func TestNoWrite(t *testing.T) {
	defer tearDown()
	lf := out.LazyFile{FileName: testFileName}
	defer lf.Close()
	_, err := os.Stat(testFileName)
	assert.True(t, os.IsNotExist(err), "Expected file not to be created")
}
