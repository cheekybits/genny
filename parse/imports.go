package parse

import (
	"io/ioutil"
	"os/exec"
)

const (
	goimportsCmd = "goimports"
)

func Process(input []byte, outputFileName string) ([]byte, error) {
	if err := ioutil.WriteFile(outputFileName, input, 0644); err != nil {
		return nil, err
	}

	cmd := exec.Command(goimportsCmd, outputFileName)
	output, err := cmd.Output()
	return output, err
}
