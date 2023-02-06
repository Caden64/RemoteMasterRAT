//go:build darwin

package darwin

import (
	"bytes"
	"os/exec"
	"strings"
)

func BuildNumberDarwin() (string, error) {
	cmd := exec.Command("sw_vers")
	cmd.Stdin = strings.NewReader("Version input")
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		return "", err
	}
	return string(out.Bytes()), nil
}
